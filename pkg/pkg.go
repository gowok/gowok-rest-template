package pkg

import (
	"github.com/gowok/gowok"
	"github.com/gowok/gowok-rest-template/pkg/config"
	"github.com/gowok/gowok-rest-template/pkg/database"
	"github.com/gowok/gowok-rest-template/pkg/repository"
	"github.com/gowok/gowok-rest-template/pkg/util/runner"
	"github.com/gowok/gowok-rest-template/pkg/validator"
	"github.com/gowok/ioc"
)

func PrepareAll() {
	runner.PrepareRuntime()

	conf, err := config.Configure()
	if err != nil {
		panic(err)
	}
	ioc.Set(func() gowok.Config {
		return conf
	})

	database.Configure(conf.Databases)

	ioc.Set(func() gowok.Validator {
		return validator.Configure()
	})

	repo := repository.NewRepository()
	ioc.Set(func() repository.Repository {
		return repo
	})

}
