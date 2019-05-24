package std

// A StrStrMap is a map[string]string and holds string values.
type StrStrMap struct {
	// Map contains all string keys and values
	Map map[string]string
}

func (m *StrStrMap) init() {
	if m.Map == nil {
		m.Map = make(map[string]string)
	}
}

// Put replaces any unboxed key in the map with the unboxed value
func (m *StrStrMap) Put(key string, value string) {
	m.init()
	m.Map[key] = value
}

// Get returns the value in a box for the given key or empty string
func (m *StrStrMap) Get(key string) string {
	m.init()
	return m.Map[key]
}

// Len returns the amount of entries in the map
func (m *StrStrMap) Len() int {
	m.init()
	return len(m.Map)
}

// Keys returns the keys as a Slice. Caution: the order of keys is not stable, so keep the StrSlice while iterating.
func (m *StrStrMap) Keys() *StrSlice {
	m.init()
	res := make([]string, len(m.Map))
	idx := 0
	for k := range m.Map {
		res[idx] = k
		idx++
	}
	return &StrSlice{res}
}

// Values returns the values as a Slice
func (m *StrStrMap) Values() *StrSlice {
	m.init()
	res := make([]string, len(m.Map))
	idx := 0
	for _, v := range m.Map {
		res[idx] = v
		idx++
	}
	return &StrSlice{res}
}

// Has checks if the unboxed key exists, not if the associated value is nil or not
func (m *StrStrMap) Has(key string) bool {
	m.init()
	_, ok := m.Map[key]
	return ok
}

// Delete removes the unboxed key from the map
func (m *StrStrMap) Delete(key string) {
	m.init()
	delete(m.Map, key)
}

// Box returns this Map in a box
func (m *StrStrMap) Box() *Box {
	return &Box{m}
}
