package internal

import (
	"errors"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/netx/addrx"
	"net"
	"strconv"
)

func GlobalUnicastAddr(address net.Addr) (net.IP, int, error) {
	host, port, err := net.SplitHostPort(address.String())
	if err != nil {
		return nil, 0, err
	}
	ip := net.ParseIP(host)
	if addrx.IsGlobalUnicastIP(ip) {
		return ip, errorx.Ignore(strconv.Atoi(port)), nil
	}
	if !ip.IsUnspecified() {
		return nil, 0, errors.New("failed to get global unicast ip")
	}
	ip, err = addrx.GlobalUnicastIP()
	if err != nil {
		return nil, 0, err
	}
	return ip, errorx.Ignore(strconv.Atoi(port)), err
}
