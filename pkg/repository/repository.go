package repository

import (
	"github.com/gowok/gowok-rest-template/pkg/repository/token"
	"github.com/gowok/gowok-rest-template/pkg/repository/user"
)

// Repository struct
type Repository struct {
	User  user.UserRepository
	Token token.TokenRepository
}

// NewRepository func
func NewRepository() Repository {
	return Repository{
		User:  user.New(),
		Token: token.New(),
	}
}
