package addrx

import (
	"net"
	"strconv"
)

// ExtractPort extract Port from net.Addr
func ExtractPort(addr net.Addr) int {
	switch v := addr.(type) {
	case *net.TCPAddr:
		return v.Port
	case *net.UDPAddr:
		return v.Port
	default:
		return 0
	}
}

// PickFreePort automatically chose a free port and return it
func PickFreePort() (int, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer l.Close()
	addr := l.Addr().String()
	_, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return 0, err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, err
	}
	return port, nil
}
