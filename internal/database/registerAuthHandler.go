package database

import (
	"fmt"
	parse "historyback/internal/parse"
	"log"
	"net/http"
	"unicode"
)

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {

	username, usersurname, login, password := parse.Parseuserinfo(r)

	// validation of username
	var nameAlphaNumeric = true
	for _, char := range username {
		if !unicode.IsLetter(char) {
			nameAlphaNumeric = false
		}
	}

	var nameLength bool
	if 5 <= len(username) && len(username) <= 100 {
		nameLength = true
	}

	if !nameLength || !nameAlphaNumeric {
		err := tpl.ExecuteTemplate(w, "register.html", "Пожалуйста, проверьте критерии для имени пользователя")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	nameAlphaNumeric = true
	for _, char := range usersurname {
		if !unicode.IsLetter(char) {
			nameAlphaNumeric = false
		}
	}

	if 5 <= len(usersurname) && len(usersurname) <= 100 {
		nameLength = true
	}

	if !nameLength || !nameAlphaNumeric {
		err := tpl.ExecuteTemplate(w, "register.html", "Пожалуйста, проверьте критерии для фамилии пользователя")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// validation of login
	nameAlphaNumeric = true
	for _, char := range login {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			nameAlphaNumeric = false
		}
	}

	if 5 <= len(login) && len(login) <= 50 {
		nameLength = true
	}

	if !nameLength || !nameAlphaNumeric {
		err := tpl.ExecuteTemplate(w, "register.html", "Пожалуйста, проверьте критерии для логина")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// validation of password
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {

		case unicode.IsLower(char):
			pswdLowercase = true

		case unicode.IsUpper(char):
			pswdUppercase = true

		case unicode.IsNumber(char):
			pswdNumber = true

		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true

		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}

	if 11 < len(password) && len(password) < 60 {
		pswdLength = true
	}

	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces {
		err := tpl.ExecuteTemplate(w, "register.html", "Пожалуйста, проверьте критерии для пароля")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	result := AddUser(username, usersurname, login, password)

	// check if the person is already in the database
	if !result {
		err := tpl.ExecuteTemplate(w, "register.html", "Такое имя пользователя уже есть")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Fprint(w, "congrats, your account has been successfully created")

}
