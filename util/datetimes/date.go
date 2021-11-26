package datetimes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/huwhy/commons/constant"
	"time"
)

type Date time.Time

func (t Date) MarshalJSON() ([]byte, error) {
	ts := time.Time(t)
	b := make([]byte, 0, len(constant.DateFormat)+2)
	b = append(b, '"')
	if !ts.IsZero() && ts.Unix() > 0 {
		b = ts.AppendFormat(b, constant.DateFormat)
	}
	b = append(b, '"')
	return b, nil
}

func (t *Date) UnmarshalJSON(data []byte) (err error) {
	value := string(data)
	if value != "" {
		now, err := time.ParseInLocation(`"`+constant.DateFormat+`"`, value, time.Local)
		*t = Date(now)
		return err
	}
	return
}

// Scan 实现 sql.Scanner 接口
func (t *Date) Scan(value interface{}) error {
	ti, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to time.Time value:", value))
	}
	*t = Date(ti)
	return nil
}

// Value 实现 driver.Valuer 接口
func (t Date) Value() (driver.Value, error) {
	ti := time.Time(t)
	return ti, nil
}

func (t Date) String() string {
	return "\"" + t.Format() + "\""
}

func (t *Date) Format() string {
	return (time.Time(*t)).Format(constant.DateFormat)
}

func (t *Date) WithFormat(format string) string {
	return (time.Time(*t)).Format(format)
}

func (t Date) ToTime() time.Time {
	return time.Time(t)
}

func (t *Date) DayOfMonth() int {
	return (time.Time(*t)).Day()
}

func (t *Date) MonthOfYear() int {
	return int((time.Time(*t)).Month())
}

func (t *Date) Year() int {
	return (time.Time(*t)).Year()
}

func (t *Date) AddDays(days int) *Date {
	tm := time.Time(*t).Add(time.Hour * time.Duration(days*24))
	*t = Date(tm)
	return t
}
