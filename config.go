package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	BlockRanges []BlockRange        `yaml:"blockRange"`
	Groups      map[string][]string `yaml:"group"`
}

func parseConfig(configAsByte []byte) config {
	config := config{BlockRanges: []BlockRange{}, Groups: map[string][]string{}}
	err := yaml.Unmarshal(configAsByte, &config)
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
