package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"./unifi"
	prettyTime "github.com/andanhm/go-prettytime"
)

func main() {
	username := flag.String("u", "", "Ubiquiti controller username")
	password := flag.String("p", "", "Ubiquiti controller username")
	controllerHost := flag.String("c", "", "Ubiquiti controller host")
	mode := flag.String("m", "cli", "mode could be cli or daemon, default to cli")
	list := flag.String("list", "client", "list [client|network]")
	block := flag.String("block", "", "mac address of device to block")
	unblock := flag.String("unblock", "", "mac address of device to unblock")
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

	if *mode == "cli" {
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
			} else {
				listNetworks(api)
			}
		}
		// if *unblock == "" && *block == "" {
		// 	for {
		// 		fmt.Println("uypupo")
		// 	}
		// }
	} else {
		log.Println("kaboom")
	}

}

func listClients(api *unifi.API) {
	log.Printf("Fetching clients...")
	clients, err := api.ListClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}
	for _, client := range clients {
		name := client.Hostname
		if client.Name != "" {
			name = client.Name
		}
		wifi := "wifi"
		if client.Wired {
			wifi = "ethernet"
		}
		time := prettyTime.Format(client.LastSeen)
		fmt.Println("device", name, "on", wifi, "seen", time, "[", client.MAC, "] blocked", client.Blocked)
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
