package repository

import (
	"github.com/olivere/elastic"
)

// constants literals
const (
	// Post const
	Post = "post"
	// Story const
	Story = "story"
	// Blogpost const
	Blogpost = "blogpost"
)

// PostRepository interface
type PostRepository interface {
	ByID(ID string) (*elastic.GetResult, error)
	/*
		ByIDs(ID ...string) PostsResult
		Upload(post models.Post) PostResult
		Publish(post models.Post) PostResult
		SearchInBody(keywords ...string) PostsResult
		SearchByTag(tags ...string) PostsResult
		SearchByCategory(category ...string) PostsResult
	*/
}

func (r repo) ByID(ID string) (*elastic.GetResult, error) {
	return r.Elastic.Get().Id(ID).Do(r.Context)
}
