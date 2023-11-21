package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	conection := "user=postgres dbname=postgres password=login_system host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err)
	}

	return db
}
