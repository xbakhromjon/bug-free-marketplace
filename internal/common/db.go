package common

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx"
)

func ConnectToDb(host, port, database, user, password string) (*pgx.Conn, error) {

	p, err := stringToUint16(port)
	if err != nil {
		log.Fatalf("Could not convert port %v to uint16", port)
	}
	config := pgx.ConnConfig{
		Host:     host,
		Port:     p,
		Database: database,
		User:     user,
		Password: password,
	}

	conn, err := pgx.Connect(config)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
		return nil, err
	}

	fmt.Println("Connected to database")

	return conn, nil
}

func stringToUint16(s string) (uint16, error) {
	value, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}

	result := uint16(value)

	return result, nil
}
