package models

// Rating model
type Rating struct {
	RatingID string
	PostID   string
	Value    int64
	AuthorID string
}
