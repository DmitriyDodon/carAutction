package db

import (
	"database/sql"
	"app/config"
	"os"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type DBConnection struct {
	connection *sql.DB
}

func NewConnection(config *config.Config) (*DBConnection, error) {
	connection, error := sql.Open(config.GetSqlDriver(), config.GetDBFilePath())

	if error != nil {
		return &DBConnection{}, error
	}

	return &DBConnection{
		connection: connection,
	}, nil
}

func (d DBConnection) Execute(query string, args ...any) (sql.Result, error) {
	if len(args) > 0 {
		return d.connection.Exec(query, args...)
	}
	return d.connection.Exec(query)
}

func (d DBConnection) Query(query string, args ...any) (*sql.Rows, error) {
	if len(args) > 0 {
		return d.connection.Query(query, args...)
	}
	return d.connection.Query(query)
}

func (d DBConnection) QueryRow(query string, args ...any) *sql.Row {
	if len(args) > 0 {
		return d.connection.QueryRow(query, args...)
	}
	return d.connection.QueryRow(query)
}

func (d DBConnection) Close() {
	d.connection.Close()
}

func (d DBConnection) RunQueryFromFile(filePath string) (sql.Result, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return d.connection.Exec(string(dat))
}
