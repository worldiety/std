package std

// A StrMap is a map[string]interface{} and holds arbitrary unboxed values.
type StrMap struct {
	// Map contains all string keys and values
	Map map[string]interface{}
}

func (m *StrMap) init() {
	if m.Map == nil {
		m.Map = make(map[string]interface{})
	}
}

// Put replaces any unboxed key in the map with the unboxed value
func (m *StrMap) Put(key string, value *Box) {
	m.init()
	m.Map[key] = value.Unbox()
}

// Get returns the value in a box for the given key or nil
func (m *StrMap) Get(key string) *Box {
	m.init()
	return Wrap(m.Map[key])
}

// Len returns the amount of entries in the map
func (m *StrMap) Len() int {
	m.init()
	return len(m.Map)
}

// Keys returns the keys as a Slice. Caution: the order of keys is not stable, so keep the StrSlice while iterating.
func (m *StrMap) Keys() *StrSlice {
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
func (m *StrMap) Values() *Slice {
	m.init()
	res := make([]interface{}, len(m.Map))
	idx := 0
	for _, v := range m.Map {
		res[idx] = v
		idx++
	}
	return &Slice{res}
}

// Has checks if the unboxed key exists, not if the associated value is nil or not
func (m *StrMap) Has(key string) bool {
	m.init()
	_, ok := m.Map[key]
	return ok
}

// Delete removes the unboxed key from the map
func (m *StrMap) Delete(key string) {
	m.init()
	delete(m.Map, key)
}

// Box returns this Map in a box
func (m *StrMap) Box() *Box {
	return &Box{m}
}
