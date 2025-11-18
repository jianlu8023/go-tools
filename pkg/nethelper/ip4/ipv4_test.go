package ip4

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIPv4Addr(t *testing.T) {
	addr := GetIPv4Addr()
	t.Log(addr)
}

func TestGetIPv4Addrs(t *testing.T) {
	addrs := GetIPv4Addrs()
	t.Log(addrs)
}

func TestIPAddrInterface(t *testing.T) {

	interfaces, _ := net.Interfaces()

	for i, iface := range interfaces {

		ifaceHardwareAddr := iface.HardwareAddr
		ifaceName := iface.Name
		ifaceFlags := iface.Flags
		ifaceMTU := iface.MTU

		// fmt.Println(fmt.Sprintf("%d ifaceName : %v, ifaceHardwareAddr: %v, ifaceFlags: %v, ifaceMTU: %v",
		// 	i, ifaceName, ifaceHardwareAddr, ifaceFlags, ifaceMTU))

		addrs, _ := iface.Addrs()

		for j, addr := range addrs {
			n := addr.(*net.IPNet) // 127.0.0.1/24
			ip := n.IP             // 127.0.0.1

			fmt.Println(fmt.Sprintf(
				"%d ifaceName : %v, ifaceHardwareAddr: %v, ifaceFlags: %v, ifaceMTU: %v, addr %d, ip4: %v ",
				i, ifaceName, ifaceHardwareAddr, ifaceFlags, ifaceMTU, j, ip))

			// network := n.Network() // ip4+net
			// s := n.String()        // 127.0.0.1/24
			// fmt.Println(fmt.Sprintf("*net.IPAddr : %v, n.network: %v, n.string: %v, IP: %v", n, network, s, i))
			// if !n.IP.IsLoopback() {
			// 	if n.IP.To4() == nil {
			// 		fmt.Println(fmt.Sprintf("interface %d addr %d ipv6 %v", i, j, n.IP.String()))
			// 	}
			// }

		}
	}

}

func TestGetLocalIPv4(t *testing.T) {
	// 测试获取本地IPv4地址
	ip, err := GetLocalIPv4()

	// 注意：在某些环境中可能无法获取IP地址，所以我们要合理处理错误情况
	// 但在大多数正常网络环境中应该能成功获取

	if err != nil {
		// 如果出现错误，检查是否是预期的错误消息
		t.Logf("GetLocalIPv4 returned error: %v", err)
	} else {
		// 如果没有错误，验证IP地址格式
		assert.NotEmpty(t, ip, "IP地址不应该为空")

		// 简单验证IP地址格式（不为空且包含点）
		assert.True(t, len(ip) >= 7, "IP地址长度应该至少为7个字符") // 最短的IPv4地址是x.x.x.x (7个字符)

		t.Logf("Successfully got local IPv4 address: %s", ip)
	}
}

func TestGetLocalIP(t *testing.T) {
	// 测试GetLocalIP别名函数
	ip, err := GetLocalIP()

	if err != nil {
		t.Logf("GetLocalIP returned error: %v", err)
	} else {
		assert.NotEmpty(t, ip, "IP地址不应该为空")

		assert.True(t, len(ip) >= 7, "IP地址长度应该至少为7个字符")

		t.Logf("Successfully got local IP address: %s", ip)
	}
}

func TestIPv4Readable(t *testing.T) {
	// 测试IPv4地址可达性检查
	// 使用一个已知的公共DNS服务器地址
	ip := "8.8.8.8"

	// 注意：这个测试可能因为网络环境而失败，所以我们只验证函数能正常执行
	result := IPv4Readable(ip)

	// 我们不强制要求结果为true或false，因为这取决于网络环境
	// 只要函数能正常执行而不panic就好
	t.Logf("IPv4Readable(%s) returned %t", ip, result)

	// 测试一个无效的IP地址
	invalidIP := "999.999.999.999"
	result2 := IPv4Readable(invalidIP)
	assert.False(t, result2, "无效IP地址应该返回false")
}
