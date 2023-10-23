package database

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	_ "github.com/lib/pq"
	"log"
)

func AddUser(name, surname, login, password string) {
	loginHash := md5.Sum([]byte(login))
	passwordHash := md5.Sum([]byte(password))
	connStr := "user = postgres password = '123456' dbname = historyback sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO userdata (name, surname, hashlogin) VALUES ($1, $2, $3)`, name, surname, hex.EncodeToString(loginHash[:]))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO logindata(hashlogin, hashpassword) VALUES ($1, $2)`, hex.EncodeToString(loginHash[:]), hex.EncodeToString(passwordHash[:]))
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUser(login string) {
	loginHash := md5.Sum([]byte(login))
	connStr := "user = postgres password = '123456' dbname = historyback sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM userdata WHERE hashlogin = $1`, hex.EncodeToString(loginHash[:]))
	if err != nil {
		log.Fatal(err)
	}
}
