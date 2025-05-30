package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type postgresAdapter struct{}

func NewPostgresAdapter() *postgresAdapter {
	return &postgresAdapter{}
}

func (pa *postgresAdapter) Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	conn, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, name,
		),
	)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()

	return conn, err
}

func (pa *postgresAdapter) Close(conn *sql.DB) error {
	return conn.Close()
}
