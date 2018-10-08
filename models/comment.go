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
	ID       int64
	PostID   int64
	AuthorID int64
	Body     string
	Posted   time.Time
}
