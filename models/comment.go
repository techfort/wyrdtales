package models

import (
	"time"
)

// CommentResult struct
type CommentResult struct {
	Result Comment
	Error  error
}

// CommentsResult struct
type CommentsResult struct {
	Result []Comment
	Error  error
}

// Comment model
type Comment struct {
	ID       string
	PostID   string
	AuthorID string
	Body     string
	Posted   time.Time
}
