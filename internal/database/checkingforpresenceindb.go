package database

import (
	"database/sql"
)

func boolcheckingforpresenceindb(db *sql.DB, login string) bool {

	stmt := "SELECT UserID FROM UserInfo WHERE login = ?"
	row := db.QueryRow(stmt, login)
	var uID string
	err := row.Scan(&uID)
	if err != sql.ErrNoRows {
		return false
	}

	return true
}
