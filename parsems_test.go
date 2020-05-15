package parsems

import (
	"testing"
)

func TestParse(t *testing.T) {
	var entries = []struct {
		input float64
		want  ParsedMilliseconds
	}{
		{
			float64(60500.345678),
			ParsedMilliseconds{
				Days:         0,
				Hours:        0,
				Minutes:      1,
				Seconds:      0,
				Milliseconds: 500,
				Microseconds: 345,
				Nanoseconds:  0,
			},
		},
		{
			float64(55000),
			ParsedMilliseconds{
				Days:         0,
				Hours:        0,
				Minutes:      0,
				Seconds:      55,
				Milliseconds: 0,
				Microseconds: 0,
				Nanoseconds:  0,
			},
		},
		{
			float64(67000),
			ParsedMilliseconds{
				Days:         0,
				Hours:        0,
				Minutes:      1,
				Seconds:      7,
				Milliseconds: 0,
				Microseconds: 0,
				Nanoseconds:  0,
			},
		},
		{
			float64(4020000),
			ParsedMilliseconds{
				Days:         0,
				Hours:        1,
				Minutes:      7,
				Seconds:      0,
				Milliseconds: 0,
				Microseconds: 0,
				Nanoseconds:  0,
			},
		},
		{
			float64(144000000),
			ParsedMilliseconds{
				Days:         1,
				Hours:        16,
				Minutes:      0,
				Seconds:      0,
				Milliseconds: 0,
				Microseconds: 0,
				Nanoseconds:  0,
			},
		},
		{
			float64(3596400000),
			ParsedMilliseconds{
				Days:         41,
				Hours:        15,
				Minutes:      0,
				Seconds:      0,
				Milliseconds: 0,
				Microseconds: 0,
				Nanoseconds:  0,
			},
		},
		{
			float64(0.000543),
			ParsedMilliseconds{
				Days:         0,
				Hours:        0,
				Minutes:      0,
				Seconds:      0,
				Milliseconds: 0,
				Microseconds: 0,
				Nanoseconds:  543,
			},
		},
	}

	for _, entry := range entries {
		got := Parse(entry.input)

		if got != entry.want {
			t.Errorf("Expected %v to be equal %v", got, entry.want)
		}
	}
}
