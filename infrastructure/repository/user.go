package repository

import (
	"context"
	"fmt"

	"github.com/gowok/gowok-rest-template/infrastructure/db"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
)

type UserReader struct {
	database.SQLQuerier
}

func NewUserReader() *UserReader {
	d := ioc.Get(db.DB{})
	return &UserReader{d}
}

func (r UserReader) GetUserByEmail(c context.Context, email string) error {
	row := r.QueryRowContext(c, `
			SELECT * FROM users
			WHERE email = $1
		`,
		email,
	)

	if err := row.Err(); err != nil {
		return err
	}

	result := struct {
		ID    int
		Email string
	}{}
	err := row.Scan(&result)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
