package auth

import (
	"context"
	"errors"

	"github.com/gowok/gowok-rest-template/infrastructure/errs"
	"github.com/gowok/gowok-rest-template/infrastructure/repository"
	"github.com/gowok/ioc"
)

type Service struct {
	userReader *repository.UserReader
}

func NewService() *Service {
	return &Service{
		userReader: ioc.Get(repository.UserReader{}),
	}
}

func (s Service) Login(c context.Context, email string, password string) error {
	user, err := s.userReader.GetUserByEmail(c, email)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return ErrInvalidEmailOrPass
		}

		return err
	}

	if user.Password == "" {
		return ErrInvalidEmailOrPass
	}

	if user.Password != password {
		return ErrInvalidEmailOrPass
	}

	return nil
}
