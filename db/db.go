package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB //lower-case is private
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:123@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
