package http

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// 包级命名空间，区分不同业务内容的 ETag 计算
const etagNamespace = "go-tools"

// ETagFor 生成内容的弱 ETag（W/"..."），同命名空间内相同内容生成相同值
func ETagFor(namespace, content string) string {
	h := sha256.New()
	h.Write([]byte(etagNamespace))
	h.Write([]byte{0})
	h.Write([]byte(namespace))
	h.Write([]byte{0})
	h.Write([]byte(content))
	return `W/"` + hex.EncodeToString(h.Sum(nil)) + `"`
}

// ETagMatches 按 RFC 9110 §13.1.2 弱比较判断 If-None-Match 头是否匹配 etag：
// W/ 前缀两侧均忽略，"*" 匹配任意，逗号分隔列表任一命中即为匹配
func ETagMatches(ifNoneMatch, etag string) bool {
	ifNoneMatch = strings.TrimSpace(ifNoneMatch)
	if ifNoneMatch == "" {
		return false
	}
	if ifNoneMatch == "*" {
		return true
	}
	etag = strings.TrimPrefix(etag, "W/")
	for _, candidate := range strings.Split(ifNoneMatch, ",") {
		if strings.TrimPrefix(strings.TrimSpace(candidate), "W/") == etag {
			return true
		}
	}
	return false
}
