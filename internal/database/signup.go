package database

import (
	"fmt"
	"historyback/internal/parse"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func startSignUp() {
	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", http.FileServer(http.Dir("stylesheet"))))

	var err error
	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Parsing Templates Error:")
		log.Fatal(err)
	}

	http.HandleFunc("/register", signUpHandler)
	http.HandleFunc("/register-auth", signUpAuthHandler)

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func signUpHandler(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func signUpAuthHandler(writer http.ResponseWriter, r *http.Request) {
	username, userSurname, login, password := parse.Parseuserinfo(r)

	validateUsername(writer, username)
	validateUserSurname(writer, userSurname)
	validateLoginName(writer, login)
	validatePassword(writer, password)

	result := TryAddUser(username, userSurname, login, password)

	if !result {
		err := tpl.ExecuteTemplate(writer, "register.html", "This username already exists")
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err := fmt.Fprint(writer, "Account has been successfully created")
	if err != nil {
		log.Fatal(err)
	}
}
