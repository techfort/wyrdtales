package repository

import (
	"github.com/olivere/elastic"
	"github.com/techfort/wyrdtales/models"
)

type postRepo struct {
	repo
}

// PostRepository interface
type PostRepository interface {
	ByID(ID string) (*elastic.GetResult, error)
	Latest(num int) (*elastic.SearchResult, error)
	SavePost(post models.Post) (*elastic.IndexResponse, error)
	UpdateStory(ID string, post models.Post) (*elastic.UpdateResponse, error)
	SearchTerm(phrase string) (*elastic.SearchResult, error)
}

// ByID retrieves a post by id
func (r postRepo) ByID(ID string) (*elastic.GetResult, error) {
	return r.Elastic.Get().Index(models.PostsIndex).Type(models.StoryType).Id(ID).Do(r.Context)
}

// Latest returns the latest posts
func (r postRepo) Latest(num int) (*elastic.SearchResult, error) {
	return r.Elastic.Search().Index(models.PostsIndex).Type(models.StoryType).From(0).Size(num).Sort("posted", false).Do(r.Context)
}

// SavePost saves a post to elasticsearch
func (r postRepo) SavePost(post models.Post) (*elastic.IndexResponse, error) {
	post.SetDefaults()
	return r.Elastic.Index().Index(models.PostsIndex).Type(models.StoryType).BodyJson(post).Do(r.Context)
}

func (r postRepo) SearchTerm(phrase string) (*elastic.SearchResult, error) {
	q := elastic.NewMatchQuery("body", phrase).MinimumShouldMatch("75%")
	bq := elastic.NewBoolQuery().Should(q)
	return r.Elastic.Search().Index(models.PostsIndex).Type(models.StoryType).Query(bq).Do(r.Context)
}

// UpdateStory updates a story document
func (r postRepo) UpdateStory(ID string, post models.Post) (*elastic.UpdateResponse, error) {
	return r.Elastic.Update().Index(models.PostsIndex).Type(models.StoryType).Doc(post).Do(r.Context)
}
