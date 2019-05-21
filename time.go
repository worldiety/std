package std

import "time"

// Time is a wrapper around time.Time
type Time struct {
	Value time.Time
}

func Now() *Time {
	return &Time{time.Now()}
}

func Unix(sec int64, nsec int64) *Time {
	return &Time{time.Unix(sec, nsec)}
}

// Duration is a wrapper around time.Duration
type Duration struct {
	Value time.Duration
}

// NewDuration creates a duration based on milliseconds
func NewDuration(milliseconds float64) *Duration {
	return &Duration{time.Duration(milliseconds * 1000 * 1000)} //milli * micro * nano
}
