package unifi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// API is an interface to a UniFi controller.
type API struct {
	httpClient *http.Client
	cookieBase *url.URL
	auth       *Auth
}

func BuildAPI(username string, password string, controllerhost string) (*API, error) {
	auth := Auth{username, password, controllerhost, nil}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	cookieBase := &url.URL{
		Scheme: "https",
		Host:   auth.ControllerHost,
	}
	jar.SetCookies(cookieBase, auth.Cookies)

	api := &API{
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					// TODO: support proper certs
					InsecureSkipVerify: true,
				},
			},
			Jar: jar,
		},
		cookieBase: cookieBase,
		auth:       &auth,
	}
	return api, nil
}

func (api *API) post(u string, src, dst interface{}, opts reqOpts) error {
	u = api.baseURL() + u
	body, err := json.Marshal(src)
	if err != nil {
		panic("internal error marshaling JSON POST body: " + err.Error())
	}
	req, err := http.NewRequest("POST", u, bytes.NewReader(body))
	if err != nil {
		panic("internal error: " + err.Error())
	}
	return api.processHTTPRequest(req, dst, opts)
}

func (api *API) get(u string, dst interface{}, opts reqOpts) error {
	u = api.baseURL() + u
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		panic("internal error: " + err.Error())
	}
	return api.processHTTPRequest(req, dst, opts)
}

type reqOpts struct {
	referer string
}

func (api *API) processHTTPRequest(req *http.Request, dst interface{}, opts reqOpts) error {
	if opts.referer != "" {
		req.Header.Set("Referer", opts.referer)
	}

	dec := struct {
		Data interface{} `json:"data"`
		Meta struct {
			Code string `json:"rc"`
			Msg  string `json:"msg"`
		} `json:"meta"`
	}{Data: dst}

	triedLogin := false
	for {
		resp, err := api.httpClient.Do(req)
		if err != nil {
			return err
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return err
		}
		fmt.Println(req.Method)
		if req.Method == "POST" {
			fmt.Println(string(body))
		}
		if err := json.Unmarshal(body, &dec); err != nil {
			return fmt.Errorf("parsing response body: %v", err)
		}

		if resp.StatusCode == 200 {
			if dec.Meta.Code != "ok" {
				return fmt.Errorf("non-ok return code %q (%s)", dec.Meta.Code, dec.Meta.Msg)
			}
			return nil
		}

		if resp.StatusCode == http.StatusUnauthorized && !triedLogin { // 401
			if dec.Meta.Code == "error" && dec.Meta.Msg == "api.err.LoginRequired" {
				if err := api.login(); err != nil {
					return err
				}
				triedLogin = true
				continue
			}
		}

		return fmt.Errorf("HTTP response %s", resp.Status)
	}
}

func (api *API) baseURL() string {
	return "https://" + api.auth.ControllerHost + ":8443"
}

func (api *API) login() error {
	req := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: api.auth.Username,
		Password: api.auth.Password,
	}
	return api.post("/api/login", &req, &json.RawMessage{}, reqOpts{
		referer: api.baseURL() + "/login",
	})
}

func (api *API) ListClients(site string) ([]Client, error) {
	var resp []Client
	if err := api.get("/api/s/"+site+"/stat/sta", &resp, reqOpts{}); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) ListAllClients(site string) ([]Client, error) {
	var resp []Client
	if err := api.get("/api/s/"+site+"/stat/alluser", &resp, reqOpts{}); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) BlockClient(site string, mac string) error {
	request := struct {
		Cmd string `json:"cmd"`
		Mac string `json:"mac"`
	}{
		Cmd: "block-sta",
		Mac: strings.ToLower(mac),
	}
	err := api.post("/api/s/"+site+"/cmd/stamgr", &request, &json.RawMessage{}, reqOpts{})
	if err != nil {
		return err
	}
	return nil
}

func (api *API) UnblockClient(site string, mac string) error {
	request := struct {
		Cmd string `json:"cmd"`
		Mac string `json:"mac"`
	}{
		Cmd: "unblock-sta", //only diff with above function
		Mac: strings.ToLower(mac),
	}
	err := api.post("/api/s/"+site+"/cmd/stamgr", &request, &json.RawMessage{}, reqOpts{})
	if err != nil {
		return err
	}
	return nil
}

func (api *API) ListWirelessNetworks(site string) ([]WirelessNetwork, error) {
	var resp []WirelessNetwork
	err := api.get("/api/s/"+site+"/list/wlanconf", &resp, reqOpts{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) EnableWirelessNetwork(site, id string, enable bool) error {
	req := struct {
		Enabled bool `json:"enabled"`
	}{enable}
	return api.post("/api/s/"+site+"/upd/wlanconf/"+id, &req, &json.RawMessage{}, reqOpts{})
}

type SwitchPort struct {
	ID         int    `json:"port_idx"`
	Name       string `json:"name"`
	POE        bool   `json:"port_poe"`
	PortConfID string `json:"portconf_id"`
	Up         bool   `json:"up"`
}

type UnifiDevice struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Wired    bool   `json:"is_wired"`

	MAC       string       `json:"mac"`
	IP        string       `json:"ip"`
	Model     string       `json:"model"` // usg/uap/usw
	Type      string       `json:"type"`
	PortTable []SwitchPort `json:"port_table"`
}

func (api *API) ListDevices(site string) ([]UnifiDevice, error) {
	var unifiDevices []UnifiDevice
	err := api.get("/api/s/"+site+"/stat/device", &unifiDevices, reqOpts{})
	if err != nil {
		return nil, err
	}
	return unifiDevices, nil
}
