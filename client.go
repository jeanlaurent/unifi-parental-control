package main

import (
	"encoding/json"
	"time"
)

type Client struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Wired    bool   `json:"is_wired"`

	MAC string `json:"mac"`
	IP  string `json:"ip"`

	Blocked bool `json:"blocked"`

	LastSeen time.Time

	// TODO: other fields
}

func (c *Client) UnmarshalJSON(data []byte) error {
	type Alias Client
	aux := struct {
		*Alias

		LastSeen int64 `json:"last_seen"`
		// TODO: do this for MAC, IP
	}{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}
