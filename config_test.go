package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleGroupInConfig(t *testing.T) {
	configYaml := `group:
  kid1: 
    - 00:22:d7:52:8b:74
    - 00:22:d7:52:8b:90
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "00:22:d7:52:8b:74", config.Groups["kid1"][0])
	assert.Equal(t, "00:22:d7:52:8b:90", config.Groups["kid1"][1])
}

func TestMultipleGroupInConfig(t *testing.T) {
	configYaml := `group:
  kid1: 
    - 00:22:d7:52:8b:74
  kid2:
    - 00:22:d7:52:8b:90
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "00:22:d7:52:8b:74", config.Groups["kid1"][0])
	assert.Equal(t, "00:22:d7:52:8b:90", config.Groups["kid2"][0])
}

func TestReadConfigFromDisk(t *testing.T) {
	config := readConfigFromDisk("testdata/devices.yaml")
	assert.Equal(t, "00:01:02:03:04:05", config.Groups["clement"][0])
	assert.Equal(t, "00:01:02:03:04:06", config.Groups["clement"][1])
	assert.Equal(t, "00:01:02:03:04:07", config.Groups["marin"][0])
	assert.Equal(t, "00:01:02:03:04:08", config.Groups["marin"][1])
}
