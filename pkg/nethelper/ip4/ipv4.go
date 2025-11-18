package ip4

import (
	"errors"
	"net"
	"strings"
	"time"

	"github.com/jianlu8023/go-tools/v2/internal/netutil"
)

// GetIPv4Addr
// @return string:
func GetIPv4Addr() string {
	// 尝试使用UDP连接获取IP地址
	dial, err := net.DialTimeout("udp", "8.8.8.8:53", 3*time.Second)
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

func GetLocalIP() (string, error) {
	return GetLocalIPv4()
}

func GetLocalIPv4() (string, error) {
	addrs, err := getAddrs()
	if err != nil {
		return "", errors.New("获取IP失败: " + err.Error())
	}
	for _, addr := range addrs {
		ip := getLocalIP(addr)
		if ip == nil {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}
		return ip.String(), nil
	}
	return "", errors.New("获取 IPv4 地址失败")
}

func getAddrs() ([]net.Addr, error) {
	var addrss []net.Addr
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // 网卡没有开启
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // 这是个环回地址
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		addrss = append(addrss, addrs...)
	}
	if len(addrss) <= 0 {
		return nil, errors.New("你好像没有接入局域网？")
	}
	return addrss, nil
}

func getLocalIP(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	default:
		return nil
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	return ip
}
