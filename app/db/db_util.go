package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func TXHandler(db *sqlx.DB, f func(*sqlx.Tx) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("start transaction failed: %w", err)
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				log.Fatalf("rollback failed: %v, err: %v", rollBackErr, err)
			}
			log.Print("Rollback operation")
			if p != nil {
				err = errors.New(fmt.Sprintf("recovered: %v", p))
			} else {
				err = fmt.Errorf("transaction: operation failed: %w", err)
			}
		}
	}()
	err = f(tx)
	if err != nil {
		return err
	}
	return nil
}
