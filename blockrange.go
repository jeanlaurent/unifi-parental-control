package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// BlockRange represent a range to be blocked
type BlockRange struct {
	Mac         string `yaml:"mac"`
	BlockTime   string `yaml:"blockTime"`
	UnblockTime string `yaml:"unblockTime"`
}

// BlockTimeStamp translate a blocked timestamp in to time.Time
func (d *BlockRange) BlockTimeStamp() time.Time {
	time, err := crappyParseTime(d.BlockTime)
	if err != nil {
		fmt.Println(err)
		// KABBOOM
	}
	return time
}

// UnblockTimeStamp translate a unblock timestamo in to time.Time
func (d *BlockRange) UnblockTimeStamp() time.Time {
	time, err := crappyParseTime(d.UnblockTime)
	if err != nil {
		fmt.Println(err)
		// KABOOM
	}
	return time
}

func crappyParseTime(timeAsString string) (time.Time, error) {
	// this is mega crap, this must be checked at marshalling time see https://stackoverflow.com/questions/49530395/custom-unmarshalyaml-how-to-implement-unmarshaler-interface-on-custom-type
	now := time.Now()
	parts := strings.Split(timeAsString, ":")
	if len(parts) != 2 {
		return now, fmt.Errorf("error while parsing date, %v", parts)
	}
	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return now, err
	}
	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return now, err
	}
	return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local), nil
}