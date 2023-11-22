package database

import (
	"log"
	"net/http"
	"unicode"
)

func validatePassword(writer http.ResponseWriter, password string) {
	var (
		lowercasePresent, uppercasePresent      bool
		numberPresent, specialCharactersPresent bool
		correctLength                           bool
		noSpaces                                = true
	)

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			lowercasePresent = true

		case unicode.IsUpper(char):
			uppercasePresent = true

		case unicode.IsNumber(char):
			numberPresent = true

		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			specialCharactersPresent = true

		case unicode.IsSpace(char):
			noSpaces = false
		}
	}

	const minLength = 9
	const maxLength = 60

	correctLength = minLength < len(password) && len(password) < maxLength

	if !lowercasePresent ||
		!uppercasePresent ||
		!numberPresent ||
		!specialCharactersPresent ||
		!correctLength ||
		!noSpaces {
		err := tpl.ExecuteTemplate(writer, "register.html", "Password is not valid")
		if err != nil {
			log.Fatal(err)
		}

		return
	}
}

func validateUsername(writer http.ResponseWriter, data string) {
	var lettersOnly, correctLength bool

	for _, char := range data {
		if !unicode.IsLetter(char) {
			lettersOnly = false
		}
	}

	const (
		minLength = 1
		maxLength = 50
	)

	correctLength = minLength <= len(data) && len(data) <= maxLength

	if !correctLength || !lettersOnly {
		err := tpl.ExecuteTemplate(
			writer,
			"register.html",
			"User name is not valid",
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func validateUserSurname(writer http.ResponseWriter, data string) {
	var lettersOnly, correctLength bool

	for _, char := range data {
		if !unicode.IsLetter(char) {
			lettersOnly = false
		}
	}

	const (
		minLength = 1
		maxLength = 50
	)

	correctLength = minLength <= len(data) && len(data) <= maxLength

	if !correctLength || !lettersOnly {
		err := tpl.ExecuteTemplate(
			writer,
			"register.html",
			"User surname is not valid",
		)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func validateLoginName(writer http.ResponseWriter, data string) {
	var correctLength bool

	const (
		minLength = 1
		maxLength = 50
	)

	correctLength = minLength <= len(data) && len(data) <= maxLength

	if !correctLength {
		err := tpl.ExecuteTemplate(
			writer,
			"register.html",
			"Login name is not valid",
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
