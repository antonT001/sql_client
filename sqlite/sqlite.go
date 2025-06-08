package sqlite

import (
	"github.com/antonT001/sql_client"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func NewClient(dataSourcePath string) (*sql_client.DataBaseImpl, error) {
	conn, err := sqlx.Connect("sqlite3", dataSourcePath)
	if err != nil {
		return nil, err
	}

	return &sql_client.DataBaseImpl{DB: conn}, nil
}
