package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMacInConfig(t *testing.T) {
	configYaml := `blockRange:
  - mac: foobar`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.BlockRanges[0].Mac)
}

func TestReadBlockTimeInConfig(t *testing.T) {
	configYaml := `blockRange:
  - blockTime: foobar
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.BlockRanges[0].BlockTime)
}

func TestReadUnblockTimeInConfig(t *testing.T) {
	configYaml := `blockRange:
  - unblockTime: foobar
`
	config := parseConfig([]byte(configYaml))
	assert.Equal(t, "foobar", config.BlockRanges[0].UnblockTime)
}

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
	assert.Equal(t, "00:22:d7:52:8b:74", config.BlockRanges[0].Mac)
	assert.Equal(t, "18:00", config.BlockRanges[0].BlockTime)
	assert.Equal(t, "19:00", config.BlockRanges[0].UnblockTime)
	assert.Equal(t, "f0:4f:7c:29:21:6c", config.BlockRanges[1].Mac)
	assert.Equal(t, "17:00", config.BlockRanges[1].BlockTime)
	assert.Equal(t, "17:30", config.BlockRanges[1].UnblockTime)
	assert.Equal(t, "00:01:02:03:04:05", config.Groups["clement"][0])
	assert.Equal(t, "00:01:02:03:04:06", config.Groups["clement"][1])
	assert.Equal(t, "00:01:02:03:04:07", config.Groups["marin"][0])
	assert.Equal(t, "00:01:02:03:04:08", config.Groups["marin"][1])
}
