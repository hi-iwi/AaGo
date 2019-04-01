package util

import (
	"net"
)

// LocalAddress 223.5.5.5 is aliyun public DNS. Google public DNS 8.8.8.8 is not connectable in China.
func LocalAddress() (addr string, err error) {
	conn, err := net.Dial("udp", "223.5.5.5:80")
	if err != nil {
		return
	}
	defer conn.Close()
	addr = conn.LocalAddr().(*net.UDPAddr).IP.String()
	return
}

// LocalAddresses For multiple network cards (virtual network cards included, e.g. docker)
func LocalAddresses() (addresses []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addresses = append(addresses, ipnet.IP.String())
			}
		}
	}
	return
}
