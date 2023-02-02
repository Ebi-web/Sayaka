package handler

import (
	"net/http"

	"Sayaka/db"
	"github.com/pkg/errors"

	"Sayaka/domain/model"
	"Sayaka/services/line"
	"Sayaka/utils"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	db *sqlx.DB
}

type CreateUserRequest struct {
	AccessToken string `json:"access_token"`
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{db}
}

func (h *UserHandler) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var req CreateUserRequest
	if err := utils.ParseRequestBody(r, &req); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	pf, err := line.GetProfileByAccessToken(req.AccessToken)
	if err != nil {
		return http.StatusForbidden, nil, err
	}

	user := &model.User{
		LINEUserID:  pf.UserID,
		DisplayName: pf.DisplayName,
		PhotoURL:    pf.PictureURL,
	}

	var createdID int64

	if err = db.TXHandler(h.db, func(tx *sqlx.Tx) error {
		id, err := InsertUser(tx, user)
		if err != nil {
			return err
		}
		createdID = id
		if err = tx.Commit(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, createdID, nil
}

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
