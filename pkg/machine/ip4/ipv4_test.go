package ip4

import (
	"fmt"
	"net"
	"testing"
)

func TestGetIPv4Addr(t *testing.T) {
	addr := GetIPv4Addr()
	t.Log(addr)
}

func TestGetIPv4Addrs(t *testing.T) {
	addrs := GetIPv4Addrs()
	t.Log(addrs)
}

func TestIPv4Readable(t *testing.T) {
	readable := IPv4Readable("192.168.58.110:3306")
	t.Log(readable)
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
