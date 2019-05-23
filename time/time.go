package bhtime

import (
	"strconv"
	"time"
)

// Date Format
const (
	YmdHms  = "2006-01-02 15:04:05"
	YmdHm   = "2006-01-02 15:04"
	YmdH    = "2006-01-02 15"
	Ymd     = "2006-01-02"
	Ym      = "2006-01"
	RFC3339 = time.RFC3339
)

// now time.Now()
func now() time.Time {
	return time.Now()
}

// FormatSecond to YmdHms
func FormatSecond(t time.Time) string {
	return t.Format(YmdHms)
}

// FormatMinute to YmdHm
func FormatMinute(t time.Time) string {
	return t.Format(YmdHm)
}

// FormatHour to YmdH
func FormatHour(t time.Time) string {
	return t.Format(YmdH)
}

// TodayTime today time
func TodayTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// FormatDay to Ymd
func FormatDay(t time.Time) string {
	return t.Format(Ymd)
}

// FormatMonth to Ym
func FormatMonth(t time.Time) string {
	return t.Format(Ym)
}

// FormatRFC3339 to RFC3339
func FormatRFC3339(t time.Time) string {
	return t.Format(RFC3339)
}

// 把日期时间转成时间戳
func Date2TimestampFormat(format, date string) int64 {
	unixtime, err := time.ParseInLocation(format, date, time.Local)
	if err != nil {
		return 0
	}
	return unixtime.Unix()
}

// 把日期时间转成时间戳，format：2006-01-02 15:04:05
func Time2Timestamp(date string) int64 {
	t := Date2TimestampFormat(YmdHms, date)
	if t < 1 {
		t = Date2TimestampFormat(YmdHm, date)
	}
	if t < 1 {
		t = Date2TimestampFormat(YmdH, date)
	}
	if t < 1 {
		t = Date2TimestampFormat(Ymd, date)
	}
	if t < 1 {
		t = Date2TimestampFormat("2006/1/2 15:04:05", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("2006/1/2 15:04", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("2006/1/2 15", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("2006/1/2", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("1/2/06 15:04:05", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("1/2/06 15:04", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("1/2/06 15", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("1/2/06", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("1-2-06 15:04:05", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("1-2-06 15:04", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("1-2-06 15", date)
	}
	if t < 1 {
		t = Date2TimestampFormat("1-2-06", date)
	}
	if t < 1 {
		if stringToFloat64(date) == 0 {
			return 0
		}
		t = int64((stringToFloat64(date)-25569)*86400) - 8*3600
	}
	return t
}

// 字符转Int32
func stringToFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		i = 0
	}
	return float64(i)
}
