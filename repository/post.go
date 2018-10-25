package repository

import (
	"encoding/json"

	"github.com/techfort/wyrdtales/utils"

	"github.com/olivere/elastic"
	"github.com/techfort/wyrdtales/models"
)

type postRepo struct {
	repo
}

// PostRepository interface
type PostRepository interface {
	ByID(ID string) (models.Post, error)
	SavePost(post models.Post) (*elastic.IndexResponse, error)
	SearchTerm(phrase string) (*elastic.SearchResult, error)
}

// ByID retrieves a post by id
func (r postRepo) ByID(ID string) (models.Post, error) {
	gr, err := r.Elastic.Get().Index(models.PostsIndex).Type(models.StoryType).Id(ID).Do(r.Context)
	var post models.Post
	utils.PanicIf(err)
	utils.PanicIf(json.Unmarshal(*gr.Source, &post))
	return post, err
}

// SavePost saves a post to elasticsearch
func (r postRepo) SavePost(post models.Post) (*elastic.IndexResponse, error) {
	post.SetDefaults()
	return r.Elastic.Index().Index(models.PostsIndex).Type(models.StoryType).BodyJson(post).Do(r.Context)
}

func (r postRepo) SearchTerm(phrase string) (*elastic.SearchResult, error) {
	q := elastic.NewMatchQuery("body", phrase).MinimumShouldMatch("75%")
	bq := elastic.NewBoolQuery().Should(q)
	return r.Elastic.Search().Query(bq).Do(r.Context)
}
