package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	blockRange []BlockRange
}

func parseConfig(configAsByte []byte) config {
	config := config{blockRange: []BlockRange{}}
	err := yaml.Unmarshal(configAsByte, &config.blockRange)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
}

func readConfigFromDisk(filepath string) config {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Can't read file #%v ", err)
	}
	return parseConfig(file)
}
