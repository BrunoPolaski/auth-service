package database

import "database/sql"

type Database interface {
	Connection() (*sql.DB, error)
	Close() error
}
