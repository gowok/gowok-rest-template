package db

import (
	"database/sql"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
	_ "github.com/lib/pq"
)

type DB struct {
	*database.PostgreSQL
}

func New() (*DB, error) {
	conf := ioc.Get(gowok.Config{})

	db, err := sql.Open("postgres", conf.Databases[0].DSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{&database.PostgreSQL{DB: db}}, nil
}
