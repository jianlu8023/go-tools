package ip4

import (
	"net"
	"strings"

	"github.com/jianlu8023/go-tools/v2/internal/netutil"
)

// GetIPv4Addr
// @return string:
func GetIPv4Addr() string {
	// 尝试使用UDP连接获取IP地址
	dial, err := net.Dial("udp", "8.8.8.8:53")
	if err == nil {
		defer dial.Close()
		localAddr := dial.LocalAddr().(*net.UDPAddr)
		ip := strings.Split(localAddr.String(), ":")[0]
		return ip
	}

	// 备用方案：当无法连接到8.8.8.8:53时，从网络接口获取IP地址
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		// 跳过非活动接口和回环接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addresses, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addresses {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				// 返回找到的第一个非回环IPv4地址
				return ipNet.IP.String()
			}
		}
	}

	return ""
}

// GetIPv4Addrs
// @return map[string]string:
func GetIPv4Addrs() map[string]string {
	IPv4Map := make(map[string]string)

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagBroadcast == net.FlagBroadcast {
			addresses, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addresses {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					IPv4Map[iface.Name] = ipNet.IP.String()
				}
			}
		}
	}
	return IPv4Map
}

// IPv4Readable 检查ip地址是否可达
func IPv4Readable(ip string) bool {

	if err := netutil.Reachable(ip); err != nil {
		return false
	} else {
		return true
	}
}
