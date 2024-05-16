package ip6

import (
	"fmt"
	"log"
	"net"

	"github.com/jianlu8023/go-tools/internal/machine"
)

// GetIPv6Addr
// @Description: GetIPv6Addr
// @author ght
// @date 2023-10-20 18:39:24
// @return string:
func GetIPv6Addr() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 127.0.0.1
	udpIP := conn.LocalAddr().(*net.UDPAddr).IP
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			return ""
		}
		if machine.ContainsIPv4(iface, udpIP) {
			return addrs[1].(*net.IPNet).IP.String()
		}
	}

	return ""
}

// GetIPv6Addrs
// @Description: GetIPv6Addrs
// @author ght
// @date 2023-10-20 18:40:36
// @return map[string]string:
func GetIPv6Addrs() map[string]string {
	IPv6Map := make(map[string]string)

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln("Failed to get interfaces:", err)
		return nil
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagBroadcast == net.FlagBroadcast {
			addresses, err := iface.Addrs()
			if err != nil {
				log.Printf("Failed to get addresses for interface %s: %v\n", iface.Name, err)
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

func IPv6Reachable(ip string) bool {
	if err := machine.Reachable(ip); err != nil {
		return false
	} else {
		return true
	}
}
