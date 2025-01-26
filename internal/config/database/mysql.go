package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
)

type MySQLAdapter struct{}

func (m MySQLAdapter) Connection() (*sql.DB, error) {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=auth-service sslmode=disable", dbHost, dbUser, dbPassword))
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
