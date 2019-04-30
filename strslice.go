package std

import "fmt"

// A StringSlice is a []string and should be considered immutable.
type StrSlice struct {
	Slice []string
}

// Len returns the length of the slice
func (s *StrSlice) Len() int {
	return len(s.Slice)
}

// Get returns the value at the index, or returns an error
func (s *StrSlice) Get(idx int) (string, error) {
	if idx < 0 || idx >= len(s.Slice) {
		return "", fmt.Errorf("index out of bounds")
	}
	return s.Slice[idx], nil
}
