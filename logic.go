package main

import (
	"log"
	"time"

	"github.com/jeanlaurent/unifi-parental-control/unifi"
)

func blockOrUnblockDevices(api *unifi.API, deviceConfig config) {
	clients, err := api.ListAllClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}

	for _, blockRange := range deviceConfig.BlockRanges {

		client := findClientByMac(blockRange.Mac, clients)
		if client == nil {
			log.Println("could not find client")
			return
		}

		now := time.Now()

		if blockRange.BlockTimeStamp().After(now) && blockRange.UnblockTimeStamp().Before(now) {
			// block device
			if client.Blocked {
				log.Println("Client ", blockRange.Mac, " already blocked, skipping")
			} else {
				log.Println("blocking ", blockRange.Mac)
				api.BlockClient("default", blockRange.Mac)
			}
		} else if client.Blocked { // unblock device
			log.Println("unblocking ", blockRange.Mac)
			api.UnblockClient("default", blockRange.Mac)
		} else {
			log.Println("Client ", blockRange.Mac, " already unblocked, skipping")
		}
		log.Println(blockRange)
	}
}

func findClientByMac(mac string, clients []unifi.Client) *unifi.Client {
	for _, client := range clients {
		if client.MAC == mac {
			return &client
		}
	}
	return nil
}
