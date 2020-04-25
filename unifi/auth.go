package unifi

import "net/http"

// Auth holds the authentication information for accessing a UniFi controller.
type Auth struct {
	Username       string
	Password       string
	ControllerHost string
	Cookies        []*http.Cookie
}
