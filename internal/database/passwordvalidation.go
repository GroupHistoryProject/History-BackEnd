package database

import (
	"log"
	"net/http"
	"unicode"
)

func passwordvalidatoin(w http.ResponseWriter, password string) {

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

}
