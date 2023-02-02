package repository

import (
	"Sayaka/domain/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func InsertUser(db *sqlx.Tx, user *model.User) (int64, error) {
	stmt, err := db.Preparex("insert into users (line_user_id, display_name, photo_url) values (?,?,?)")
	if err != nil {
		return 0, errors.Wrap(err, "failed to set prepared statement")
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	result, err := stmt.Exec(user.LINEUserID, user.DisplayName, user.PhotoURL)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute insert user")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get last_insert_id")
	}

	return id, nil
}

func FindUserByLINEID(db *sqlx.DB, id string) (*model.User, error) {
	var user model.User
	err := db.Get(&user, "select id, line_user_id,display_name,photo_url from users where line_user_id = ?", id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user")
	}
	return &user, nil
}
