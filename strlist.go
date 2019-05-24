package std

// A StrList covers a []string slice with various helper methods.
type StrList struct {
	entries []string
}

// Len returns the amount of elements in the list
func (l *StrList) Len() int {
	return len(l.entries)
}

// Add appends another unboxed element
func (l *StrList) Add(v string) {
	l.entries = append(l.entries, v)
}

// AddAll appends all other elements
func (l *StrList) AddAll(v StrList) {
	l.entries = append(l.entries, v.entries...)
}

// Remove deletes the element at the given index. If out of bounds, false is returned instead of panic.
func (l *StrList) Remove(idx int) bool {
	if idx < 0 || idx >= len(l.entries) {
		return false
	}
	copy(l.entries[idx:], l.entries[idx+1:])
	l.entries[len(l.entries)-1] = "" //be GC friendly
	l.entries = l.entries[:len(l.entries)-1]
	return true
}

// Get returns the element at the given position or the empty string
func (l *StrList) Get(idx int) string {
	if idx < 0 || idx >= len(l.entries) {
		return ""
	}
	return l.entries[idx]
}

// Set replaces the unboxed value. If index is out of bounds, false is returned
func (l *StrList) Set(idx int, value string) bool {
	if idx < 0 || idx >= len(l.entries) {
		return false
	}
	l.entries[idx] = value
	return true
}

// Clear removes all entries
func (l *StrList) Clear() {
	// let the GC do its job
	l.entries = nil
}

// Box returns this list as a box
func (l *StrList) Box() *Box {
	return Wrap(l)
}

// Slice returns the underlying slice in a slice container as a defensive copy
func (l *StrList) Slice() *StrSlice {
	cpy := make([]string, l.Len())
	copy(cpy, l.entries)
	return &StrSlice{cpy}
}
