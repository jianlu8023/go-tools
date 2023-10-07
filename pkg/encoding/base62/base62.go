package base62

const (
	base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// IntToBase62
// @Description: IntToBase62
// @author ght
// @date 2023-10-07 16:53:21
// @param n:
// @return string:
func IntToBase62(n int) string {
	if n == 0 {
		return string(base62[0])
	}
	var result []byte
	for n > 0 {
		result = append(result, base62[n%62])
		n /= 62
	}
	// 反转字符串
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
