package std

import "fmt"

// IsString checks if the boxed interface represents a string
func (i *Box) IsString() bool {
	if _, ok := i.Value.(string); ok {
		return true
	}
	return false
}

// AsString performs a type assertion to a string.
func (i *Box) AsString() (string, error) {
	if v, ok := i.Value.(string); ok {
		return v, nil
	}
	return "", fmt.Errorf("not a string: %s", i.String())
}

// IsInt8 checks if the boxed interface represents a signed int8 or a byte
func (i *Box) IsInt8() bool {
	if _, ok := i.Value.(int8); ok {
		return true
	}
	return false
}

// IsInt16 checks if the boxed interface represents a signed int16 or a short
func (i *Box) IsInt16() bool {
	if _, ok := i.Value.(int16); ok {
		return true
	}
	return false
}

// IsInt32 checks if the boxed interface represents a signed int32 or an int
func (i *Box) IsInt32() bool {
	if _, ok := i.Value.(int32); ok {
		return true
	}
	return false
}

// IsInt32 checks if the boxed interface represents a signed int64 or a long
func (i *Box) IsInt64() bool {
	if _, ok := i.Value.(int64); ok {
		return true
	}
	return false
}

// IsBoxSlice checks if the type is one of []interface{}|*[]interface{}
func (i *Box) IsBoxSlice() bool {
	switch i.Value.(type) {
	case []interface{}:
		return true
	case *[]interface{}:
		return true
	case *ArrayList:
		return true
	}
	return false
}

// AsBoxSlice wraps the underlying slice into a BoxSlice or returns nil
func (i *Box) AsBoxSlice() *BoxSlice {
	switch t := i.Value.(type) {
	case []interface{}:
		return &BoxSlice{t}
	case *[]interface{}:
		return &BoxSlice{*t}

	}
	return nil
}

// IsStrMap checks if the type is map[string]interface{}
func (i *Box) IsStrMap() bool {
	if _, ok := i.Value.(map[string]interface{}); ok {
		return true
	}
	return false
}

// AsStrMap tries to wrap a map[string]interface{} or returns nil
func (i *Box) AsStrMap() *StrMap {
	switch t := i.Value.(type) {
	case map[string]interface{}:
		return &StrMap{smap{t}}

	}
	return nil
}
