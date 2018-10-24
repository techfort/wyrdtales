package repository

import (
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/techfort/wyrdtales/models"
)

// PostRepository interface
type PostRepository interface {
	ByID(ID string) (models.Post, error)
	SavePost(post models.Post) (*elastic.IndexResponse, error)
	SearchTerm(term string) (*elastic.SearchResult, error)
}

// ByID retrieves a post by id
func (r repo) ByID(ID string) (models.Post, error) {
	gr, err := r.Elastic.Get().Index(models.PostsIndex).Type(models.StoryType).Id(ID).Do(r.Context)
	var post models.Post
	if err != nil || !gr.Found {
		return post, err
	}
	err = json.Unmarshal(*gr.Source, &post)
	return post, err
}

// SavePost saves a post to elasticsearch
func (r repo) SavePost(post models.Post) (*elastic.IndexResponse, error) {
	post.SetDefaults()
	return r.Elastic.Index().Index(models.PostsIndex).Type(models.StoryType).BodyJson(post).Do(r.Context)
}

func (r repo) SearchTerm(term string) (*elastic.SearchResult, error) {
	q := elastic.NewTermQuery("body", term)
	return r.Elastic.Search().Query(q).Do(r.Context)
}
