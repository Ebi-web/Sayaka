package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	datasource string
}

func NewMySQL(datasource string) *MySQL {
	return &MySQL{
		datasource: datasource,
	}
}

func (db *MySQL) Open() (*sqlx.DB, error) {
	return sqlx.Open("mysql", db.datasource)
}
