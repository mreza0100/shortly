package models

import "time"

type Link struct {
	Short       string
	Destination string
	UserId      string
	CreatedAt   time.Time
}
