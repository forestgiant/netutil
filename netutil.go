package netutil

import (
	"errors"
	"net"
)

// IsLocalhost takes an address and checks to see if
// it matches the local computers
func IsLocalhost(target string) bool {
	ip := net.ParseIP(target)
	if ip == nil {
		return false
	}

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
// If it can't get the local ip it returns 127.0.0.1
func LocalIP() net.IP {
	loopback := net.ParseIP("127.0.0.1")

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return loopback
	}

	for _, address := range addrs {
		// check the address type and make sure it's not loopback
		if ipnet, ok := address.(*net.IPNet); ok {
			if !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.To4()
				}
			}
		}
	}

	return loopback
}

// ConvertToLocalIP takes a loopback address and converts it to LocalIP
func ConvertToLocalIP(addr string) (string, error) {
	// Break a part addresses
	ip, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}

	// If local host convert to external ip
	if IsLocalhost(ip) || ip == "" {
		ip = LocalIP().String()
	} else {
		return "", errors.New("Address host must be localhost")
	}

	// Combine back
	return net.JoinHostPort(ip, port), nil
}
