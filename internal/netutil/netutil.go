package netutil

import (
	"net"
)

// ContainsIPv4 检查网络接口是否包含指定的IPv4地址
func ContainsIPv4(iface net.Interface, target net.IP) bool {
	addrs, _ := iface.Addrs()

	for _, v := range addrs {
		addr := v.(*net.IPNet)
		if addr.Contains(target) {
			return true
		}
	}
	return false
}
