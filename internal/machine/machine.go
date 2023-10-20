package machine

import (
	"net"
)

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
