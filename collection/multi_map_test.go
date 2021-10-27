package collection

import "testing"

func Test_map(t *testing.T) {

	m := NewMultiMap()

	m.Add("张三", "1")
	m.Add("张三", "2")
	m.Add("李四", 1)
	m.Add("李四", "32")

	t.Log(m, m.Get("张三"), m.Size(), m.keys())
}
