package convert_test

import (
	"testing"
	"time"

	. "github.com/mreza0100/shortly/pkg/convert"
)

func TestHourToDuration(t *testing.T) {
	if HourToDuration(1) != time.Hour {
		t.Errorf("HourToDuration(1) != time.Hour")
	}
}
