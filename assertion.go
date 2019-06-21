package std

import (
	"context"
	"fmt"
	"strconv"
)

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

// AsInt8 performs a type assertion to an int8.
func (i *Box) AsInt8() (int8, error) {
	if v, ok := i.Value.(int8); ok {
		return v, nil
	}
	return 0, fmt.Errorf("not int8: %s", i.String())
}

// IsInt16 checks if the boxed interface represents a signed int16 or a short
func (i *Box) IsInt16() bool {
	if _, ok := i.Value.(int16); ok {
		return true
	}
	return false
}

// AsInt16 performs a type assertion to an int16.
func (i *Box) AsInt16() (int16, error) {
	if v, ok := i.Value.(int16); ok {
		return v, nil
	}
	return 0, fmt.Errorf("not int16: %s", i.String())
}

// IsInt32 checks if the boxed interface represents a signed int32 or an int
func (i *Box) IsInt32() bool {
	if _, ok := i.Value.(int32); ok {
		return true
	}
	return false
}

// AsInt32 performs a type assertion to an int32.
func (i *Box) AsInt32() (int32, error) {
	if v, ok := i.Value.(int32); ok {
		return v, nil
	}
	return 0, fmt.Errorf("not int32: %s", i.String())
}

// IsInt64 checks if the boxed interface represents a signed int64 or a long
func (i *Box) IsInt64() bool {
	if _, ok := i.Value.(int64); ok {
		return true
	}
	return false
}

// AsInt64 performs a type assertion to an int64.
func (i *Box) AsInt64() (int64, error) {
	if v, ok := i.Value.(int64); ok {
		return v, nil
	}
	return 0, fmt.Errorf("not int64: %s", i.String())
}

// IsBoxSlice checks if the type is one of []interface{}|*[]interface{}|*Slice
func (i *Box) IsSlice() bool {
	if i == nil {
		return false
	}
	switch i.Value.(type) {
	case []interface{}:
		return true
	case *[]interface{}:
		return true
	case *Slice:
		return true
	}
	return false
}

// AsSlice wraps the underlying slice into a Slice or returns nil
func (i *Box) AsSlice() *Slice {
	if i == nil {
		return nil
	}
	switch t := i.Value.(type) {
	case []interface{}:
		return &Slice{t}
	case *[]interface{}:
		return &Slice{*t}
	case *Slice:
		return t
	}
	return nil
}

// IsMap checks if the type is map[string]interface{}
func (i *Box) IsMap() bool {
	if i == nil {
		return false
	}
	if _, ok := i.Value.(map[string]interface{}); ok {
		return true
	}
	return false
}

// AsMap tries to wrap a map[interface{}]interface{} or returns nil
func (i *Box) AsMap() *Map {
	if i == nil {
		return nil
	}
	switch t := i.Value.(type) {
	case map[interface{}]interface{}:
		return &Map{t}
	case *Map:
		return t
	}
	return nil
}

// IsStrMap checks if the type is map[string]interface{}
func (i *Box) IsStrMap() bool {
	if i == nil {
		return false
	}
	if _, ok := i.Value.(map[string]interface{}); ok {
		return true
	}
	return false
}

// AsStrMap tries to wrap a map[string]interface{} or returns nil
func (i *Box) AsStrMap() *StrMap {
	if i == nil {
		return nil
	}
	return duckTypeStrMap(i.Value)
}

// AsContext tries to wrap a context.Context or asserts a std.Context, otherwise returns nil
func (i *Box) AsContext() *Context {
	if i == nil {
		return nil
	}
	switch t := i.Value.(type) {
	case context.Context:
		return &Context{Value: nil}
	case *Context:
		return t
	}
	return nil
}

func (i *Box) Int() int {
	return int(i.Int64())
}

func (i *Box) Int32() int32 {
	return int32(i.Int64())
}

// Int64 tries to interpret any string or number gracefully into an int64. If it fails, returns just 0
func (i *Box) Int64() int64 {
	if i == nil {
		return 0
	}
	return duckTypeInt64(i.Value)
}

// Float tries to interpret any number gracefully into a Float. If it fails, returns NaN
func (i *Box) Float64() float64 {
	if i == nil {
		return 0
	}
	switch t := i.Value.(type) {
	case int8:
		return float64(t)
	case uint8:
		return float64(t)
	case int16:
		return float64(t)
	case uint16:
		return float64(t)
	case int32:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case int64:
		return float64(t)
	case int:
		return float64(t)
	case string:
		i, _ := strconv.ParseFloat(t, 64)
		return i
	case float32:
		return float64(t)
	case float64:
		return t
	case bool:
		if t {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// Bool tries to interpret any number or string gracefully into a boolean. If it fails, returns false.
func (i *Box) Bool() bool {
	if i == nil {
		return false
	}
	switch t := i.Value.(type) {
	case string:
		b, _ := strconv.ParseBool(t)
		return b
	default:
		b := i.Int()
		if b == 1 {
			return true
		}
		return false
	}
}

// String returns a string representation of the boxed type, or the empty string if box is nil or refers to nil.
func (i *Box) String() string {
	if i == nil {
		return ""
	}
	return duckTypeString(i.Value)
}

func duckTypeString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%+v", v)
}

func duckTypeInt64(v interface{}) int64 {
	switch t := v.(type) {
	case int8:
		return int64(t)
	case uint8:
		return int64(t)
	case int16:
		return int64(t)
	case uint16:
		return int64(t)
	case int32:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case int64:
		return t
	case int:
		return int64(t)
	case string:
		i, _ := strconv.ParseInt(t, 10, 64)
		return i
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case bool:
		if t {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func duckTypeStrMap(v interface{}) *StrMap {
	switch t := v.(type) {
	case map[string]interface{}:
		return &StrMap{t}
	case *StrMap:
		return t
	}
	return nil
}
