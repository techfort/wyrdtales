package models

import (
	"time"
)

// Post models
type Post struct {
	ID     int64
	Title  string
	Body   string
	Tags   []string
	Posted time.Time
}
