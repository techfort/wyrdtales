package models

import (
	"time"
)

const (
	PostsIndex   = "posts"
	StoryType    = "story"
	BlogpostType = "blogpost"
)

// Post models
type Post struct {
	AuthorID string    `json:"authorid"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Tags     []string  `json:"tags"`
	Posted   time.Time `json:"posted"`
	Status   string    `json:"status"`
	Category string    `json:"category"`
}
