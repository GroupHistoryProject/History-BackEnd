package database

import (
	"log"
	"net/http"
	"unicode"
)

func validation(w http.ResponseWriter, userdata string, typeofdata string) {

	var nameAlphaNumeric = true
	for _, char := range userdata {
		if !unicode.IsLetter(char) {
			nameAlphaNumeric = false
		}
	}

	var nameLength bool
	if 5 <= len(userdata) && len(userdata) <= 100 {
		nameLength = true
	}

	if !nameLength || !nameAlphaNumeric {

		if typeofdata == "username" {
			err := tpl.ExecuteTemplate(w, "register.html", "Пожалуйста, проверьте критерии для имени пользователя")
			if err != nil {
				log.Fatal(err)
			}
		} else if typeofdata == "usersurname" {
			err := tpl.ExecuteTemplate(w, "register.html", "Пожалуйста, проверьте критерии для фамилии пользователя")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := tpl.ExecuteTemplate(w, "register.html", "Пожалуйста, проверьте критерии для логина")
			if err != nil {
				log.Fatal(err)
			}
		}

		return
	}
}
