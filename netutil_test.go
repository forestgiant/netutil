package netutil

import (
	"net"
	"testing"
)

func TestIsLocalhost(t *testing.T) {
	tests := []struct {
		address string
		expect  bool
	}{
		{"127.0.0.1", true},
		{"", false},
		{"):*", false},
	}
	for _, test := range tests {
		result := IsLocalhost(test.address)
		if result != test.expect {
			t.Fatalf("Result was %t, expected %t", result, test.expect)
		}
	}
}

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
}

func TestCovertToLocalIP(t *testing.T) {
	tests := []struct {
		address string
		expect  bool
	}{
		{"127.0.0.1:9000", true},
		{":9000", true},
		{"", false},
		{"192.168.199.199:9000", false},
		{"):*", false},
	}
	for _, test := range tests {
		_, err := ConvertToLocalIP(test.address)
		result := err == nil
		if result != test.expect {
			t.Fatalf("Result was %t, expected %t", result, test.expect)
		}
	}
}
