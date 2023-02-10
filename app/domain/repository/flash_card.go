package repository

import (
	"Sayaka/domain/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func InsertFlashCard(db *sqlx.Tx, c *model.FlashCard) (int64, error) {
	stmt, err := db.Preparex("insert into flash_cards (user_id, front, back) values ($1,$2,$3) RETURNING id")
	if err != nil {
		return 0, errors.Wrap(err, "failed to set prepared statement")
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	err = stmt.QueryRow(c.UserID, c.Front, c.Back).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute insert flash card")
	}

	return id, nil
}

func GetFlashCardsByUserID(db *sqlx.DB, id int64) ([]model.FlashCard, error) {
	var cs []model.FlashCard
	err := db.Select(&cs, "select id, user_id,front,back from flash_cards where user_id = $1", id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch flash cards by user id")
	}
	return cs, nil
}
