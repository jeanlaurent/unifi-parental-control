package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// An AuthStore is an interface for loading and saving authentication information.
// See FileAuthStore for a file-based implementation.
type AuthStore interface {
	Load() (*Auth, error)
	Save(*Auth) error
}

// DefaultAuthFile is a default place to store authentication information.
// Pass this to FileAuthStore if an alternate path isn't required.
var DefaultAuthFile = filepath.Join(os.Getenv("HOME"), ".unifi-auth")

// FileAuthStore returns an AuthStore that stores authentication information in a named file.
func FileAuthStore(filename string) AuthStore {
	return fileAuthStore{filename}
}

type fileAuthStore struct {
	filename string
}

func (f fileAuthStore) Load() (*Auth, error) {
	// Security check.
	fi, err := os.Stat(f.filename)
	if err != nil {
		return nil, err
	}
	if fi.Mode()&0077 != 0 {
		return nil, fmt.Errorf("security check failed on %s: mode is %04o; it should not be accessible by group/other", f.filename, fi.Mode())
	}

	raw, err := ioutil.ReadFile(f.filename)
	if err != nil {
		return nil, err
	}
	auth := new(Auth)
	if err := json.Unmarshal(raw, auth); err != nil {
		return nil, fmt.Errorf("bad auth file %s: %v", f.filename, err)
	}
	return auth, nil
}

func (f fileAuthStore) Save(auth *Auth) error {
	raw, err := json.Marshal(auth)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f.filename, raw, 0600)
}
