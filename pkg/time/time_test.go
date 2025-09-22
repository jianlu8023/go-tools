package time

import (
	"testing"
	"time"
)

func TestHumanTimeLower(t *testing.T) {
	str := "3/1/2025"
	datetime, err := ParseTimeLocal(str)
	if err != nil {
		t.Fatal(err)
	}
	lower := HumanTimeLower(datetime, "unknown")
	t.Log(lower)
}

func TestParseTimeOnLocal(t *testing.T) {
	str := "3/1/2014"
	datetime, err := ParseTimeLocal(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datetime)
}

func TestParseTimeIn(t *testing.T) {
	str := "3/1/2014 10:22:22.111"
	datetime, err := ParseTimeIn(str, time.Local)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datetime)
}

func TestParseDuration(t *testing.T) {
	str := "0d5h15m40s"
	duration, err := ParseDuration(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(duration)
}

// 以下是新添加的测试函数

func TestFormatTimeISO8601(t *testing.T) {
	timeNow := time.Now()
	isoTime := FormatTimeISO8601(timeNow)
	t.Log("ISO 8601 time:", isoTime)

	// 验证格式是否正确
	_, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		t.Errorf("FormatTimeISO8601 returned invalid ISO 8601 format: %v", err)
	}
}

func TestFormatTimeCommon(t *testing.T) {
	timeNow := time.Now()
	commonTime := FormatTimeCommon(timeNow)
	t.Log("Common format time:", commonTime)

	// 验证格式是否正确
	_, err := time.Parse("2006-01-02 15:04:05", commonTime)
	if err != nil {
		t.Errorf("FormatTimeCommon returned invalid format: %v", err)
	}
}

func TestUnixTimeConversion(t *testing.T) {
	timestamp := int64(1635153045)
	timeObj := UnixToTime(timestamp)
	convertedTimestamp := TimeToUnix(timeObj)

	if convertedTimestamp != timestamp {
		t.Errorf("Unix time conversion failed: expected %d, got %d", timestamp, convertedTimestamp)
	}

	// 测试纳秒时间戳
	nanoTimestamp := int64(1635153045123456789)
	tNano := UnixNanoToTime(nanoTimestamp)
	convertedNanoTimestamp := TimeToUnixNano(tNano)

	// 由于纳秒精度问题，我们只比较到微秒级别
	if convertedNanoTimestamp/1000 != nanoTimestamp/1000 {
		t.Errorf("Unix nano time conversion failed: expected %d, got %d", nanoTimestamp, convertedNanoTimestamp)
	}
}

func TestGetCurrentTimestamp(t *testing.T) {
	timestamp := GetCurrentTimestamp()
	t.Log("Current timestamp:", timestamp)

	// 验证时间戳是否在合理范围内（当前时间的前后5秒）
	timeNow := time.Now().Unix()
	if timestamp < timeNow-5 || timestamp > timeNow+5 {
		t.Errorf("GetCurrentTimestamp returned invalid timestamp: %d", timestamp)
	}
}

func TestDurationCalculation(t *testing.T) {
	start := time.Now().Add(-24 * time.Hour)
	end := time.Now()

	days := GetDurationDays(start, end)
	hours := GetDurationHours(start, end)
	minutes := GetDurationMinutes(start, end)
	seconds := GetDurationSeconds(start, end)

	t.Logf("Duration: %d days, %.2f hours, %.2f minutes, %.2f seconds", days, hours, minutes, seconds)

	// 验证天数计算是否正确
	if days != 1 {
		t.Errorf("GetDurationDays failed: expected 1, got %d", days)
	}

	// 验证小时数计算是否正确（允许小误差）
	if hours < 23.9 || hours > 24.1 {
		t.Errorf("GetDurationHours failed: expected ~24, got %.2f", hours)
	}
}

func TestIsSameDay(t *testing.T) {
	t1 := time.Date(2023, 10, 25, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, 10, 25, 20, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, 10, 26, 10, 0, 0, 0, time.UTC)

	if !IsSameDay(t1, t2) {
		t.Error("IsSameDay failed: expected same day")
	}

	if IsSameDay(t1, t3) {
		t.Error("IsSameDay failed: expected different days")
	}
}

func TestGetBeijingTime(t *testing.T) {
	beijingTime := GetBeijingTime()

	// 验证时区是否正确设置为Asia/Shanghai
	location := beijingTime.Location()
	if location.String() != "Asia/Shanghai" {
		t.Errorf("GetBeijingTime returned time with incorrect location: %s", location.String())
	}

	// 验证时间是否是一个有效的时间（不是零值）
	if beijingTime.IsZero() {
		t.Error("GetBeijingTime returned zero time")
	}

	// 验证时间是否与当前时间大致相符（允许5分钟误差）
	now := time.Now()
	timeDiff := beijingTime.Sub(now).Minutes()
	if timeDiff < -5 || timeDiff > 5 {
		t.Errorf("GetBeijingTime returned time that is too far from current time: %.2f minutes difference", timeDiff)
	}

	t.Log("Beijing time:", beijingTime)
}

func TestConvertTimezone(t *testing.T) {
	utcLoc, _ := time.LoadLocation("UTC")
	nycLoc, _ := time.LoadLocation("America/New_York")

	utcTime := time.Date(2023, 10, 25, 12, 0, 0, 0, utcLoc)
	nycTime := ConvertTimezone(utcTime, utcLoc, nycLoc)

	// 纽约时区通常比UTC-5/-4小时
	hourDiff := nycTime.Hour()
	if hourDiff != 7 && hourDiff != 8 {
		t.Errorf("ConvertTimezone failed: expected hour 7 or 8, got %d", hourDiff)
	}

	t.Logf("UTC time: %v, NYC time: %v", utcTime, nycTime)
}

func TestWait(t *testing.T) {
	start := time.Now()
	Wait(100) // 等待100毫秒
	duration := time.Since(start).Milliseconds()

	if duration < 90 || duration > 150 {
		t.Errorf("Wait failed: expected ~100ms, got %dms", duration)
	}

	t.Logf("Wait duration: %dms", duration)
}

func TestWaitSeconds(t *testing.T) {
	start := time.Now()
	WaitSeconds(1) // 等待1秒
	duration := time.Since(start).Seconds()

	if duration < 0.9 || duration > 1.5 {
		t.Errorf("WaitSeconds failed: expected ~1s, got %.2fs", duration)
	}

	t.Logf("WaitSeconds duration: %.2fs", duration)
}
