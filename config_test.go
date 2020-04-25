package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMacInConfig(t *testing.T) {
	configYaml := `
mac: foobar
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.Mac)
}

func TestReadBlockTimeInConfig(t *testing.T) {
	configYaml := `
blockTime: foobar
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.BlockTime)
}

func TestReadUnblockTimeInConfig(t *testing.T) {
	configYaml := `
unblockTime: foobar
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.UnblockTime)
}

func TestReadConfigFromDisk(t *testing.T) {

	config := readConfigFromDisk("testdata/device.yaml")
	assert.Equal(t, "00:22:d7:52:8b:74", config.Mac)
	assert.Equal(t, "18:00", config.BlockTime)
	assert.Equal(t, "19:00", config.UnblockTime)
}
