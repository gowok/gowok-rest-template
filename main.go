package main

import (
	"context"
	"log"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok-rest-template/infrastructure/db"
	"github.com/gowok/gowok-rest-template/infrastructure/repository"
	"github.com/gowok/ioc"
)

func init() {
	conf, err := gowok.Configure()
	if err != nil {
		panic(err)
	}

	ioc.Set(func() gowok.Config { return conf })
}

func init() {
	d, err := db.New()
	if err != nil {
		panic(err)
	}

	ioc.Set(func() db.DB { return *d })
}

func init() {
	ioc.Set(func() repository.UserReader { return *repository.NewUserReader() })
}

func main() {
	userReader := ioc.Get(repository.UserReader{})
	log.Fatal(userReader.GetUserByEmail(context.Background(), "alexunder@gmail.com"))
}
