package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	Groups map[string][]string `yaml:"group"`
}

func parseConfig(configAsByte []byte) config {
	config := emptyConfig()
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

func emptyConfig() config {
	return config{Groups: map[string][]string{}}
}
