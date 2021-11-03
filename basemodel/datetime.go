package basemodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/huwhy/commons/constant"
	"time"
)

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	ts := time.Time(t)
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

// Scan 实现 sql.Scanner 接口
func (t *DateTime) Scan(value interface{}) error {
	ti, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to time.Time value:", value))
	}
	*t = DateTime(ti)
	return nil
}

// Value 实现 driver.Valuer 接口
func (t DateTime) Value() (driver.Value, error) {
	ti := time.Time(t)
	return ti, nil
}

func (t DateTime) String() string {
	return "\"" + t.Format() + "\""
}

func (t *DateTime) Format() string {
	return (time.Time(*t)).Format(constant.DateTimeFormat)
}

func (t *DateTime) WithFormat(format string) string {
	return (time.Time(*t)).Format(format)
}

func (t DateTime) ToTime() time.Time {
	return time.Time(t)
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
