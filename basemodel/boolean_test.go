package basemodel

import (
	"encoding/json"
	"testing"
	"time"
)

type A struct {
	Status  *Boolean  `json:"status"`
	Created *DateTime `json:"created"`
}

func Test_boolean_json(t *testing.T) {
	b, e := json.Marshal(False)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(string(b))
	}

	b, e = json.Marshal(True)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(string(b))
	}
	var a = &A{}
	dateTime := DateTime(time.Now())
	a.Created = &dateTime
	b, e = json.Marshal(a)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(string(b))
	}
	var s = "{\"status\":null,\"created\":\"2021-10-20 09:19:24\"}"
	e = json.Unmarshal([]byte(s), a)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(a, a.Status)
	}
}

func Test_datetime(t *testing.T) {
	m := DateTime(time.Now())
	a := &m
	t.Log(m, a)
}

func Test_boolean(t *testing.T) {
	var s = "{\"status\":null,\"created\":\"2021-10-20 09:19:24\"}"
	var a = &A{}
	e := json.Unmarshal([]byte(s), a)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(a, a.Status)
	}
}
