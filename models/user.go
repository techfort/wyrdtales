package models

// User model
type User struct {
	UserID    string
	Username  string
	FirstName string
	LastName  string
	Age       int64
	Address   string
	Country   string
	AuthID    string
	PenNames  []*PenName
}
