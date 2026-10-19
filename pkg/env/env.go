package env

import (
	"os"
	"strconv"
	"time"
)

// GetInt 读取整数环境变量，未设置或解析失败时返回默认值
// @param name 环境变量名
// @param defaultValue 默认值
// @return int 环境变量值或默认值
func GetInt(name string, defaultValue int) int {
	v, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		return defaultValue
	}
	return v
}

// GetString 读取字符串环境变量，未设置时返回默认值
// @param name 环境变量名
// @param defaultValue 默认值
// @return string 环境变量值或默认值
func GetString(name string, defaultValue string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}
	return defaultValue
}

// GetBool 读取布尔环境变量，未设置或解析失败时返回默认值
// @param name 环境变量名
// @param defaultValue 默认值
// @return bool 环境变量值或默认值
func GetBool(name string, defaultValue bool) bool {
	v, err := strconv.ParseBool(os.Getenv(name))
	if err != nil {
		return defaultValue
	}
	return v
}

// GetDuration 读取 time.Duration 环境变量（如 "5s"、"1m"），未设置或解析失败时返回默认值
// @param name 环境变量名
// @param defaultValue 默认值
// @return time.Duration 环境变量值或默认值
func GetDuration(name string, defaultValue time.Duration) time.Duration {
	v, err := time.ParseDuration(os.Getenv(name))
	if err != nil {
		return defaultValue
	}
	return v
}
