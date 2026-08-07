package ip6

import (
	"net"
	"time"

	"github.com/jianlu8023/go-tools/v2/internal/netutil"
)

// GetIPv6Addr
// @return string:
func GetIPv6Addr() string {
	// 尝试使用UDP连接获取IP地址
	conn, err := net.DialTimeout("udp", "8.8.8.8:53", 3*time.Second)
	if err == nil {
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(conn)
		// 127.0.0.1
		udpIP := conn.LocalAddr().(*net.UDPAddr).IP
		interfaces, err := net.Interfaces()
		if err != nil {

		} else {
			for _, iface := range interfaces {
				addrs, err := iface.Addrs()
				if err != nil {

					continue
				}
				if netutil.ContainsIPv4(iface, udpIP) && len(addrs) > 1 {
					// 遍历所有地址，安全地找到第一个非 IPv4 的 IPNet 地址（通常是 IPv6）
					// 不再硬编码索引 [1]，不再做未检查的类型断言
					for _, addr := range addrs {
						ipNet, ok := addr.(*net.IPNet)
						if !ok {
							continue
						}
						// 跳过 IPv4 地址，返回找到的第一个 IPv6 地址
						if ipNet.IP.To4() == nil {
							return ipNet.IP.String()
						}
					}
				}
			}
		}
	}

	// 备用方案：当无法连接到8.8.8.8:53时，从网络接口获取IPv6地址
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
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() == nil {
				// 返回找到的第一个非回环IPv6地址
				return ipNet.IP.String()
			}
		}
	}

	return ""
}

// GetIPv6Addrs
// @return map[string]string:
func GetIPv6Addrs() map[string]string {
	IPv6Map := make(map[string]string)

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
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() == nil {
					IPv6Map[iface.Name] = ipNet.IP.String()
				}
			}
		}
	}
	return IPv6Map
}

// IPv6Reachable 检查IPv6地址是否可达
//
// 参数:
//   - ip: 要检查的IPv6地址
//
// 返回值:
//   - bool: 如果地址可达返回true，否则返回false
func IPv6Reachable(ip string) bool {
	if err := netutil.Reachable(ip); err != nil {
		return false
	} else {
		return true
	}
}
