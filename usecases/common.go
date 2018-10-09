package usecases

import (
	"github.com/olivere/elastic"
	"github.com/techfort/wyrdtales/models"
)

type usecase struct {
	Elastic *elastic.Client
}

// UseCase interface
type UseCase interface {
	SavePost(post models.Post)
}
