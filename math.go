package std

import "math"

//IsNaN returns true, if the given float is not a number
func IsNaN(f float64) bool {
	return math.IsNaN(f)
}
