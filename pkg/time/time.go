package time

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

// humanDuration returns a human-readable approximation of a
// duration (eg. "About a minute", "4 hours ago", etc.).
func humanDuration(d time.Duration) string {
	seconds := int(d.Seconds())

	switch {
	case seconds < 1:
		return "Less than a second"
	case seconds == 1:
		return "1 second"
	case seconds < 60:
		return fmt.Sprintf("%d seconds", seconds)
	}

	minutes := int(d.Minutes())
	switch {
	case minutes == 1:
		return "About a minute"
	case minutes < 60:
		return fmt.Sprintf("%d minutes", minutes)
	}

	hours := int(math.Round(d.Hours()))
	switch {
	case hours == 1:
		return "About an hour"
	case hours < 48:
		return fmt.Sprintf("%d hours", hours)
	case hours < 24*7*2:
		return fmt.Sprintf("%d days", hours/24)
	case hours < 24*30*2:
		return fmt.Sprintf("%d weeks", hours/24/7)
	case hours < 24*365*2:
		return fmt.Sprintf("%d months", hours/24/30)
	}

	return fmt.Sprintf("%d years", int(d.Hours())/24/365)
}

func HumanTime(t time.Time, zeroValue string) string {
	return humanTime(t, zeroValue)
}

func HumanTimeLower(t time.Time, zeroValue string) string {
	return strings.ToLower(humanTime(t, zeroValue))
}

func humanTime(t time.Time, zeroValue string) string {
	if t.IsZero() {
		return zeroValue
	}

	delta := time.Since(t)
	if int(delta.Hours())/24/365 < -20 {
		return "Forever"
	} else if delta < 0 {
		return humanDuration(-delta) + " from now"
	}

	return humanDuration(delta) + " ago"
}

// ParseDuration 解析时间
// @param d: 字符串
// @return time.Duration: 时间
// @return error: 错误
// @example ParseDuration("0d5h15m40s") => 5h15m40s
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	// 仅在 "d" 紧跟在数字后面且处于单位位置时才视为天数单位
	// 使用正则匹配 "<number>d<rest>" 形式，避免 "abcd" 这类误匹配
	dayPattern := regexp.MustCompile(`^([+-]?\d+)d(.*)$`)
	if matches := dayPattern.FindStringSubmatch(d); matches != nil {
		hour, err := strconv.Atoi(matches[1])
		if err != nil {
			return 0, fmt.Errorf("invalid day count in duration %q: %w", d, err)
		}
		dr = time.Hour * 24 * time.Duration(hour)
		rest := matches[2]
		if rest == "" {
			return dr, nil
		}
		ndr, err := time.ParseDuration(rest)
		if err != nil {
			return 0, fmt.Errorf("invalid duration suffix in %q: %w", d, err)
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

func ParseTimeLocal(datetime string) (time.Time, error) {
	return dateparse.ParseLocal(datetime)
}

func ParseTimeIn(datetime string, in *time.Location) (time.Time, error) {
	return dateparse.ParseIn(datetime, in)
}

// FormatTimeISO8601 将时间格式化为ISO 8601格式
// @param t: 时间
// @return string: ISO 8601格式的时间字符串
// @example FormatTimeISO8601(time.Now()) => "2023-10-25T15:30:45Z"
func FormatTimeISO8601(t time.Time) string {
	return t.Format(time.RFC3339)
}

// FormatTimeCommon 将时间格式化为常见格式 (YYYY-MM-DD HH:MM:SS)
// @param t: 时间
// @return string: 格式化后的时间字符串
// @example FormatTimeCommon(time.Now()) => "2023-10-25 15:30:45"
func FormatTimeCommon(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// UnixToTime 将Unix时间戳转换为time.Time
// @param timestamp: Unix时间戳(秒)
// @return time.Time: 对应的时间
// @example UnixToTime(1635153045) => 2021-10-25 15:30:45 +0000 UTC
func UnixToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// UnixNanoToTime 将Unix纳秒时间戳转换为time.Time
// @param timestamp: Unix纳秒时间戳
// @return time.Time: 对应的时间
func UnixNanoToTime(timestamp int64) time.Time {
	return time.Unix(0, timestamp)
}

// TimeToUnix 将time.Time转换为Unix时间戳(秒)
// @param t: 时间
// @return int64: Unix时间戳
func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

// TimeToUnixNano 将time.Time转换为Unix纳秒时间戳
// @param t: 时间
// @return int64: Unix纳秒时间戳
func TimeToUnixNano(t time.Time) int64 {
	return t.UnixNano()
}

// GetCurrentTimestamp 获取当前Unix时间戳(秒)
// @return int64: 当前Unix时间戳
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// GetCurrentTimestampNano 获取当前Unix纳秒时间戳
// @return int64: 当前Unix纳秒时间戳
func GetCurrentTimestampNano() int64 {
	return time.Now().UnixNano()
}

// GetDurationDays 计算两个时间之间的天数差
// @param start: 开始时间
// @param end: 结束时间
// @return int64: 天数差
func GetDurationDays(start, end time.Time) int64 {
	return int64(end.Sub(start).Hours() / 24)
}

// GetDurationHours 计算两个时间之间的小时差
// @param start: 开始时间
// @param end: 结束时间
// @return float64: 小时差
func GetDurationHours(start, end time.Time) float64 {
	return end.Sub(start).Hours()
}

// GetDurationMinutes 计算两个时间之间的分钟差
// @param start: 开始时间
// @param end: 结束时间
// @return float64: 分钟差
func GetDurationMinutes(start, end time.Time) float64 {
	return end.Sub(start).Minutes()
}

// GetDurationSeconds 计算两个时间之间的秒差
// @param start: 开始时间
// @param end: 结束时间
// @return float64: 秒差
func GetDurationSeconds(start, end time.Time) float64 {
	return end.Sub(start).Seconds()
}

// IsSameDay 判断两个时间是否在同一天
// @param t1: 第一个时间
// @param t2: 第二个时间
// @return bool: 是否在同一天
func IsSameDay(t1, t2 time.Time) bool {
	t1Year, t1Month, t1Day := t1.Date()
	t2Year, t2Month, t2Day := t2.Date()
	return t1Year == t2Year && t1Month == t2Month && t1Day == t2Day
}

// beijingLocation 缓存北京时区，避免每次调用 GetBeijingTime 都重新加载
// 若 LoadLocation 失败（如系统缺少 tzdata），回退到固定 UTC+8 偏移
var beijingLocation = func() *time.Location {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.FixedZone("CST", 8*60*60)
	}
	return loc
}()

// GetBeijingTime 获取当前北京时间(UTC+8)
// @return time.Time: 当前北京时间
func GetBeijingTime() time.Time {
	return time.Now().In(beijingLocation)
}

// ConvertTimezone 将时间从一个时区转换到另一个时区
// @param t: 原始时间
// @param fromLoc: 原始时区
// @param toLoc: 目标时区
// @return time.Time: 转换后的时间
func ConvertTimezone(t time.Time, fromLoc, toLoc *time.Location) time.Time {
	return t.In(fromLoc).In(toLoc)
}

// Wait 等待指定的毫秒数
// @param ms: 等待的毫秒数
func Wait(ms int64) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// WaitSeconds 等待指定的秒数
// @param seconds: 等待的秒数
func WaitSeconds(seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
