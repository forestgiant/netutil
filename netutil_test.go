package netutil

import "testing"

func TestLocalIP(t *testing.T) {
	e, err := LocalIP()
	if err != nil {
		t.Error("LocalIP errored:", err)
	}

	if !IsLocalhost(e) {
		t.Error("Should be local ip. IsLocalhost is false")
	}
}
