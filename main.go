package main

import (
	"fmt"
	"log"
)

func main() {
	// this is pure garbage
	// we should ask login/password everytime and keep cookies in memory for the time being
	api, err := NewAPI(FileAuthStore(DefaultAuthFile))
	if err != nil {
		log.Fatalf("unifi.NewClient: %v", err)
	}
	defer func() {
		if err := api.WriteConfig(); err != nil {
			log.Printf("api.WriteConfig: %v", err)
		}
	}()

	log.Printf("Fetching clients...")
	clients, err := api.ListClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}
	for _, client := range clients {
		fmt.Printf("%+v\n", client)
	}

	log.Printf("Fetching wireless networks...")
	wlans, err := api.ListWirelessNetworks("default")
	if err != nil {
		log.Fatalf("Fetching wireless networks: %v", err)
	}
	for _, wlan := range wlans {
		fmt.Printf("%+v\n", wlan)
	}
	log.Println("blocking 48:ba:4e:87:92:2f")
	err = api.UnblockClient("default", "48:ba:4e:87:92:2f")
	if err != nil {
		log.Fatalf("Fetching wireless networks: %v", err)
	}
	fmt.Println("48:ba:4e:87:92:2f should be unblocked")
}
