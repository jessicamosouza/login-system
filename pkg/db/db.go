package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	conection := "user=null dbname=postgres password=disable host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err)
	}

	return db
}
