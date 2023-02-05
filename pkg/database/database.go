package database

import (
	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
)

func Configure(conf []config.Database) {
	for _, c := range conf {
		var db *database.PostgreSQL
		db, err := database.NewPostgresql(c)

		if err != nil {
			panic(err)
		}

		ioc.Set(func() database.PostgreSQL { return *db })
	}
}
