package times

import (
	"git.huwhy.cn/commons/constant"
	"strconv"
	"time"
)

func PlusDays(date time.Time, days int) time.Time {
	t := date.Unix()
	t += int64(constant.OneDaySeconds * days)
	return time.Unix(t, 0)
}

func Parse(value, format string) (time.Time, error) {
	return time.ParseInLocation(format, value, time.Local)
}

func ParseDateTime(value string) (time.Time, error) {
	return Parse(value, constant.DateTimeFormat)
}
func FormatDate(t time.Time) string {
	return t.Format(constant.DateFormat)
}
func FormatDateTime(t time.Time) string {
	return t.Format(constant.DateTimeFormat)
}

func TodayString() string {
	return time.Now().Format(constant.DateFormat)
}
func TodayTimeString() string {
	return time.Now().Format(constant.DateTimeFormat)
}

func NowNumber() int64 {
	s := time.Now().Format(constant.DateTimeNumberFormat)
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func NowHour() int64 {
	s := time.Now().Format(constant.DateTimeHourNumberFormat)
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func PlusMonth(t time.Time) time.Time {
	var v = time.Date(t.Year(), t.Month()+1, t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	return v
}

func GetMonthStart(t time.Time) time.Time {
	var v = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	return v
}

func GetMonthEnd(t time.Time) time.Time {
	var v = time.Date(t.Year(), t.Month()+1, 1, 23, 59, 59, 0, time.Local)
	return v.Add(time.Hour * -24)
}

func KeepYMD(t time.Time) time.Time {
	var v = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return v
}
