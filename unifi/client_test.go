package unifi

import (
	"testing"

	"gotest.tools/assert"
)

func TestNameisHostname(t *testing.T) {
	client := Client{Name: "", Hostname: "hostname"}

	assert.Equal(t, "hostname", client.getName())
}

func TestNameisName(t *testing.T) {
	client := Client{Name: "name", Hostname: "polka"}

	assert.Equal(t, "name", client.getName())
}
