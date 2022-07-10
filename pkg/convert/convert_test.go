package convert_test

import (
	"testing"
	"time"

	. "github.com/mreza0100/shortly/pkg/convert"
)

// Check if HourToDuration works correctly
func TestHourToDuration(t *testing.T) {
	cases := []struct {
		hour   int
		result time.Duration
	}{
		{1, 1 * time.Hour},
		{2, 2 * time.Hour},
		{3, 3 * time.Hour},
		{4, 4 * time.Hour},
		{5, 5 * time.Hour},
		{6, 6 * time.Hour},
		{7, 7 * time.Hour},
		{8, 8 * time.Hour},
		{9, 9 * time.Hour},
		{10, 10 * time.Hour},
	}

	for _, c := range cases {
		if HourToDuration(c.hour) != c.result {
			t.Errorf("HourToDuration(%d) != %d", c.hour, c.result)
		}
	}
}
