package datetimes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"huwhy.cn/commons/constant"
	"huwhy.cn/commons/util/times"
	"time"
)

type DateTime time.Time

func (t *DateTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	ts := time.Time(*t)
	b := make([]byte, 0, len(constant.DateTimeFormat)+2)
	b = append(b, '"')
	if !ts.IsZero() && ts.Unix() > 0 {
		b = ts.AppendFormat(b, constant.DateTimeFormat)
	}
	b = append(b, '"')
	return b, nil
}

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	value := string(data)
	if value != "" {
		now, err := time.ParseInLocation(`"`+constant.DateTimeFormat+`"`, value, time.Local)
		*t = DateTime(now)
		return err
	}
	return
}

// 实现 sql.Scanner 接口
func (j *DateTime) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to time.Time value:", value))
	}
	*j = DateTime(t)
	return nil
}

// 实现 driver.Valuer 接口
func (j DateTime) Value() (driver.Value, error) {
	t := time.Time(j)
	return t, nil
}

func (t *DateTime) Format() string {
	return (time.Time(*t)).Format(constant.DateTimeFormat)
}

func (t *DateTime) WithFormat(format string) string {
	return (time.Time(*t)).Format(format)
}

func (t *DateTime) DayOfMonth() int {
	return (time.Time(*t)).Day()
}

func (t *DateTime) MonthOfYear() int {
	return int((time.Time(*t)).Month())
}

func (t *DateTime) Year() int {
	return (time.Time(*t)).Year()
}

func (t *DateTime) AddDays(days int) *DateTime {
	tm := time.Time(*t).Add(time.Hour * time.Duration(days*24))
	*t = DateTime(tm)
	return t
}

func (t *DateTime) AddHours(hours int) *DateTime {
	tm := time.Time(*t).Add(time.Hour * time.Duration(hours))
	*t = DateTime(tm)
	return t
}

func Now() DateTime {
	return DateTime(time.Now())
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
