package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	prettyTime "github.com/andanhm/go-prettytime"
)

func main() {
	username := flag.String("u", "", "Ubiquiti controller username")
	password := flag.String("p", "", "Ubiquiti controller username")
	controllerHost := flag.String("c", "", "Ubiquiti controller host")
	mode := flag.String("m", "cli", "mode could be cli or server, default to cli")
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
	api, err := buildAPI(*username, *password, *controllerHost)
	if err != nil {
		log.Fatalf("unifi.NewClient: %v", err)
	}
	if *mode == "cli" {

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
			fmt.Println("device", name, "on", wifi, "seen", time)
		}
		log.Println("=========================")
		log.Printf("Fetching wireless networks...")
		wlans, err := api.ListWirelessNetworks("default")
		if err != nil {
			log.Fatalf("Fetching wireless networks: %v", err)
		}
		for _, wlan := range wlans {
			fmt.Printf("%+v\n", wlan)
		}
		log.Println("=========================")
		log.Println("unblocking 48:ba:4e:87:92:2f")
		err = api.UnblockClient("default", "48:ba:4e:87:92:2f")
		if err != nil {
			log.Fatalf("Fetching wireless networks: %v", err)
		}
		fmt.Println("48:ba:4e:87:92:2f should be unblocked")
	} else {
		start(api)
	}

}
