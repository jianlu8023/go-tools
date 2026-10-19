package ssrf

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsPrivateIP(t *testing.T) {
	privates := []string{
		"10.0.0.1", "192.168.1.1", "172.16.0.1", "127.0.0.1",
		"169.254.1.1", "0.0.0.0", "100.64.0.1", "198.18.0.1",
		"::1", "fe80::1", "fd12::1", "::",
	}
	for _, s := range privates {
		assert.True(t, IsPrivateIP(net.ParseIP(s)), s)
	}
	publics := []string{"8.8.8.8", "1.1.1.1", "2606:4700:4700::1111"}
	for _, s := range publics {
		assert.False(t, IsPrivateIP(net.ParseIP(s)), s)
	}
}

func TestIsIPInCIDRList(t *testing.T) {
	assert.True(t, IsIPInCIDRList(net.ParseIP("1.2.3.4"), []string{"1.2.3.0/24"}))
	assert.True(t, IsIPInCIDRList(net.ParseIP("1.2.3.4"), []string{"1.2.3.4"})) // 单IP回退
	assert.False(t, IsIPInCIDRList(net.ParseIP("1.2.3.4"), []string{"1.2.4.0/24"}))
	assert.False(t, IsIPInCIDRList(net.ParseIP("1.2.3.4"), nil))
}

func TestNewInvalidPort(t *testing.T) {
	_, err := New(false, true, true, nil, nil, []string{"99999"}, false)
	assert.Error(t, err)
	_, err = New(false, true, true, nil, nil, []string{"200-100"}, false)
	assert.Error(t, err)
	p, err := New(false, true, true, nil, nil, []string{"80", "8000-8002"}, false)
	require.NoError(t, err)
	assert.ElementsMatch(t, []int{80, 8000, 8001, 8002}, p.AllowedPorts)
}

func TestValidateURLProtocol(t *testing.T) {
	p, err := New(true, false, false, nil, nil, nil, false)
	require.NoError(t, err)
	err = p.ValidateURL("ftp://example.com/file")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol")
}

func TestValidateURLDomainWhitelist(t *testing.T) {
	p, err := New(false, true, false, []string{"example.com", "*.good.org"}, nil, nil, false)
	require.NoError(t, err)

	assert.NoError(t, p.ValidateURL("https://example.com/a"))
	assert.NoError(t, p.ValidateURL("https://sub.good.org/a"))
	err = p.ValidateURL("https://evil.com/a")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "domain not in whitelist")
}

func TestValidateURLDomainBlacklist(t *testing.T) {
	p, err := New(false, false, false, []string{"bad.com"}, nil, nil, false)
	require.NoError(t, err)

	assert.NoError(t, p.ValidateURL("https://good.com/a"))
	assert.Error(t, p.ValidateURL("https://bad.com/a"))
}

func TestValidateURLPrivateIP(t *testing.T) {
	p, err := New(false, false, false, nil, nil, nil, false)
	require.NoError(t, err)
	assert.Error(t, p.ValidateURL("http://127.0.0.1/admin"))
	assert.Error(t, p.ValidateURL("http://192.168.0.1:8080/admin"))

	// 允许私网
	pAllow, err := New(true, false, false, nil, nil, nil, false)
	require.NoError(t, err)
	assert.NoError(t, pAllow.ValidateURL("http://127.0.0.1/admin"))
}

func TestValidateURLPortWhitelist(t *testing.T) {
	p, err := New(false, false, false, nil, nil, []string{"443"}, false)
	require.NoError(t, err)

	assert.NoError(t, p.ValidateURL("https://8.8.8.8/"))     // 默认443
	assert.NoError(t, p.ValidateURL("https://8.8.8.8:443/")) // 显式443
	assert.Error(t, p.ValidateURL("http://8.8.8.8:80/"))     // 80 不在白名单
	assert.Error(t, p.ValidateURL("https://8.8.8.8:8443/"))  // 8443 不在白名单
}

func TestValidateNetworkTargetInvalidInput(t *testing.T) {
	p, err := New(false, false, false, nil, nil, nil, false)
	require.NoError(t, err)

	assert.Error(t, p.ValidateNetworkTarget("", 80))
	assert.Error(t, p.ValidateNetworkTarget("example.com", 0))
	assert.Error(t, p.ValidateNetworkTarget("example.com", 70000))
}

func TestValidateResolvedIP(t *testing.T) {
	// 白名单IP模式：私网IP被拒且错误信息带 host
	p, err := New(false, false, true, nil, []string{"8.8.8.0/24"}, nil, false)
	require.NoError(t, err)

	assert.NoError(t, p.ValidateResolvedIP("example.com", net.ParseIP("8.8.8.8")))
	err = p.ValidateResolvedIP("example.com", net.ParseIP("10.0.0.1"))
	require.Error(t, err)
	assert.ErrorContains(t, err, "resolves to")

	// 黑名单IP模式
	pBlack, err := New(true, false, false, nil, []string{"10.0.0.0/8"}, nil, false)
	require.NoError(t, err)
	assert.NoError(t, pBlack.ValidateResolvedIP("example.com", net.ParseIP("192.168.1.1")))
	assert.Error(t, pBlack.ValidateResolvedIP("example.com", net.ParseIP("10.1.1.1")))
	_ = errors.New
}
