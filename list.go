package std

// A list is the private basic list implementation
type list struct {
	entries []interface{}
}

// Size returns the amount of elements in the list
func (l *list) Size() int {
	return len(l.entries)
}

// Add appends another element
func (l *list) Add(v interface{}) {
	l.entries = append(l.entries, v)
}

// AddAll appends all other elements
func (l *list) AddAll(v []interface{}) {
	l.entries = append(l.entries, v...)
}

// Remove deletes the element at the given index. If out of bounds, false is returned instead of panic.
func (l *list) Remove(idx int) bool {
	if idx < 0 || idx >= len(l.entries) {
		return false
	}
	copy(l.entries[idx:], l.entries[idx+1:])
	l.entries[len(l.entries)-1] = nil
	l.entries = l.entries[:len(l.entries)-1]
	return true
}

// Get returns the element at the given position or nil, if otherwise an out of bounds would be thrown
func (l *list) Get(idx int) interface{} {
	if idx < 0 || idx >= len(l.entries) {
		return nil
	}
	return l.entries[idx]
}

//==

// An ArrayList covers an interface slice
type ArrayList struct {
	list
}

// Add appends another element
func (l *ArrayList) Add(v *Box) {
	l.list.Add(v)
}

// AddAll appends all other elements
func (l *ArrayList) AddAll(v ArrayList) {
	l.list.AddAll(v.entries)
}

// Get returns the element at the given position or nil, if otherwise an out of bounds would be thrown
func (l *ArrayList) Get(idx int) *Box {
	return l.list.Get(idx).(*Box)
}

//==

// A StrList is an ArrayList of strings
type StrList struct {
	list
}

// Add appends another element
func (l *StrList) Add(v string) {
	l.list.Add(v)
}

// AddAll appends all other elements
func (l *StrList) AddAll(v StrList) {
	l.list.AddAll(v.entries)
}

// Get returns the element at the given position or the empty string, if otherwise an out of bounds would be thrown
func (l *StrList) Get(idx int) string {
	return l.list.Get(idx).(string)
}
