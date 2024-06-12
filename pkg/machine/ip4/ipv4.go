package ip4

import (
	"log"
	"net"
	"strings"

	"github.com/jianlu8023/go-tools/internal/machine"
)

// GetIPv4Addr
// @return string:
func GetIPv4Addr() string {
	// use udp get ip4 address
	dial, err := net.Dial("udp", "8.8.8.8:53")
	defer dial.Close()
	if err != nil {
		log.Fatalln("GetIPv4Addr Error:", err.Error())
		return ""
	}
	localAddr := dial.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip
}

// GetIPv4Addrs
// @return map[string]string:
func GetIPv4Addrs() map[string]string {
	IPv4Map := make(map[string]string)

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

	if err := machine.Reachable(ip); err != nil {
		return false
	} else {
		return true
	}
}
