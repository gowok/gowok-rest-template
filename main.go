package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gowok/gowok"
	"github.com/gowok/gowok-rest-template/infrastructure/db"
	"github.com/gowok/gowok-rest-template/infrastructure/repository"
	"github.com/gowok/gowok-rest-template/module/auth"
	"github.com/gowok/ioc"
)

// init dependencies
func init() {
	conf, err := gowok.Configure()
	if err != nil {
		panic(err)
	}

	ioc.Set(func() gowok.Config { return conf })

	d, err := db.New()
	if err != nil {
		panic(err)
	}

	ioc.Set(func() db.DB { return *d })

	ioc.Set(func() fiber.App { return *fiber.New() })
}

// init repositories
func init() {
	ioc.Set(func() repository.UserReader { return *repository.NewUserReader() })
}

// init modules
func init() {
	auth.InitModule()
}

func main() {
	go gowok.GracefulStop(func() {
		log.Default().Println("shutdown ...")
	})

	conf := ioc.Get(gowok.Config{})
	app := ioc.Get(fiber.App{})
	app.Listen(conf.App.Rest.Host)
}
