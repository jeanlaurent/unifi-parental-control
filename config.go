package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	blockRanges []BlockRange
}

func parseConfig(configAsByte []byte) config {
	config := config{blockRanges: []BlockRange{}}
	err := yaml.Unmarshal(configAsByte, &config.blockRanges)
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
