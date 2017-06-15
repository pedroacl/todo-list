package models

// Comment model
type Comment struct {
	PostID  string `json:"postId"`
	Content string `json:"content"`
}
