package jwt

import "time"

func HourToDuration(hour int) time.Duration {
	return time.Hour * time.Duration(hour)
}
