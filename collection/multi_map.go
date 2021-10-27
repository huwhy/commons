package collection

type MultiMap struct {
	values map[interface{}][]interface{}
}

func NewMultiMap() *MultiMap {
	return &MultiMap{values: make(map[interface{}][]interface{})}
}

func (m *MultiMap) Add(key, value interface{}) {
	if v, ok := m.values[key]; ok {
		v = append(v, value)
		m.values[key] = v
	} else {
		v = make([]interface{}, 0)
		v = append(v, value)
		m.values[key] = v
	}
}

func (m *MultiMap) Get(key interface{}) []interface{} {
	return m.values[key]
}

func (m *MultiMap) Remove(key interface{}) []interface{} {
	if v, ok := m.values[key]; ok {
		delete(m.values, key)
		return v
	} else {
		return nil
	}
}

func (m *MultiMap) Size() int {
	return len(m.values)
}

func (m *MultiMap) keys() []interface{} {
	keys := make([]interface{}, m.Size())
	var i = 0
	for k := range m.values {
		keys[i] = k
		i++
	}
	return keys
}
