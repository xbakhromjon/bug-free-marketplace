package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "db"
	user     = "postgres"
	password = "postgres"
	dbname   = "marketplace_db"
)

func OpenDatabaseConnection() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	return db, err
}
