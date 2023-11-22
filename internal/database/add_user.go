package database

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func TryAddUser(name, surname, login, password string) bool {
	db, err := sql.Open("mysql", "root:FuFa2020@tcp(127.0.0.1:3306)/HistoryProject")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if !userExists(db, login) {
		return false
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	_, err = db.Exec(
		`INSERT INTO userdata (name, surname, Login, Password) VALUES ($1, $2, $3, $3)`,
		name,
		surname,
		login,
		hash,
	)
	if err != nil {
		return false
	}

	return true
}

func userExists(db *sql.DB, login string) bool {
	stmt := "SELECT UserID FROM UserInfo WHERE login = ?"

	row := db.QueryRow(stmt, login)

	var uID string
	err := row.Scan(&uID)
	if !errors.Is(err, sql.ErrNoRows) {
		return false
	}

	return true
}
