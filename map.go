package std

// smap is the internal map implementation
type smap struct {
	// the actual map reference
	Map map[string]interface{}
}

func (m *smap) init() {
	if m.Map == nil {
		m.Map = make(map[string]interface{})
	}
}

// Keys returns all keys a StrList, which is a defensive copy
func (m *smap) Keys() *StrList {
	m.init()
	list := &StrList{}
	list.entries = make([]interface{}, len(m.Map))
	i := 0
	for k := range m.Map {
		list.entries[i] = k
		i++
	}
	return list
}

// Put replaces any value in the map with the given interface value
func (m *smap) Put(key string, value interface{}) {
	m.init()
	m.Map[key] = value
}

// Get returns the value for the given key or nil
func (m *smap) Get(key string) interface{} {
	m.init()
	v := m.Map[key]
	return v
}

// Has checks if the key exists, not if the associated value is nil or not
func (m *smap) Has(key string) bool {
	m.init()
	_, ok := m.Map[key]
	return ok
}

//==

// A StrMap wraps a Map[string]interface{}
type StrMap struct {
	smap
}

// Put replaces any value in the map with the given interface value
func (m *StrMap) Put(key string, value *Box) {
	if value == nil {
		m.smap.Put(key, nil)
	} else {
		m.smap.Put(key, &Box{value})
	}

}

// Get returns the value for the given key or nil
func (m *StrMap) Get(key string) *Box {
	v := m.smap.Get(key)
	if v != nil {
		return &Box{v}
	}
	return nil
}

// AsInterface wraps this type into an Box
func (m *StrMap) AsInterface() *Box {
	return &Box{m}
}

//==

// A StrStrMap wraps a Map[string]string
type StrStrMap struct {
	smap
}

// Put replaces any value in the map with the given interface value
func (m *StrStrMap) Put(key string, value string) {
	m.smap.Put(key, value)
}

// Get returns the value for the given key or nil
func (m *StrStrMap) Get(key string) string {
	v := m.smap.Get(key)
	if i, ok := v.(string); ok {
		return i
	}
	return ""
}

// AsInterface wraps this type into an Box
func (m *StrStrMap) AsInterface() *Box {
	return &Box{m}
}

// AsMap converts this strstrmap into a map[string]interface{}
func (m *StrMap) AsMap() *StrMap {
	return &StrMap{m.smap}
}
