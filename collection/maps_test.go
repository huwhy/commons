package collection

import (
	"testing"
	"time"
)

func TestStructToMap(t *testing.T) {
	user := user{5, "zhangsan", "pwd", time.Now()}
	data := StructToMap(user)
	t.Log(data)
}

type user struct {
	Id        int64
	UserName  string
	Password  string
	LoginTime time.Time
}
