package models

import "time"

// User type
type User struct {
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Users type
type Users []User
