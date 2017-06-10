package models

import "time"

// Post model
type Post struct {
	Title     string
	Contente  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Posts model
type Posts []Post
