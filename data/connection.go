package data

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection interface {
}

type PostgresSQL struct {
	db *sqlx.DB
}

func New(connection string) (Connection, error) {
	db, err := sqlx.Connect("postgres", connection)
	if err != nil {
		return nil, err
	}
	return &PostgresSQL{db}, nil
}
