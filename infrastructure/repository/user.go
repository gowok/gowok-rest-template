package repository

import (
	"context"

	"github.com/gowok/gowok-rest-template/infrastructure/db"
	"github.com/gowok/gowok-rest-template/infrastructure/errs"
	"github.com/gowok/gowok-rest-template/infrastructure/repository/model"
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

func (r UserReader) GetUserByEmail(c context.Context, email string) (model.User, error) {
	rows, err := r.QueryContext(c, `
			SELECT * FROM users
			WHERE email = $1
		`,
		email,
	)

	if err != nil {
		return model.User{}, err
	}

	result := []model.User{}
	for rows.Next() {
		r := model.User{}
		err := rows.Scan(
			&r.ID,
			&r.Email,
			&r.Password,
		)
		if err != nil {
			return model.User{}, err
		}

		result = append(result, r)
	}

	if len(result) == 0 {
		return model.User{}, errs.ErrNotFound
	}

	return result[0], nil
}
