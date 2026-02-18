package models

import "time"

type Job struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	DateAdded   time.Time `json:"dateAdded"`
}
