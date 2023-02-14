package handler

import (
	"net/http"

	"Sayaka/db"
	"Sayaka/domain/model"
	"Sayaka/domain/repository"
	"Sayaka/utils"
	"github.com/jmoiron/sqlx"
)

type FlashCardHandler struct {
	db *sqlx.DB
}

type CreateFlashCardRequest struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

func NewFlashCardHandler(db *sqlx.DB) *FlashCardHandler {
	return &FlashCardHandler{db}
}

func (h *FlashCardHandler) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var req CreateFlashCardRequest
	if err := utils.ParseRequestBody(r, &req); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	usr, err := utils.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusUnauthorized, nil, nil
	}

	c := &model.FlashCard{
		UserID: usr.ID,
		Front:  req.Front,
		Back:   req.Back,
	}

	var createdID int64
	if err = db.TXHandler(h.db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertFlashCard(tx, c)
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

func (h *FlashCardHandler) Index(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	usr, err := utils.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusUnauthorized, nil, err
	}
	fc, err := repository.GetFlashCardsByUserID(h.db, usr.ID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, fc, nil
}
