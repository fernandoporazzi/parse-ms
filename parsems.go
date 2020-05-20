package parsems

import "math"

// ParsedMilliseconds holds the parsed content
type ParsedMilliseconds struct {
	Days         float64 `json:"days"`
	Hours        float64 `json:"hours"`
	Minutes      float64 `json:"minutes"`
	Seconds      float64 `json:"seconds"`
	Milliseconds float64 `json:"milliseconds"`
	Microseconds float64 `json:"microseconds"`
	Nanoseconds  float64 `json:"nanoseconds"`
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
		Days:         float64(roundTowardsZero(m / 86400000)),
		Hours:        float64(roundTowardsZero(m/3600000) % 24),
		Minutes:      float64(roundTowardsZero(m/60000) % 60),
		Seconds:      float64(roundTowardsZero(m/1000) % 60),
		Milliseconds: float64(roundTowardsZero(m) % 1000),
		Microseconds: float64(roundTowardsZero(m*1000) % 1000),
		Nanoseconds:  float64(roundTowardsZero(m*1000000) % 1000),
	}

	return ps
}
