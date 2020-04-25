package main

type Device struct {
	Mac         string `yaml:"mac"`
	BlockTime   string `yaml:"blockTime"`
	UnblockTime string `yaml:"unblockTime"`
}
