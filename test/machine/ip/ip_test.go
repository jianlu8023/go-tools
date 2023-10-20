package ip

import (
	"fmt"
	"net"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/machine/ip"
)

func TestGetIp(t *testing.T) {
	addr := ip.GetIPv4Addr()
	fmt.Println(addr)
}

func TestGetIPv4Addrs(t *testing.T) {
	addrs := ip.GetIPv4Addrs()
	fmt.Println(addrs)
}

func TestGetIPv6Addr(t *testing.T) {
	addr := ip.GetIPv6Addr() // fe80::20c:29ff:fe82:c12f
	fmt.Println(addr)
}

func TestGetIPv6Addrs(t *testing.T) {
	addrs := ip.GetIPv6Addrs()
	fmt.Println(addrs)
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
				"%d ifaceName : %v, ifaceHardwareAddr: %v, ifaceFlags: %v, ifaceMTU: %v, addr %d, ip: %v ",
				i, ifaceName, ifaceHardwareAddr, ifaceFlags, ifaceMTU, j, ip))

			// network := n.Network() // ip+net
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
