package database

type Database interface {
	Connection() error
	Close() error
}
