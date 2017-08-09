package rpc

import (
	"net"
	"time"
)

var (
	dialer = net.DialTimeout
)

// SetDialer set new a dialer for replace origin dailer.
// You could implemented a new dialer to connect with socks server
func SetDialer(newDialer func(network, address string, timeout time.Duration) (net.Conn, error)) {
	dialer = newDialer
}
