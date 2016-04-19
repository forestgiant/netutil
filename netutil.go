package netutil

import "net"

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
