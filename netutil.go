package netutil

import (
	"errors"
	"net"
)

// IsLocalhost takes an address and checks to see if
// it matches the local computers
func IsLocalhost(target string) bool {
	ip := net.ParseIP(target)

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if ip.Equal(ipnet.IP) {
				return true
			}
		}
	}

	return false
}

// LocalIP return the network address of the computer
func LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// check the address type and make sure it's not loopback
		if ipnet, ok := address.(*net.IPNet); ok {
			if !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.To4().String(), nil
				}
			}
		}
	}

	return "", errors.New("Couldn't get local IP")
}
