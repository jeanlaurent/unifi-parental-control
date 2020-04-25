package unifi

type WirelessNetwork struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`

	Security string `json:"security"`
	WPAMode  string `json:"wpa_mode"`

	Guest bool `json:"is_guest,omitempty"`

	// TODO: other fields
}
