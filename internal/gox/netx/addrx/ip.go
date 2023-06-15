package addrx

import (
	"fmt"
	"net"
)

// ExtractIP extract IP from net.Addr
func ExtractIP(addr net.Addr) net.IP {
	switch v := addr.(type) {
	case *net.IPAddr:
		return v.IP
	case *net.IPNet:
		return v.IP
	case *net.TCPAddr:
		return v.IP
	case *net.UDPAddr:
		return v.IP
	default:
		return nil
	}
}

// GlobalUnicastIP get a global unicast IP address
func GlobalUnicastIP() (net.IP, error) {
	ips := IPs()
	for _, ip := range ips {
		if ip.IsUnspecified() {
			continue
		}
		if ip.IsLoopback() {
			continue
		}
		if ip.IsLinkLocalUnicast() {
			continue
		}

		if ip.IsInterfaceLocalMulticast() {
			continue
		}
		if ip.IsLinkLocalMulticast() {
			continue
		}
		if ip.IsMulticast() {
			continue
		}
		if ip.IsGlobalUnicast() {
			return ip, nil
		}
	}
	return nil, fmt.Errorf("no found global unicast IP")
}

// GlobalUnicastIPString get a global unicast IP address string
func GlobalUnicastIPString() (string, error) {
	ip, err := GlobalUnicastIP()
	if err != nil {
		return "", err
	}
	return ip.String(), nil
}

// IPs get all IP addresses
func IPs() []net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	var ips []net.IP
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ip := ExtractIP(addr)
			if len(ip) == 0 {
				continue
			}
			ips = append(ips, ip)
		}
	}
	return ips
}
