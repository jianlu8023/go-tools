package mac

import (
	"net"

	"github.com/jianlu8023/go-tools/internal/machine"
)

// GetMacAddr
// @Description: GetMacAddr
// @author ght
// @date 2023-10-20 18:51:25
// @return string:
func GetMacAddr() string {
	conn, _ := net.Dial("udp", "8.8.8.8:53")

	udpIP := conn.LocalAddr().(*net.UDPAddr).IP
	interfaces, _ := net.Interfaces()

	for _, iface := range interfaces {
		if machine.ContainsIPv4(iface, udpIP) {
			return iface.HardwareAddr.String()
		}

	}
	return ""

}
