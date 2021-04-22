package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"magellan/pkg/config"
)

// DB ...
type DB struct {
	Client *sql.DB
}

func formatConnString(host string, port string, user string, password string, dbname string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

// Connects to DB
func Connect(config config.DBConfig) (*sql.DB, error) {
	connStr := formatConnString(config.DBHost, config.DBPort, config.DBUser, config.DBPwd, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Close ...
func (d *DB) Close() error {
	return d.Client.Close()
}
