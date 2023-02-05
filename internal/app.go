package internal

import (
	"github.com/gowok/gowok-rest-template/internal/service"
	"github.com/gowok/ioc"
)

type App struct {
	Service *service.Service
}

func NewApp() *App {
	return &App{
		Service: service.NewService(),
	}
}

func PrepareAll() {
	app := NewApp()
	ioc.Set(func() App {
		return *app
	})
}
