package collection

type multiMap struct {
	values map[interface{}][]interface{}
}

func NewMultiMap() *multiMap {
	return &multiMap{values: make(map[interface{}][]interface{})}
}

func (m *multiMap) Add(key, value interface{}) {
	if v, ok := m.values[key]; ok {
		v = append(v, value)
		m.values[key] = v
	} else {
		v = make([]interface{}, 0)
		v = append(v, value)
		m.values[key] = v
	}
}

func (m *multiMap) Get(key interface{}) []interface{} {
	return m.values[key]
}

func (m *multiMap) Remove(key interface{}) []interface{} {
	if v, ok := m.values[key]; ok {
		delete(m.values, key)
		return v
	} else {
		return nil
	}
}

func (m *multiMap) Size() int {
	return len(m.values)
}

func (m *multiMap) keys() []interface{} {
	keys := make([]interface{}, m.Size())
	var i = 0
	for k := range m.values {
		keys[i] = k
		i++
	}
	return keys
}
