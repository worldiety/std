package std

import "fmt"

// A Slice is a []interfaces{} and should be considered immutable. Internally boxes are removed.
type Slice struct {
	Slice []interface{}
}

// Len returns the length of the slice
func (s *Slice) Len() int {
	return len(s.Slice)
}

// Get returns the value at the index, or returns an error
func (s *Slice) Get(idx int) (*Box, error) {
	if idx < 0 || idx >= len(s.Slice) {
		return nil, fmt.Errorf("index out of bounds")
	}
	return &Box{s.Slice[idx]}, nil
}
