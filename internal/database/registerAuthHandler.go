package database

import (
	"fmt"
	parse "historyback/internal/parse"
	"log"
	"net/http"
)

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {

	username, usersurname, login, password := parse.Parseuserinfo(r)

	validation(w, username, "username")
	validation(w, usersurname, "usersurname")
	validation(w, login, "login")
	passwordvalidatoin(w, password)

	result := AddUser(username, usersurname, login, password)

	if !result {
		err := tpl.ExecuteTemplate(w, "register.html", "Такое имя пользователя уже есть")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Fprint(w, "congrats, your account has been successfully created")

}
