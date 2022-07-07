package convert

import "time"

func HourToDuration(h int) time.Duration {
	return time.Hour * time.Duration(h)
}
