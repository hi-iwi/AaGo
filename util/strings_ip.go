package util

import (
	"net"
	"net/http"
	"strings"
)

// GetSelfDefinedHeader
func GetSelfDefinedHeader(h http.Header, k string) string {
	var v string
	if v = h.Get(k); v != "" {
		return v
	}
	if v = h.Get(strings.ToUpper(k)); v != "" {
		return v
	}
	if v = h.Get(strings.ToLower(k)); v != "" {
		return v
	}
	return ""
}
func splitHost(addr string) string {
	host, _, err := net.SplitHostPort(addr)
	if err == nil {
		return host
	}
	return addr
}

// RemoteIp 客户端真实IP地址
func RemoteIp(r *http.Request) string {
	var addr string
	if addr = GetSelfDefinedHeader(r.Header, "X-Real-Ip"); addr != "" {
		return splitHost(addr)
	}
	if addr = GetSelfDefinedHeader(r.Header, "X-Forwarded-For"); addr != "" {
		return splitHost(addr)
	}
	// may be the last proxy  有可能是最后一层代理IP地址
	return splitHost(r.RemoteAddr)
}

// LocalAddress 223.5.5.5 is aliyun public DNS. Google public DNS 8.8.8.8 is not connectable in China.
func LocalIP() (addr string, err error) {
	conn, err := net.Dial("udp", "223.5.5.5:80")
	if err != nil {
		return
	}
	defer conn.Close()
	addr = conn.LocalAddr().(*net.UDPAddr).IP.String()
	return
}

// LocalAddresses For multiple network cards (virtual network cards included, e.g. docker)
func LocalIPs() (addresses []string, err error) {
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
