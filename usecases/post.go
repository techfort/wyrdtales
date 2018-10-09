package usecases

import (
	"time"

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
}

// GetPost returns a post
func (puc postUsecase) GetPost(ID string) (models.Post, error) {
	gr, err := puc.Repo.ByID(ID)
	var post models.Post
	if err != nil || !gr.Found {
		return post, err
	}
	post.PostID = gr.Id
	post.AuthorID = gr.Fields["authorId"].(string)
	post.Body = gr.Fields["body"].(string)
	post.Category = gr.Fields["category"].(string)
	post.Posted = gr.Fields["posted"].(time.Time)
	post.Status = gr.Fields["status"].(string)
	post.Tags = gr.Fields["tags"].([]string)
	post.Title = gr.Fields["title"].(string)
	return post, err
}
