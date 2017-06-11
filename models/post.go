package models

import "time"

// Post model
type Post struct {
	Title     string    `json:"name" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
}

// Posts model
type Posts []Post
