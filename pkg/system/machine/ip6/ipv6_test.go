package ip6

import (
	"testing"
)

func TestGetIPv6Addrs(t *testing.T) {
	addrs := GetIPv6Addrs()
	t.Log(addrs)
}

func TestGetIPv6Addr(t *testing.T) {
	addr := GetIPv6Addr()
	t.Log(addr)
}

func TestIPv6Reachable(t *testing.T) {
	reachable := IPv6Reachable("[fc00::64]:3306")
	t.Log(reachable)
}
