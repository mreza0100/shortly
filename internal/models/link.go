package models

import "time"

type Link struct {
	Short       string
	Destination string
	UserEmail   string
	CreatedAt   time.Time
}
