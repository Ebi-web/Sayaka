package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	datasource string
}

func NewPostgreSQL(datasource string) *PostgreSQL {
	return &PostgreSQL{
		datasource: datasource,
	}
}

func (db *PostgreSQL) Open() (*sqlx.DB, error) {
	return sqlx.Open("postgres", db.datasource)
}
