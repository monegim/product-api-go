package data

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection interface {
	IsConnected() (bool, error)
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

func (c *PostgresSQL) IsConnected() (bool, error) {
	err := c.db.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}
