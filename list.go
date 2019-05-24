package std

// A List covers an []interface{} slice with various helper methods.
type List struct {
	entries []interface{}
}

// Len returns the amount of elements in the list
func (l *List) Len() int {
	return len(l.entries)
}

// Add appends another unboxed element
func (l *List) Add(v *Box) {
	l.entries = append(l.entries, v.Unbox())
}

// AddAll appends all other elements
func (l *List) AddAll(v List) {
	l.entries = append(l.entries, v.entries...)
}

// Remove deletes the element at the given index. If out of bounds, false is returned instead of panic.
func (l *List) Remove(idx int) bool {
	if idx < 0 || idx >= len(l.entries) {
		return false
	}
	copy(l.entries[idx:], l.entries[idx+1:])
	l.entries[len(l.entries)-1] = nil //be GC friendly
	l.entries = l.entries[:len(l.entries)-1]
	return true
}

// Get returns the element at the given position or nil, if otherwise an out of bounds would be thrown
func (l *List) Get(idx int) *Box {
	if idx < 0 || idx >= len(l.entries) {
		return nil
	}
	return Wrap(l.entries[idx])
}

// Set replaces the unboxed value. If index is out of bounds, false is returned
func (l *List) Set(idx int, value *Box) bool {
	if idx < 0 || idx >= len(l.entries) {
		return false
	}
	l.entries[idx] = value.Unbox()
	return true
}

// Clear removes all entries
func (l *List) Clear() {
	// let the GC do its job
	l.entries = nil
}

// Box returns this list as a box
func (l *List) Box() *Box {
	return Wrap(l)
}

// Slice returns the underlying slice in a slice container as a defensive copy
func (l *List) Slice() *Slice {
	cpy := make([]interface{}, l.Len())
	copy(cpy, l.entries)
	return &Slice{cpy}
}
