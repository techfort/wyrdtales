package models

import (
	"time"
)

// Post models
type Post struct {
	PostID   string
	Title    string
	Body     string
	Tags     []string
	Posted   time.Time
	Status   string
	Category string
}
