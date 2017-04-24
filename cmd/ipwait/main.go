package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/forestgiant/netutil"
)

func main() {
	var (
		delayPtr   = flag.Int("delay", 0, "Delay (milliseconds) inbetween checking the interface")
		timeoutPtr = flag.Int("timeout", 20, "Max timeout(seconds)")
	)
	flag.Parse()

	// Block till we find a non loopback ipv6
	var ip net.IP
	timeout := time.Second * time.Duration(*timeoutPtr)
	startTime := time.Now()
	for ip = netutil.LocalIPv6(); ip.IsLoopback() && time.Since(startTime) < timeout; ip = netutil.LocalIPv6() {
		time.Sleep(time.Duration(*delayPtr) * time.Millisecond)
	}

	if !ip.IsLoopback() {
		fmt.Println(ip)
	}
}
