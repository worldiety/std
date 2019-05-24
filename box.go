package std

import "reflect"

// An Box is used to wrap the go interface{} in a gomobile compatible way.
type Box struct {
	// Value just holds the actual interface{}, which may be nil.
	Value interface{}
}

// Wrap is not available in gomobile and is a factory for boxes and returns a nil box if value is nil.
// It is not called NewBox() because the unsupported type would create a private constructor.
func Wrap(value interface{}) *Box {
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
func (i *Box) Set(other *Box) *Box {
	i.Value = other.Unbox()
	return i
}

// SetString puts the value into the box and returns the box.
func (i *Box) SetString(v string) *Box {
	i.Value = v
	return i
}

// SetFloat64 puts the value into the box and returns the box.
func (i *Box) SetFloat64(v float64) *Box {
	i.Value = v
	return i
}

// SetFloat32 puts the value into the box and returns the box.
func (i *Box) SetFloat32(v float32) *Box {
	i.Value = v
	return i
}

// SetInt64 puts the value into the box and returns the box.
func (i *Box) SetInt64(v int64) *Box {
	i.Value = v
	return i
}

// SetInt32 puts the value into the box and returns the box.
func (i *Box) SetInt32(v int32) *Box {
	i.Value = v
	return i
}

// SetInt16 puts the value into the box and returns the box.
func (i *Box) SetInt16(v int16) *Box {
	i.Value = v
	return i
}

// SetInt8 puts the value into the box and returns the box.
func (i *Box) SetInt8(v int8) *Box {
	i.Value = v
	return i
}

// SetInt puts the value into the box and returns the box.
func (i *Box) SetInt(v int) *Box {
	i.Value = v
	return i
}

// SetBool puts the value into the box and returns the box.
func (i *Box) SetBool(v bool) *Box {
	i.Value = v
	return i
}

// SetSlice unwraps the slice from the box and boxes it and returns the box.
func (i *Box) SetSlice(v *Slice) *Box {
	i.Value = v.Slice
	return i
}

// SetStrSlice unwraps the slice from the box and boxes it and returns the box.
func (i *Box) SetStrSlice(v *StrSlice) *Box {
	i.Value = v.Slice
	return i
}

// SetStrMap takes the map and unwraps the map into this box and returns the box.
func (i *Box) SetStrMap(v *StrMap) *Box {
	i.Value = v.Map
	return i
}

// SetMap takes the map and unwraps the map into this box and returns the box.
func (i *Box) SetMap(v *Map) *Box {
	i.Value = v.Map
	return i
}

// IsNil checks if the box or the boxed interface is nil. Usually you don't need that.
func (i *Box) IsNil() bool {
	return i == nil || i.Value == nil
}
