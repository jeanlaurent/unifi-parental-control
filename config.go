package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func parseConfig(configAsByte []byte) Device {
	readDevice := Device{}
	err := yaml.Unmarshal(configAsByte, &readDevice)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return readDevice
}

func readConfigFromDisk(filepath string) Device {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Can't read file #%v ", err)
	}
	return parseConfig(file)
}
