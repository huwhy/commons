package collection

import (
	"reflect"
	"strings"
)

func StructToMap(m interface{}) map[string]interface{} {
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		name = strings.ToLower(name[:1]) + name[1:]
		data[name] = v.Field(i).Interface()
	}
	return data
}
