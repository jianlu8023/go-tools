package mac

import (
	"net"

	"github.com/jianlu8023/go-tools/internal/machine"
)

// GetMacAddr
// @return string:
func GetMacAddr() string {
	conn, _ := net.Dial("udp", "8.8.8.8:53")
	defer conn.Close()

	udpIP := conn.LocalAddr().(*net.UDPAddr).IP
	interfaces, _ := net.Interfaces()

	for _, iface := range interfaces {
		if machine.ContainsIPv4(iface, udpIP) {
			return iface.HardwareAddr.String()
		}

	}
	return ""

}
