package parse

import (
	"log"
	"net/http"
)

func Parseuserinfo(r *http.Request) (string, string, string, string) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	username := r.FormValue("username")
	usersurname := r.FormValue("usersurname")
	login := r.FormValue("login")
	password := r.FormValue("password")

	return username, usersurname, login, password
}
