package netutil

import (
	"fmt"
	"net"
	"testing"
)

func TestLocalIP(t *testing.T) {
	l := LocalIP()
	_, err := net.InterfaceAddrs()
	if err != nil {
		// if there was an error with the interface
		// then it should be loopback
		if !l.IsLoopback() {
			t.Error("LocalIP should be loopback")
		}
	}

	// Make sure it is localhost
	if !IsLocalhost(l.String()) {
		t.Error("LocalIP should be localhost")
	}

	fmt.Println(l)
}
