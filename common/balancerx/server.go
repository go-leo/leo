package balancerx

import (
	"net/netip"
)

type Server struct {
	Scheme string
	Host   string
	Port   uint16
}

func (srv *Server) GetAddr() (netip.Addr, error) {
	return netip.ParseAddr(srv.Host)
}

func (srv *Server) GetAddrPort() (netip.AddrPort, error) {
	addr, err := netip.ParseAddr(srv.Host)
	if err != nil {
		return netip.AddrPort{}, err
	}
	return netip.AddrPortFrom(addr, srv.Port), nil
}
