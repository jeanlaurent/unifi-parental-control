package main

import (
	"flag"
	"log"
	"os"

	prettyTime "github.com/andanhm/go-prettytime"
	"github.com/jeanlaurent/unifi-parental-control/unifi"
)

func main() {
	username := flag.String("u", "", "Unifi controller username")
	password := flag.String("p", "", "Unifi controller username")
	controllerHost := flag.String("c", "", "Unifi controller host")
	list := flag.String("list", "", "list [client|network|all|device]")
	block := flag.String("block", "", "mac address of device to block")
	unblock := flag.String("unblock", "", "mac address of device to unblock")
	port := flag.Int("port", 0, "Port to allow poe on, must be used in conjunction with either -poeon or -poweroff")
	poeon := flag.String("poeon", "", "DeviceID of switch to enable poe on, to be use in conjunction of -port")
	poeoff := flag.String("poeoff", "", "DeviceID of switch to disable poe on, to be use in conjunction of -port")
	config := flag.String("config", "", "config file to launch in cron mode")

	flag.Parse()

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
		listClients(api)
		log.Println("=========================")
		log.Println("blocking", *block)
		err = api.BlockClient("default", *block)
		if err != nil {
			log.Fatalf("blocking client: %v", err)
		}
		log.Println(*block, "should be blocked")
		listClients(api)
	}

	if *unblock != "" {
		listClients(api)
		log.Println("=========================")
		log.Println("unblocking", *unblock)
		err = api.UnblockClient("default", *unblock)
		if err != nil {
			log.Fatalf("unblocking client: %v", err)
		}
		log.Println(*unblock, "should be unblocked")
		listClients(api)
	}

	if *poeon != "" {
		if *port == 0 {
			log.Fatal("Please provide a port number")
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
			log.Fatal("Please provide a port number")
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

	if *config != "" {
		deviceConfig := readConfigFromDisk(*config)
		startPollingScheduler(api, deviceConfig)
	}

}

func listAllClients(api *unifi.API) {
	log.Printf("Fetching clients...")
	clients, err := api.ListAllClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}
	displayClients(clients)
}

func listClients(api *unifi.API) {
	log.Printf("Fetching clients...")
	clients, err := api.ListClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}
	displayClients(clients)
}

func displayClients(clients []unifi.Client) {
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
		time := prettyTime.Format(client.LastSeen)
		log.Println("device", name, "(", client.Manufacturer, ")", "on", wifi, "seen", time, "[", client.MAC, "] blocked", client.Blocked)
	}
}

func listUnifiDevices(api *unifi.API) {
	log.Printf("Fetching unifi devices...")
	unifiDevices, err := api.ListDevices("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}

	for _, device := range unifiDevices {
		log.Println("device", device.Name, "(", device.Type, device.Model, ")", "ID:", device.ID, "[", device.MAC, "]")
		if len(device.PortTable) > 0 {
			for _, port := range device.PortTable {
				log.Println("\t", port.Name, "HasPoe:", port.POE, "up:", port.Up, "PortConf:", port.PortConfID)
			}

		}
	}
}

func listNetworks(api *unifi.API) {
	log.Printf("Fetching wireless networks...")
	wlans, err := api.ListWirelessNetworks("default")
	if err != nil {
		log.Fatalf("Fetching wireless networks: %v", err)
	}
	for _, wlan := range wlans {
		log.Printf("%+v\n", wlan)
	}
}
