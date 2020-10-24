package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jeanlaurent/unifi-parental-control/unifi"
)

func startPollingScheduler(api *unifi.API, deviceConfig config) {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(1).Seconds().Do(func() {
		blockOrUnblockDevices(api, deviceConfig)
	})
	<-scheduler.StartAsync()
}

func blockOrUnblockDevices(api *unifi.API, deviceConfig config) {
	clients, err := api.ListAllClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}

	for _, blockRange := range deviceConfig.blockRange {

		client := findClientByMac(blockRange.Mac, clients)
		if client == nil {
			fmt.Println("could not find client")
			return
		}

		now := time.Now()

		if blockRange.BlockTimeStamp().After(now) && blockRange.UnblockTimeStamp().Before(now) {
			// block device
			if client.Blocked {
				fmt.Println("Client ", blockRange.Mac, " already blocked, skipping")
			} else {
				fmt.Println("blocking ", blockRange.Mac)
				api.BlockClient("default", blockRange.Mac)
			}
		} else if client.Blocked { // unblock device
			fmt.Println("unblocking ", blockRange.Mac)
			api.UnblockClient("default", blockRange.Mac)
		} else {
			fmt.Println("Client ", blockRange.Mac, " already unblocked, skipping")
		}
		fmt.Println(blockRange)
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
