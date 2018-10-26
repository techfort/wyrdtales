package repository

import (
	"github.com/olivere/elastic"
	"github.com/techfort/wyrdtales/models"
)

type userRepo struct {
	repo
}

// UserRepository interface
type UserRepository interface {
	SaveUser(user models.User) (*elastic.IndexResponse, error)
	ByID(ID string) (models.User, error)
	Update(ID string, user models.User) (*elastic.IndexResponse, error)
}

// SaveUser saves a user
func (r userRepo) SaveUser(user models.User) (*elastic.IndexResponse, error) {
	panic("Not implemented")
}

// ByID returns a user by its id
func (r userRepo) ByID(ID string) (models.User, error) {
	panic("Not implemented")
}

// Update updates a user
func (r userRepo) Update(ID string, user models.User) (*elastic.IndexResponse, error) {
	panic("Not implemented")
}
