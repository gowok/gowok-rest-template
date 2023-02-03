package auth

import (
	"errors"

	"github.com/gowok/ioc"
)

var ErrInvalidEmailOrPass = errors.New("invalid email or password")

func InitModule() {
	ioc.Set(func() Service { return *NewService() })
	InitHttp()
}
