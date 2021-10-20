package basemodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Boolean int8

var (
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
	switch v := value.(type) {
	case int64:
		if v == 1 || v == 0 {
			*m = Boolean(v)
		}
	case []uint8:
		vv := string(v)
		vv = strings.ReplaceAll(vv, "\"", "")
		if vv == "1" || vv == "true" {
			*m = True
		} else if vv == "0" || vv == "false" {
			*m = False
		}
	case uint8:
		if v == 1 || v == 0 {
			*m = Boolean(v)
		}
	default:
		return errors.New(fmt.Sprint("Failed to uint8 value:", value))
	}
	return nil
}

// Value 实现 driver.Valuer 接口
func (m Boolean) Value() (driver.Value, error) {
	if m == True {
		return true, nil
	} else if m == False {
		return false, nil
	}
	return nil, nil
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

func (m Boolean) IsTrue() bool {
	return m == True
}
