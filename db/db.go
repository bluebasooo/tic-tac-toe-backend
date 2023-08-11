package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB //lower-case is private
}

func NewDatabase() (*Database, error) {
	login := os.Getenv("QOVERY_POSTGRESQL_")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", login, password, host, port, dbName))
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
