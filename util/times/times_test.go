package times

import (
	"testing"
	"time"
)

func TestGetMonthEnd(t *testing.T) {
	tm := GetMonthStart(time.Now())
	t.Log(tm)
	t.Log(GetMonthEnd(time.Now()))
	t.Log(KeepYMD(time.Now()))
	t.Log(PlusMonth(time.Now()))
}
