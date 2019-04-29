package std

import "fmt"

// A BoxSlice is a []interfaces{}
type BoxSlice struct {
	Slice []interface{}
}

// Len returns the length of the slice
func (s *BoxSlice) Len() int {
	return len(s.Slice)
}

// Get returns the value at the index, or returns an error
func (s *BoxSlice) Get(idx int) (*Box, error) {
	if idx < 0 || idx >= len(s.Slice) {
		return nil, fmt.Errorf("index out of bounds")
	}
	return &Box{s.Slice[idx]}, nil
}

//==

// A StringSlice is a []string
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
