package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/techfort/wyrdtales/utils"

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
	GetLatest(num int) ([]models.Post, error)
	SearchTerm(term string) ([]models.Post, error)
	Unmarshal(bytes []byte) (models.Post, error)
	/*
		SavePost(post models.Post) (models.Post, error)
		Unpublish(ID string) (models.Post, error)
	*/
}

// GetPost returns a post
func (puc postUsecase) GetPost(ID string) (models.Post, error) {
	gr := utils.Bypass(puc.Repo.ByID(ID)).(*elastic.GetResult)
	return puc.Unmarshal(*gr.Source)
}

func (puc postUsecase) GetLatest(num int) ([]models.Post, error) {
	sr := utils.Bypass(puc.Repo.Latest(num)).(*elastic.SearchResult)
	return UnmarshalPostSlice(sr)

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
	// post.Posted = time.Now()
	fmt.Println(fmt.Sprintf("POST: %+v", post))
	ir, err = puc.Repo.SavePost(post)
	fmt.Println(fmt.Sprintf("%+v", ir))
	return ir, err
}

// SearchTerm searches for documents with the term
func (puc postUsecase) SearchTerm(term string) ([]models.Post, error) {
	sr, err := puc.Repo.SearchTerm(term)
	if err != nil {
		return []models.Post{}, err
	}
	return UnmarshalPostSlice(sr)
}

func (puc postUsecase) Unmarshal(bytes []byte) (models.Post, error) {
	var post models.Post
	err := json.Unmarshal(bytes, &post)
	return post, err
}
