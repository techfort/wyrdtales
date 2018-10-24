package models

import (
	"time"
)

const (
	// PostsIndex constant
	PostsIndex = "posts"
	// StoryType constant
	StoryType = "story"
	// BlogpostType constant
	BlogpostType = "blogpost"
	// PublishedStatus constant
	PublishedStatus = "published"
	// DraftStatus constant
	DraftStatus = "draft"
	// CtgShortStory category
	CtgShortStory = "short story"
	// CtgNovel cateogry
	CtgNovel = "novel"
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

// SetDefaults set post defaults
func (p *Post) SetDefaults() {
	if p.Category == "" {
		p.Category = CtgShortStory
	}
	if p.Status == "" {
		p.Status = DraftStatus
	}
	if p.Posted.Year() == 1 {
		p.Posted = time.Now()
	}
}
