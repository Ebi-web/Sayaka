package repository

import (
	"Sayaka/domain/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func InsertFlashCard(db *sqlx.Tx, c *model.FlashCard) (int64, error) {
	stmt, err := db.Preparex("insert into flash_cards (user_id, front, back) values (?,?,?)")
	if err != nil {
		return 0, errors.Wrap(err, "failed to set prepared statement")
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	result, err := stmt.Exec(c.UserID, c.Front, c.Back)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute insert flash card")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get last_insert_id")
	}

	return id, nil
}

func GetFlashCardsByUserID(db *sqlx.DB, id int64) ([]model.FlashCard, error) {
	var cs []model.FlashCard
	err := db.Select(&cs, "select id, user_id,front,back from flash_cards where user_id = ?", id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch flash cards by user id")
	}
	return cs, nil
}
