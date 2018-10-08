package models

// Follow struct
type Follow struct {
	Followed  User
	Followers []*User
}
