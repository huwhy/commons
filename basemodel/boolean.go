package basemodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Boolean int8

const (
	True  = Boolean(uint8(1))
	False = Boolean(uint8(0))
)

// MarshalJSON 实现序列化接口
func (m Boolean) MarshalJSON() ([]byte, error) {
	if m == True {
		return json.Marshal(true)
	} else {
		return json.Marshal(false)
	}
}

// UnmarshalJSON 反序列化
func (m *Boolean) UnmarshalJSON(data []byte) (err error) {
	if data == nil || len(data) == 0 {
		return nil
	}
	v := string(data)
	v = strings.ReplaceAll(v, "\"", "")
	if v == "1" || v == "true" {
		*m = True
	} else if v == "0" || v == "false" {
		*m = False
	}
	return
}

// Scan 实现 sql.Scanner 接口
func (m *Boolean) Scan(value interface{}) error {
	t, ok := value.(uint8)
	if !ok {
		return errors.New(fmt.Sprint("Failed to uint8 value:", value))
	}
	if t == 1 || t == 0 {
		*m = Boolean(t)
	} else {
		return errors.New(fmt.Sprint("Failed to Boolean value:", value))
	}
	return nil
}

// Value 实现 driver.Valuer 接口
func (m Boolean) Value() (driver.Value, error) {
	v := uint8(m)
	return v, nil
}

// String
func (m Boolean) String() string {
	if m == True {
		return "true"
	} else if m == False {
		return "false"
	}
	return "nil"
}
