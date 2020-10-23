package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jeanlaurent/unifi-parental-control/unifi"
)

func startPollingScheduler(api *unifi.API, deviceConfig Device) {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(1).Seconds().Do(func() {
		blockOrUnblockDevices(api, deviceConfig)
	})
	<-scheduler.StartAsync()
}

func blockOrUnblockDevices(api *unifi.API, deviceConfig Device) {
	clients, err := api.ListAllClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}

	now := time.Now()

	client := findClientByMac(deviceConfig.Mac, clients)
	if client == nil {
		fmt.Println("could not find client")
		return
	}

	if deviceConfig.BlockTimeStamp().After(now) && deviceConfig.UnblockTimeStamp().Before(now) {
		// block device
		if client.Blocked {
			fmt.Println("Client ", deviceConfig.Mac, " already blocked, skipping")
		} else {
			fmt.Println("blocking ", deviceConfig.Mac)
			api.BlockClient("default", deviceConfig.Mac)
		}
	} else if client.Blocked { // unblock device
		fmt.Println("unblocking ", deviceConfig.Mac)
		api.UnblockClient("default", deviceConfig.Mac)
	} else {
		fmt.Println("Client ", deviceConfig.Mac, " already unblocked, skipping")
	}
	fmt.Println(deviceConfig)
}

func findClientByMac(mac string, clients []unifi.Client) *unifi.Client {
	for _, client := range clients {
		if client.MAC == mac {
			return &client
		}
	}
	return nil
}
