package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	prettyTime "github.com/andanhm/go-prettytime"
	"github.com/jeanlaurent/unifi-parental-control/unifi"
)

func main() {
	username := flag.String("u", "", "Unifi controller username")
	password := flag.String("p", "", "Unifi controller username")
	controllerHost := flag.String("c", "", "Unifi controller host")
	list := flag.String("list", "", "list [client|network|all]")
	block := flag.String("block", "", "mac address of device to block")
	unblock := flag.String("unblock", "", "mac address of device to unblock")
	config := flag.String("config", "", "config file")
	flag.Parse()

	if *username == "" {
		fmt.Println("Missing username")
		fmt.Println("usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *password == "" {
		fmt.Println("Missing password")
		fmt.Println("usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *controllerHost == "" {
		fmt.Println("Missing controller host")
		fmt.Println("usage:")
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
		fmt.Println(*block, "should be blocked")
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
		fmt.Println(*unblock, "should be unblocked")
		listClients(api)
	}

	if *list != "" {
		if *list == "client" {
			listClients(api)
		} else if *list == "all" {
			listAllClients(api)
		} else {
			listNetworks(api)
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
		fmt.Println("device", name, "(", client.Manufacturer, ")", "on", wifi, "seen", time, "[", client.MAC, "] blocked", client.Blocked)
	}
}

func listNetworks(api *unifi.API) {
	log.Printf("Fetching wireless networks...")
	wlans, err := api.ListWirelessNetworks("default")
	if err != nil {
		log.Fatalf("Fetching wireless networks: %v", err)
	}
	for _, wlan := range wlans {
		fmt.Printf("%+v\n", wlan)
	}
}
