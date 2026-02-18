package models

import "time"

type Job struct {
	ID          int64
	Title       string
	URL         string
	Description string
	DateAdded   time.Time
}
