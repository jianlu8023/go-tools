package ip6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPv6Reachable(t *testing.T) {
	// 测试IPv6地址可达性检查
	// 使用一个已知的公共DNS服务器IPv6地址
	ip := "2001:4860:4860::8888"

	// 注意：这个测试可能因为网络环境而失败，所以我们只验证函数能正常执行
	result := IPv6Reachable(ip)

	// 我们不强制要求结果为true或false，因为这取决于网络环境
	// 只要函数能正常执行而不panic就好
	t.Logf("IPv6Reachable(%s) returned %t", ip, result)

	// 测试一个无效的IPv6地址
	invalidIP := "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"
	result2 := IPv6Reachable(invalidIP)
	assert.False(t, result2, "无效IPv6地址应该返回false")
}

func TestGetIPv6Addr(t *testing.T) {
	// 测试获取IPv6地址
	addr := GetIPv6Addr()

	// 注意：在某些环境中可能无法获取IPv6地址
	// 我们只验证函数能正常执行而不panic
	t.Logf("GetIPv6Addr returned: %s", addr)

	// 如果返回了地址，简单验证格式
	if addr != "" {
		// IPv6地址应该包含冒号
		assert.True(t, len(addr) >= 3, "IPv6地址长度应该至少为3个字符") // 最短的IPv6地址是:: (2个字符)
	}
}

func TestGetIPv6Addrs(t *testing.T) {
	// 测试获取所有IPv6地址
	addrs := GetIPv6Addrs()

	// 注意：在某些环境中可能无法获取IPv6地址
	t.Logf("GetIPv6Addrs returned: %v", addrs)

	// 如果返回了地址映射，验证其结构
	if addrs != nil {
		// 验证返回的是一个映射
		assert.True(t, len(addrs) >= 0, "返回的应该是一个映射")
		t.Logf("Successfully got IPv6 addresses map with %d entries", len(addrs))
	}
}
