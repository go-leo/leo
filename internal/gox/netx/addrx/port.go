package addrx

import (
	"net"
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
