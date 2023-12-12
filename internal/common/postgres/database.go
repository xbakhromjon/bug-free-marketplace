package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	*pgxpool.Pool
}

// New provides PostgresDB struct init
func New(host, port, name, user, password, sslmode string) (*PostgresDB, error) {
	db := PostgresDB{}
	if err := db.connect(host, port, name, user, password, sslmode); err != nil {
		return nil, err
	}

	return &db, nil
}

func (p *PostgresDB) configToStr(host, port, name, user, password, sslmode string) string {
	var conn []string

	if len(host) != 0 {
		conn = append(conn, "host="+host)
	}

	if len(port) != 0 {
		conn = append(conn, "port="+port)
	}

	if len(user) != 0 {
		conn = append(conn, "user="+user)
	}

	if len(password) != 0 {
		conn = append(conn, "password="+password)
	}

	if len(name) != 0 {
		conn = append(conn, "dbname="+name)
	}

	if len(sslmode) != 0 {
		conn = append(conn, "sslmode="+sslmode)
	}

	return strings.Join(conn, " ")
}

func (p *PostgresDB) connect(host, port, name, user, password, sslmode string) error {
	pool, err := pgxpool.New(context.Background(), p.configToStr(host, port, name, user, password, sslmode))
	if err != nil {
		return fmt.Errorf("unable to connect database config: %w", err)
	}

	p.Pool = pool
	return nil
}

func (p *PostgresDB) Close() {
	p.Pool.Close()
}

func (p *PostgresDB) Error(err error) error {
	return err
}

func (p *PostgresDB) ErrSQLBuild(err error, message string) error {
	return fmt.Errorf("error during sql build, %s: %w", message, err)
}
