package repository

import (
	"Sayaka/domain/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func InsertUser(db *sqlx.Tx, user *model.User) (int64, error) {
	stmt, err := db.Preparex("insert into users (line_user_id, display_name, photo_url) values ($1,$2,$3) RETURNING id")
	if err != nil {
		return 0, errors.Wrap(err, "failed to set prepared statement")
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	err = stmt.QueryRow(user.LINEUserID, user.DisplayName, user.PhotoURL).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute insert user")
	}

	return id, nil
}

func FindUserByLINEID(db *sqlx.DB, id string) (*model.User, error) {
	var user model.User
	err := db.Get(&user, "select id, line_user_id,display_name,photo_url from users where line_user_id = $1", id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user")
	}
	return &user, nil
}
