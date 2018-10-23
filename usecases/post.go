package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"

	"github.com/techfort/wyrdtales/models"
	"github.com/techfort/wyrdtales/repository"
)

type postUsecase struct {
	Repo repository.PostRepository
}

// NewPostUseCase returns a PostUseCase interface
func NewPostUseCase(repo repository.PostRepository) PostUseCase {
	return postUsecase{repo}
}

// PostUseCase interface
type PostUseCase interface {
	// GetPost returns a post by ud
	GetPost(ID string) (models.Post, error)
	SavePost(post models.Post) (*elastic.IndexResponse, error)
	/*
		SavePost(post models.Post) (models.Post, error)
		Unpublish(ID string) (models.Post, error)
	*/
}

// GetPost returns a post
func (puc postUsecase) GetPost(ID string) (models.Post, error) {
	gr, err := puc.Repo.ByID(ID)
	var post models.Post
	if err != nil || !gr.Found {
		return post, err
	}
	fmt.Println(fmt.Sprintf("Post: %+v", gr.Fields))
	err = json.Unmarshal(*gr.Source, &post)
	return post, err
}

func (puc postUsecase) SavePost(post models.Post) (*elastic.IndexResponse, error) {
	ir, err := puc.Repo.SavePost(post)
	fmt.Println(fmt.Sprintf("IR: %+v", ir))
	return ir, err
}
