package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMacInConfig(t *testing.T) {
	configYaml := `- mac: foobar
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.blockRanges[0].Mac)
}

func TestReadBlockTimeInConfig(t *testing.T) {
	configYaml := `- blockTime: foobar
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.blockRanges[0].BlockTime)
}

func TestReadUnblockTimeInConfig(t *testing.T) {
	configYaml := `- unblockTime: foobar
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.blockRanges[0].UnblockTime)
}

func TestReadConfigFromDisk(t *testing.T) {
	config := readConfigFromDisk("testdata/devices.yaml")
	assert.Equal(t, "00:22:d7:52:8b:74", config.blockRanges[0].Mac)
	assert.Equal(t, "18:00", config.blockRanges[0].BlockTime)
	assert.Equal(t, "19:00", config.blockRanges[0].UnblockTime)
	assert.Equal(t, "f0:4f:7c:29:21:6c", config.blockRanges[1].Mac)
	assert.Equal(t, "17:00", config.blockRanges[1].BlockTime)
	assert.Equal(t, "17:30", config.blockRanges[1].UnblockTime)
}
