package models

import (
	"net/mail"

	"github.com/jessicamosouza/login-system/db"
)

type User struct {
	FirstName string `json:"fname" db:"firstname"`
	Lastname  string `json:"lname" db:"lastname"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

func SearchAllUsers() []User {
	db := db.InitDB()
	defer db.Close()

	selectAllusers, err := db.Query("select * from users")
	if err != nil {
		panic(err.Error())
	}

	u := User{}
	users := []User{}

	for selectAllusers.Next() {
		var fname, lname, email, password string

		err = selectAllusers.Scan(&fname, &lname, &email, &password)
		if err != nil {
			panic(err.Error())
		}

		u.FirstName = fname
		u.Lastname = lname
		u.Email = email
		u.Password = password

		users = append(users, u)
	}

	return users
}

func InsertUser(firstName, lastName, password string, email *mail.Address) {
	db := db.InitDB()
	defer db.Close()

	addUserDB, err := db.Prepare("insert into users (firstname, lastname, email, password)  values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	addUserDB.Exec(firstName, lastName, email, password)
	// tratar erro

}
