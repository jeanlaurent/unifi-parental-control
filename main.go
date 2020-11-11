package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	prettyTime "github.com/andanhm/go-prettytime"
	"github.com/jeanlaurent/unifi-parental-control/unifi"
	"github.com/olekukonko/tablewriter"
)

func main() {
	username := flag.String("u", "", "Unifi controller username")
	password := flag.String("p", "", "Unifi controller password")
	controllerHost := flag.String("c", "", "Unifi controller host")

	list := flag.String("list", "", "List [client|network|all|device]")

	block := flag.String("block", "", "Mac address or group of device to block")
	unblock := flag.String("unblock", "", "Mac address or group of device to unblock")
	config := flag.String("config", "", "Configuration file holding group definition")

	port := flag.Int("port", 0, "Port to allow poe on, must be used in conjunction with either -poeon or -poweroff")
	poeon := flag.String("poeon", "", "DeviceID of switch to enable poe on, to be use in conjunction of -port")
	poeoff := flag.String("poeoff", "", "DeviceID of switch to disable poe on, to be use in conjunction of -port")

	displayDateTime := flag.Bool("time", false, "Prefix output with datetime default is false")

	flag.Parse()

	if !*displayDateTime {
		log.SetFlags(0)
	}

	if *username == "" {
		log.Println("Missing username")
		log.Println("usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *password == "" {
		log.Println("Missing password")
		log.Println("usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *controllerHost == "" {
		log.Println("Missing controller host")
		log.Println("usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	api, err := unifi.BuildAPI(*username, *password, *controllerHost)
	if err != nil {
		log.Fatalf("buildApi Error: %v", err)
	}

	if *block != "" {
		deviceConfig := emptyConfig()
		log.Println("blocking", *block)
		if *config != "" {
			deviceConfig = readConfigFromDisk(*config)
		}
		group := deviceConfig.Groups[*block]
		if group == nil {
			group = []string{*block}
		}
		for _, mac := range group {
			err = api.BlockClient("default", mac)
			if err != nil {
				log.Println("blocking client:", err)
			}
			log.Println(mac, "will be shortly blocked")
		}
	}

	if *unblock != "" {
		deviceConfig := emptyConfig()
		log.Println("unblocking", *unblock)
		if *config != "" {
			deviceConfig = readConfigFromDisk(*config)
		}
		group := deviceConfig.Groups[*unblock]
		if group == nil {
			group = []string{*unblock}
		}
		for _, mac := range group {
			err = api.UnblockClient("default", mac)
			if err != nil {
				log.Println("unblocking client:", err)
			}
			log.Println(mac, "will be shortly unblocked")
		}
	}

	if *poeon != "" {
		if *port == 0 {
			log.Println("Please provide a port number")
			log.Println("usage:")
			flag.PrintDefaults()
			os.Exit(1)
		}
		devices, err := api.ListDevices("default")
		if err != nil {
			log.Fatalf("could not list devices: %v", err)
		}
		for _, device := range devices {
			if device.ID == *poeon {
				log.Println("Enabling POE on port", *port, " of switch ", device.Name)
				api.EnablePortPOE("default", device.ID, *port, true)
			}
		}
	}

	if *poeoff != "" {
		if *port == 0 {
			log.Println("Please provide a port number")
			log.Println("usage:")
			flag.PrintDefaults()
			os.Exit(1)
		}
		devices, err := api.ListDevices("default")
		if err != nil {
			log.Fatalf("could not list devices: %v", err)
		}
		for _, device := range devices {
			if device.ID == *poeoff {
				log.Println("Disabling POE on port", *port, " of switch ", device.Name)
				api.EnablePortPOE("default", device.ID, *port, false)
			}
		}
	}

	if *list != "" {
		if *list == "client" {
			listClients(api)
		} else if *list == "all" {
			listAllClients(api)
		} else if *list == "network" {
			listNetworks(api)
		} else {
			listUnifiDevices(api)
		}
	}

}

func listAllClients(api *unifi.API) {
	clients, err := api.ListAllClients("default")
	if err != nil {
		log.Fatalf("Listing all clients: %v", err)
	}
	displayClients(clients)
}

func listClients(api *unifi.API) {
	clients, err := api.ListClients("default")
	if err != nil {
		log.Fatalf("Listing w active clients: %v", err)
	}
	displayClients(clients)
}

func displayClients(clients []unifi.Client) {
	data := [][]string{}
	for _, client := range clients {
		name := "unknow"
		if client.Hostname != "" {
			name = client.Hostname
		}
		if client.Name != "" {
			name = client.Name
		}
		wifi := "wifi"
		if client.Wired {
			wifi = "ethernet"
		}
		data = append(data, []string{name, client.Manufacturer, wifi, prettyTime.Format(client.LastSeen), client.MAC, strconv.FormatBool(client.Blocked)})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Manufacturer", "Wifi/Wired", "Last Seen", "Mac", "Blocked"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}

func listUnifiDevices(api *unifi.API) {
	unifiDevices, err := api.ListDevices("default")
	if err != nil {
		log.Fatalf("Fetching unifi devices: %v", err)
	}

	data := [][]string{}
	for _, device := range unifiDevices {
		data = append(data, []string{device.Name, device.Type, device.Model, device.ID, device.MAC, "", "", "", ""})
		if len(device.PortTable) > 0 {
			for _, port := range device.PortTable {
				data = append(data, []string{"", "", "", "", "", port.Name, strconv.FormatBool(port.POE), strconv.FormatBool(port.Up), port.PortConfID})
			}
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "Model", "ID", "Mac", "Port", "PortHasPOE", "PortIsUP", "PortConf"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}

func listNetworks(api *unifi.API) {
	wlans, err := api.ListWirelessNetworks("default")
	if err != nil {
		log.Fatalf("Fetching wireless networks: %v", err)
	}
	data := [][]string{}
	for _, wlan := range wlans {
		data = append(data, []string{wlan.Name, wlan.Security, strconv.FormatBool(wlan.Enabled), strconv.FormatBool(wlan.Guest)})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Security", "Enabled", "Guest"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}
