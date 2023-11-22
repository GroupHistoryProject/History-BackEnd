package database

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(name, surname, login, password string) bool {

	db, err := sql.Open("mysql", "root:FuFa2020@tcp(127.0.0.1:3306)/HistoryProject")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if !boolcheckingforpresenceindb(db, login) {
		return false
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	_, err = db.Exec(`INSERT INTO userdata (name, surname, Login, Password) VALUES ($1, $2, $3, $3)`, name, surname, login, hash)
	if err != nil {
		return false
	}

	return true
}
