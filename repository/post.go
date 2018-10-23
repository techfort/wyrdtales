package repository

import (
	"github.com/olivere/elastic"
	"github.com/techfort/wyrdtales/models"
)

// PostRepository interface
type PostRepository interface {
	ByID(ID string) (*elastic.GetResult, error)
	SavePost(post models.Post) (*elastic.IndexResponse, error)
	/*
		ByIDs(ID ...string) PostsResult
		Upload(post models.Post) PostResult
		Publish(post models.Post) PostResult
		SearchInBody(keywords ...string) PostsResult
		SearchByTag(tags ...string) PostsResult
		SearchByCategory(category ...string) PostsResult
	*/
}

// ByID retrieves a post by id
func (r repo) ByID(ID string) (*elastic.GetResult, error) {
	return r.Elastic.Get().Index(models.PostsIndex).Type(models.StoryType).Id(ID).Do(r.Context)
}

// SavePost saves a post to elasticsearch
func (r repo) SavePost(post models.Post) (*elastic.IndexResponse, error) {
	return r.Elastic.Index().Index(models.PostsIndex).Type(models.StoryType).BodyJson(post).Do(r.Context)
}
