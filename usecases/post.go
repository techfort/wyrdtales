package usecases

import (
	"encoding/json"
	"fmt"
	"time"

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
	SavePost(rawJSON string) (*elastic.IndexResponse, error)
	/*
		SavePost(post models.Post) (models.Post, error)
		Unpublish(ID string) (models.Post, error)
	*/
}

// GetPost returns a post
func (puc postUsecase) GetPost(ID string) (models.Post, error) {
	return puc.Repo.ByID(ID)
}

// SavePost saves a post
func (puc postUsecase) SavePost(rawJSON string) (*elastic.IndexResponse, error) {
	var (
		post models.Post
		ir   *elastic.IndexResponse
	)
	err := json.Unmarshal([]byte(rawJSON), &post)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %+v", err.Error()))
		return ir, err
	}
	post.Posted = time.Now()
	fmt.Println(fmt.Sprintf("POST: %+v", post))
	ir, err = puc.Repo.SavePost(post)
	fmt.Println(fmt.Sprintf("%+v", ir))
	return ir, err
}
