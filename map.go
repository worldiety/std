package std

// A Map can use more or less arbitrary key/value boxes and is a map[interface{}]interface{}. The key
// must be a compatible map type, like string or int.
type Map struct {
	// Map contains all unboxed keys and values
	Map map[interface{}]interface{}
}

func (m *Map) init() {
	if m.Map == nil {
		m.Map = make(map[interface{}]interface{})
	}
}

// Put replaces any unboxed key in the map with the unboxed value
func (m *Map) Put(key *Box, value *Box) {
	m.init()
	m.Map[key.Unbox()] = value.Unbox()
}

// Get returns the value in a box for the given key or nil.
func (m *Map) Get(key *Box) *Box {
	m.init()
	return Wrap(m.Map[key.Unbox()])
}

// Len returns the amount of entries in the map
func (m *Map) Len() int {
	m.init()
	return len(m.Map)
}

// Keys returns the keys as a Slice. Caution: the order of keys is not stable, so keep the StrSlice while iterating.
func (m *Map) Keys() *Slice {
	m.init()
	res := make([]interface{}, len(m.Map))
	idx := 0
	for k := range m.Map {
		res[idx] = k
		idx++
	}
	return &Slice{res}
}

// Values returns the values as a Slice
func (m *Map) Values() *Slice {
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
func (m *Map) Has(key *Box) bool {
	m.init()
	_, ok := m.Map[key.Unbox()]
	return ok
}

// Delete removes the unboxed key from the map
func (m *Map) Delete(key *Box) {
	m.init()
	delete(m.Map, key.Unbox())
}

// Box returns this Map in a box
func (m *Map) Box() *Box {
	return &Box{m}
}
