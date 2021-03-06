package datetimes

import (
	"github.com/huwhy/commons/constant"
	"github.com/huwhy/commons/util/times"
	"time"
)

func ParseTime(value string) *DateTime {
	return Parse(value, constant.DateTimeFormat)
}

func Parse(value, format string) *DateTime {
	time, err := time.ParseInLocation(format, value, time.Local)
	if err != nil {
		panic(err)
	}
	v := DateTime(time)
	return &v
}

func Now() DateTime {
	return DateTime(time.Now())
}
func Now2() *DateTime {
	v := DateTime(time.Now())
	return &v
}

func ValueOf(t time.Time) DateTime {
	return DateTime(t)
}

func ValueOfString(v string) DateTime {
	dateTime, err := times.ParseDateTime(v)
	if err != nil {
		panic(err)
	}
	return DateTime(dateTime)
}

func FormatWithDayStartTime(date *DateTime) string {
	return time.Time(*date).Format(constant.DateFormat) + " 00:00:00"
}

func FormatWithDayEndTime(date *DateTime) string {
	return time.Time(*date).Format(constant.DateFormat) + " 23:59:59"
}
