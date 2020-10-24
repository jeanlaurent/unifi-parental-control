package main

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jeanlaurent/unifi-parental-control/unifi"
)

func startPollingScheduler(api *unifi.API, deviceConfig config) {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(1).Minutes().Do(func() {
		blockOrUnblockDevices(api, deviceConfig)
	})
	<-scheduler.StartAsync()
}
