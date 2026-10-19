package ssrf

import (
	"fmt"
	"net"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

// Protection SSRF防护配置
type Protection struct {
	AllowPrivateIP         bool
	DomainFilterMode       bool     // true: 白名单, false: 黑名单
	DomainList             []string // 域名格式, 如 example.com, *.example.com
	IPFilterMode           bool     // true: 白名单, false: 黑名单
	IPList                 []string // CIDR 或单个IP
	AllowedPorts           []int    // 允许的端口
	ApplyIPFilterForDomain bool     // 对域名启用解析后IP过滤
}

// Default 默认SSRF防护配置：拒绝私有地址，域名/IP均为空白名单（不放行任何域名）
var Default = &Protection{
	AllowPrivateIP:   false,
	DomainFilterMode: true,
	DomainList:       []string{},
	IPFilterMode:     true,
	IPList:           []string{},
	AllowedPorts:     []int{},
}

// New 从参数构建SSRF防护配置
// allowedPorts 支持格式: "80", "443", "8000-9000"
func New(allowPrivateIP bool, domainFilterMode bool, ipFilterMode bool, domainList, ipList, allowedPorts []string, applyIPFilterForDomain bool) (*Protection, error) {
	ports, err := parsePortRanges(allowedPorts)
	if err != nil {
		return nil, fmt.Errorf("invalid port configuration: %v", err)
	}
	return &Protection{
		AllowPrivateIP:         allowPrivateIP,
		DomainFilterMode:       domainFilterMode,
		DomainList:             domainList,
		IPFilterMode:           ipFilterMode,
		IPList:                 ipList,
		AllowedPorts:           ports,
		ApplyIPFilterForDomain: applyIPFilterForDomain,
	}, nil
}

// privateIPv4Nets IPv4 私有/保留/特殊用途网段
// 参考 IANA IPv4 Special-Purpose Address Registry
// https://www.iana.org/assignments/iana-ipv4-special-registry/
var privateIPv4Nets = []net.IPNet{
	{IP: net.IPv4(0, 0, 0, 0), Mask: net.CIDRMask(8, 32)},          // 0.0.0.0/8 ("This network" / 未指定)
	{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)},         // 10.0.0.0/8 (私有)
	{IP: net.IPv4(100, 64, 0, 0), Mask: net.CIDRMask(10, 32)},      // 100.64.0.0/10 (运营商级 NAT / CGNAT)
	{IP: net.IPv4(127, 0, 0, 0), Mask: net.CIDRMask(8, 32)},        // 127.0.0.0/8 (回环)
	{IP: net.IPv4(169, 254, 0, 0), Mask: net.CIDRMask(16, 32)},     // 169.254.0.0/16 (链路本地)
	{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(12, 32)},      // 172.16.0.0/12 (私有)
	{IP: net.IPv4(192, 0, 0, 0), Mask: net.CIDRMask(24, 32)},       // 192.0.0.0/24 (IETF 协议分配)
	{IP: net.IPv4(192, 0, 2, 0), Mask: net.CIDRMask(24, 32)},       // 192.0.2.0/24 (TEST-NET-1)
	{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(16, 32)},     // 192.168.0.0/16 (私有)
	{IP: net.IPv4(198, 18, 0, 0), Mask: net.CIDRMask(15, 32)},      // 198.18.0.0/15 (基准测试)
	{IP: net.IPv4(198, 51, 100, 0), Mask: net.CIDRMask(24, 32)},    // 198.51.100.0/24 (TEST-NET-2)
	{IP: net.IPv4(203, 0, 113, 0), Mask: net.CIDRMask(24, 32)},     // 203.0.113.0/24 (TEST-NET-3)
	{IP: net.IPv4(224, 0, 0, 0), Mask: net.CIDRMask(4, 32)},        // 224.0.0.0/4 (组播)
	{IP: net.IPv4(240, 0, 0, 0), Mask: net.CIDRMask(4, 32)},        // 240.0.0.0/4 (保留)
	{IP: net.IPv4(255, 255, 255, 255), Mask: net.CIDRMask(32, 32)}, // 255.255.255.255/32 (受限广播)
}

// privateIPv6Nets IPv6 私有/保留/特殊用途网段
// 参考 IANA IPv6 Special-Purpose Address Registry
// https://www.iana.org/assignments/iana-ipv6-special-registry/
var privateIPv6Nets = func() []net.IPNet {
	cidrs := []string{
		"::/128",        // 未指定地址
		"::1/128",       // 回环
		"::ffff:0:0/96", // IPv4-mapped
		"64:ff9b::/96",  // IPv4/IPv6 translation
		"100::/64",      // Discard-Only
		"2001::/23",     // IETF Protocol Assignments
		"2001:db8::/32", // 文档
		"fc00::/7",      // Unique Local Address (ULA)
		"fe80::/10",     // 链路本地
		"ff00::/8",      // 组播
	}
	nets := make([]net.IPNet, 0, len(cidrs))
	for _, c := range cidrs {
		if _, n, err := net.ParseCIDR(c); err == nil && n != nil {
			nets = append(nets, *n)
		}
	}
	return nets
}()

// IsPrivateIP 检查IP是否为私有/保留/特殊用途地址
func IsPrivateIP(ip net.IP) bool {
	if ip == nil {
		return true
	}
	// 未指定地址 (0.0.0.0, ::)
	if ip.IsUnspecified() {
		return true
	}
	// 回环、链路本地 (unicast/multicast)
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}
	// 接口本地组播 (IPv6 ff01::/16 等)
	if ip.IsInterfaceLocalMulticast() {
		return true
	}

	if v4 := ip.To4(); v4 != nil {
		for _, privateNet := range privateIPv4Nets {
			if privateNet.Contains(v4) {
				return true
			}
		}
		return false
	}

	// IPv6 检查
	for _, privateNet := range privateIPv6Nets {
		if privateNet.Contains(ip) {
			return true
		}
	}
	// 兜底: Go 标准库识别的其他私有地址
	if ip.IsPrivate() {
		return true
	}
	return false
}

// IsIPInCIDRList 检查IP是否在CIDR列表（或单IP列表）中
func IsIPInCIDRList(ip net.IP, cidrList []string) bool {
	for _, cidr := range cidrList {
		_, network, err := net.ParseCIDR(cidr)
		if err != nil {
			// 尝试作为单个IP处理
			if listIP := net.ParseIP(cidr); listIP != nil && ip.Equal(listIP) {
				return true
			}
			continue
		}
		if network.Contains(ip) {
			return true
		}
	}
	return false
}

// parsePortRanges 解析端口配置
// 支持格式: "80", "443", "8000-9000"
func parsePortRanges(portConfigs []string) ([]int, error) {
	var ports []int

	for _, config := range portConfigs {
		config = strings.TrimSpace(config)
		if config == "" {
			continue
		}

		if strings.Contains(config, "-") {
			// 端口范围 "8000-9000"
			parts := strings.Split(config, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid port range format: %s", config)
			}
			startPort, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				return nil, fmt.Errorf("invalid start port in range %s: %v", config, err)
			}
			endPort, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return nil, fmt.Errorf("invalid end port in range %s: %v", config, err)
			}
			if startPort > endPort {
				return nil, fmt.Errorf("invalid port range %s: start port cannot be greater than end port", config)
			}
			if startPort < 1 || startPort > 65535 || endPort < 1 || endPort > 65535 {
				return nil, fmt.Errorf("port range %s contains invalid port numbers (must be 1-65535)", config)
			}
			for port := startPort; port <= endPort; port++ {
				ports = append(ports, port)
			}
		} else {
			// 单个端口 "80"
			port, err := strconv.Atoi(config)
			if err != nil {
				return nil, fmt.Errorf("invalid port number: %s", config)
			}
			if port < 1 || port > 65535 {
				return nil, fmt.Errorf("invalid port number %d (must be 1-65535)", port)
			}
			ports = append(ports, port)
		}
	}
	return ports, nil
}

// isAllowedPort 检查端口是否被允许
func (p *Protection) isAllowedPort(port int) bool {
	if len(p.AllowedPorts) == 0 {
		return true
	}
	return slices.Contains(p.AllowedPorts, port)
}

// isDomainListed 检查域名是否在列表中（支持精确与通配符 *.example.com）
func isDomainListed(domain string, list []string) bool {
	if len(list) == 0 {
		return false
	}
	domain = strings.ToLower(domain)
	for _, item := range list {
		item = strings.ToLower(strings.TrimSpace(item))
		if item == "" {
			continue
		}
		if domain == item {
			return true
		}
		if after, ok := strings.CutPrefix(item, "*."); ok {
			suffix := after
			if strings.HasSuffix(domain, "."+suffix) || domain == suffix {
				return true
			}
		}
	}
	return false
}

func (p *Protection) isDomainAllowed(domain string) bool {
	listed := isDomainListed(domain, p.DomainList)
	if p.DomainFilterMode { // 白名单
		return listed
	}
	// 黑名单
	return !listed
}

// IsIPAccessAllowed 检查IP是否允许访问
func (p *Protection) IsIPAccessAllowed(ip net.IP) bool {
	// 私有IP限制
	if IsPrivateIP(ip) && !p.AllowPrivateIP {
		return false
	}
	listed := IsIPInCIDRList(ip, p.IPList)
	if p.IPFilterMode { // 白名单
		return listed
	}
	// 黑名单
	return !listed
}

func (p *Protection) ipAccessError(host string, ip net.IP) error {
	if host != "" {
		if IsPrivateIP(ip) && !p.AllowPrivateIP {
			return fmt.Errorf("private IP address not allowed: %s resolves to %s", host, ip.String())
		}
		if p.IPFilterMode {
			return fmt.Errorf("ip not in whitelist: %s resolves to %s", host, ip.String())
		}
		return fmt.Errorf("ip in blacklist: %s resolves to %s", host, ip.String())
	}
	if IsPrivateIP(ip) && !p.AllowPrivateIP {
		return fmt.Errorf("private IP address not allowed: %s", ip.String())
	}
	if p.IPFilterMode {
		return fmt.Errorf("ip not in whitelist: %s", ip.String())
	}
	return fmt.Errorf("ip in blacklist: %s", ip.String())
}

// ValidateNetworkTarget 在拨号前校验主机与端口
func (p *Protection) ValidateNetworkTarget(host string, port int) error {
	host = strings.TrimSpace(host)
	if host == "" {
		return fmt.Errorf("invalid host")
	}
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port: %d", port)
	}
	if !p.isAllowedPort(port) {
		return fmt.Errorf("port %d is not allowed", port)
	}

	if ip := net.ParseIP(host); ip != nil {
		if !p.IsIPAccessAllowed(ip) {
			return p.ipAccessError("", ip)
		}
		return nil
	}

	if !p.isDomainAllowed(host) {
		if p.DomainFilterMode {
			return fmt.Errorf("domain not in whitelist: %s", host)
		}
		return fmt.Errorf("domain in blacklist: %s", host)
	}
	return nil
}

// ValidateResolvedIP 在拨号前校验域名解析出的IP（防止 DNS 重绑定到内网）
func (p *Protection) ValidateResolvedIP(host string, ip net.IP) error {
	if !p.IsIPAccessAllowed(ip) {
		return p.ipAccessError(host, ip)
	}
	return nil
}

// ValidateURL 验证URL是否安全（协议、端口、域名/IP 黑白名单、私有地址）
func (p *Protection) ValidateURL(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("invalid URL format: %v", err)
	}

	// 只允许HTTP/HTTPS协议
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("unsupported protocol: %s (only http/https allowed)", u.Scheme)
	}

	// 解析主机和端口
	host, portStr, err := net.SplitHostPort(u.Host)
	if err != nil {
		// 没有端口，使用默认端口
		host = u.Hostname()
		if u.Scheme == "https" {
			portStr = "443"
		} else {
			portStr = "80"
		}
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid port: %s", portStr)
	}

	if err := p.ValidateNetworkTarget(host, port); err != nil {
		return err
	}

	// 如果 host 是 IP，或未启用对域名应用 IP 过滤，则到此通过
	if net.ParseIP(host) != nil || !p.ApplyIPFilterForDomain {
		return nil
	}

	// 解析域名对应IP并检查
	ips, err := net.LookupIP(host)
	if err != nil {
		return fmt.Errorf("DNS resolution failed for %s: %v", host, err)
	}
	for _, ip := range ips {
		if err := p.ValidateResolvedIP(host, ip); err != nil {
			return err
		}
	}
	return nil
}
