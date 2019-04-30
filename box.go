package std

import "reflect"

// An Box is used to wrap the go interface{} in a gomobile compatible way.
type Box struct {
	// Value just holds the actual interface{}, which may be nil.
	Value interface{}
}

// NewBox is not available in gomobile and is a factory for boxes and returns a nil box if value is nil
func NewBox(value interface{}) *Box {
	if value == nil {
		return nil
	}
	return &Box{value}
}

// Unbox is not available in gomobile and is a convenience method which also works on nil-boxes
func (i *Box) Unbox() interface{} {
	if i == nil {
		return nil
	}
	return i.Value
}

// TypeName returns a string representation of the boxed type
func (i *Box) TypeName() string {
	return reflect.TypeOf(i.Value).String()
}

// Set unboxes the value from the other box and sets it into this box
func (i *Box) Set(other *Box) {
	i.Value = other.Unbox()
}

// SetString puts the value into the box
func (i *Box) SetString(v string) {
	i.Value = v
}

// SetFloat64 puts the value into the box
func (i *Box) SetFloat64(v float64) {
	i.Value = v
}

// SetFloat32 puts the value into the box
func (i *Box) SetFloat32(v float32) {
	i.Value = v
}

// SetInt64 puts the value into the box
func (i *Box) SetInt64(v int64) {
	i.Value = v
}

// SetInt32 puts the value into the box
func (i *Box) SetInt32(v int32) {
	i.Value = v
}

// SetInt16 puts the value into the box
func (i *Box) SetInt16(v int16) {
	i.Value = v
}

// SetInt8 puts the value into the box
func (i *Box) SetInt8(v int8) {
	i.Value = v
}

// SetInt puts the value into the box
func (i *Box) SetInt(v int) {
	i.Value = v
}

// SetBool puts the value into the box
func (i *Box) SetBool(v bool) {
	i.Value = v
}

// SetSlice unwraps the slice from the box and boxes it
func (i *Box) SetSlice(v *Slice) {
	i.Value = v.Slice
}

// SetStrSlice unwraps the slice from the box and boxes it
func (i *Box) SetStrSlice(v *StrSlice) {
	i.Value = v.Slice
}

// SetStrMap takes the map and unwraps the map into this box.
func (i *Box) SetStrMap(v *StrMap) {
	i.Value = v.Map
}

// SetMap takes the map and unwraps the map into this box.
func (i *Box) SetMap(v *Map) {
	i.Value = v.Map
}

// IsNil checks if the box or the boxed interface is nil. Usually you don't need that.
func (i *Box) IsNil() bool {
	return i == nil || i.Value == nil
}
