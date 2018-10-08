package repository

import "github.com/techfort/wyrdtales/models"

// PostResult struct
type PostResult struct {
	Result models.Post
	Error  error
}

// PostsResult struct
type PostsResult struct {
	Result []models.Post
	Error  error
}

// PostRepository interface
type PostRepository interface {
	ByID(ID string) PostResult
	ByIDs(ID ...string) PostsResult
}
