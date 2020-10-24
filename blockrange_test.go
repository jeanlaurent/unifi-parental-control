package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadBlockTime(t *testing.T) {
	device := BlockRange{Mac: "foobar", BlockTime: "18:00"}
	now := time.Now()
	blockTime := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, time.Local)
	assert.True(t, blockTime.Equal(device.BlockTimeStamp()))
}

func TestReadUnBlockTime(t *testing.T) {
	device := BlockRange{Mac: "barqix", UnblockTime: "5:07"}
	now := time.Now()
	unblockTime := time.Date(now.Year(), now.Month(), now.Day(), 5, 7, 0, 0, time.Local)
	assert.True(t, unblockTime.Equal(device.UnblockTimeStamp()))
}
