package repository

import (
	"context"

	"github.com/olivere/elastic"
)

type repo struct {
	Elastic *elastic.Client
	Context context.Context
}

// RepoFactory interface
type RepoFactory interface {
	GetPostRepository() PostRepository
}

// NewRepoFactory returns a repo struct but with the interface of the desired repo
func NewRepoFactory(es *elastic.Client) RepoFactory {
	return repo{es, context.Background()}
}

// GetPostRespository returns a PostRepository
func (r repo) GetPostRepository() PostRepository {
	return postRepo{r}
}

func (r repo) GetUserRepository() UserRepository {
	return userRepo{r}
}
