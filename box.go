package std

import (
	"reflect"
)

// An Box is used to wrap the go interface{} in a gomobile compatible way.
type Box struct {
	// Value just holds the actual interface{}, which may be nil.
	Value interface{}
}

// String returns a string representation of the boxed type, or the empty string if box is nil or refers to nil.
func (i *Box) String() string {
	if i == nil || i.Value == nil {
		return ""
	}
	return reflect.TypeOf(i.Value).String()
}

// Set unboxes the value from the other box and sets it into this box
func (i *Box) Set(other *Box) {
	i.Value = other.Value
}

// SetString puts the value into the box
func (i *Box) SetString(v string) {
	i.Value = v
}

// SetFloat64 puts the value into the box
func (i *Box) SetFloat64(v float64) {
	i.Value = v
}

// SetInt64 puts the value into the box
func (i *Box) SetInt64(v int64) {
	i.Value = v
}

// SetBool puts the value into the box
func (i *Box) SetBool(v bool) {
	i.Value = v
}

// SetBoxSlice unwraps the slice from the box and boxes it
func (i *Box) SetBoxSlice(v *BoxSlice) {
	i.Value = v.Slice
}

// SetStrMap takes the map and unwraps the map into this box.
func (i *Box) SetStrMap(v *StrMap) {
	i.Value = v.Map
}

// IsNil checks if the box or the boxed interface is nil. Usually you don't need that.
func (i *Box) IsNil() bool {
	return i == nil || i.Value == nil
}
