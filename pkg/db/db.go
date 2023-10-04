package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	conexao := "user=jessica dbname=usersDB password=disable host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}

	return db
}
