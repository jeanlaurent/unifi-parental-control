package unifi

import (
	"encoding/json"
	"time"
)

type Client struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	Hostname     string `json:"hostname"`
	Wired        bool   `json:"is_wired"`
	Manufacturer string `json:"oui"`

	MAC string `json:"mac"`
	IP  string `json:"ip"`

	Blocked bool `json:"blocked"`

	LastSeen time.Time
}

func (c *Client) getName() string {
	if c.Name == "" {
		return c.Hostname
	}
	return c.Name
}

func (c *Client) UnmarshalJSON(data []byte) error {
	type Alias Client
	aux := struct {
		*Alias
		LastSeen int64 `json:"last_seen"`
		// TODO: do this for MAC, IP
	}{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}
