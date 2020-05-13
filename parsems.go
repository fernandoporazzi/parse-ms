package parsems

import "math"

// ParsedMilliseconds holds the parsed content
type ParsedMilliseconds struct {
	Days         int `json:"days"`
	Hours        int `json:"hours"`
	Minutes      int `json:"minutes"`
	Seconds      int `json:"seconds"`
	Milliseconds int `json:"milliseconds"`
	Microseconds int `json:"microseconds"`
	Nanoseconds  int `json:"nanoseconds"`
}

func roundTowardsZero(m float64) int {
	if m > 0 {
		return int(math.Floor(m))
	}

	return int(math.Ceil(m))
}

// Parse takes a value in milliseconds and returns a struct with parsed data
func Parse(m float64) ParsedMilliseconds {
	ps := ParsedMilliseconds{
		Days:         roundTowardsZero(m / 86400000),
		Hours:        roundTowardsZero(m/3600000) % 24,
		Minutes:      roundTowardsZero(m/60000) % 60,
		Seconds:      roundTowardsZero(m/1000) % 60,
		Milliseconds: roundTowardsZero(m) % 1000,
		Microseconds: roundTowardsZero(m*1000) % 1000,
		Nanoseconds:  roundTowardsZero(m*1000000) % 1000,
	}

	return ps
}
