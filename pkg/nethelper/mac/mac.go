package mac

import (
	"net"

	"github.com/jianlu8023/go-tools/v2/internal/netutil"
)

// GetMacAddr
// @return string:
func GetMacAddr() string {
	// 尝试使用UDP连接获取MAC地址
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err == nil {
		defer func(conn net.Conn) {
			_ = conn.Close()

		}(conn)

		udpIP := conn.LocalAddr().(*net.UDPAddr).IP
		interfaces, _ := net.Interfaces()

		for _, iface := range interfaces {
			if netutil.ContainsIPv4(iface, udpIP) && len(iface.HardwareAddr) > 0 {
				return iface.HardwareAddr.String()
			}
		}
	}

	// 备用方案：当无法连接到8.8.8.8:53时，直接获取活动网络接口的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		// 跳过非活动接口、回环接口和没有MAC地址的接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || len(iface.HardwareAddr) == 0 {
			continue
		}

		// 优先选择有广播功能的接口（通常是以太网接口）
		if iface.Flags&net.FlagBroadcast != 0 {
			return iface.HardwareAddr.String()
		}
	}

	// 如果没有找到有广播功能的接口，返回第一个可用的MAC地址
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && len(iface.HardwareAddr) > 0 {
			return iface.HardwareAddr.String()
		}
	}

	return ""
}
