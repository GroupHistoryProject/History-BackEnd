package database

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"unicode"
)

var tpl *template.Template
var db *sql.DB


func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****registerHandler running*****")
	tpl.ExecuteTemplate(w, "main.html", nil)
}

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.FormValue("username")

	// проверка на то, что в имени только буквы
	var nameAlphaNumeric = true
	for _, char := range username {
		if unicode.IsLetter(char) == false {
			nameAlphaNumeric = false
		}
	}

	// проверка длины имени пользлвателя
	var nameLength bool
	if 5 <= len(username) && len(username) <= 100 {
		nameLength = true
	}

	// возвращаем на регистрацию, если не так
	if !nameLength || !nameAlphaNumeric {
		tpl.ExecuteTemplate(w, "register.html", "Пожалуйства проверьте критерии для имени пользователя")
		return
	}

	usersurname := r.FormValue("usersurname")

	// проверка на то, что в имени только буквы и цифры
	nameAlphaNumeric = true
	for _, char := range usersurname {
		if unicode.IsLetter(char) == false {
			nameAlphaNumeric = false
		}
	}

	// проверка длины имени пользлвателя
	if 5 <= len(usersurname) && len(usersurname) <= 100 {
		nameLength = true
	}

	// возвращаем на регистрацию, если не так
	if !nameLength || !nameAlphaNumeric {
		tpl.ExecuteTemplate(w, "register.html", "Пожалуйства проверьте критерии для фамилии пользователя")
		return
	}

	// проверка login
	login := r.FormValue("login")

	// проверка на то, что в имени только буквы и цифры
	nameAlphaNumeric = true
	for _, char := range login {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}

	// проверка длины имени пользлвателя
	if 5 <= len(login) && len(login) <= 50 {
		nameLength = true
	}

	// возвращаем на регистрацию, если не так
	if !nameLength || !nameAlphaNumeric {
		tpl.ExecuteTemplate(w, "register.html", "Пожалуйства проверьте критерии для логина")
		return
	}

	// проверка валидности пароля
	password := r.FormValue("password")
	fmt.Println("password:", password, "\npswdLength:", len(password))

	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {

		// есть прописные буквы
		case unicode.IsLower(char):
			pswdLowercase = true

		// есть заглавные буквы
		case unicode.IsUpper(char):
			pswdUppercase = true

		// есть цифра
		case unicode.IsNumber(char):
			pswdNumber = true

		// есть символы
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true

		// нет пробелов
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}

	// длина пароля
	if 11 < len(password) && len(password) < 60 {
		pswdLength = true
	}

	// если не выполнены условия возвращаемся на регистрацию
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces {
		tpl.ExecuteTemplate(w, "register.html", "Пожалуйства проверьте критерии для пароля")
		return
	}

	result := AddUser(username, usersurname, login, password)
	// если пользователь уже есть, возвращаем на регистрацию
	if !result {
		tpl.ExecuteTemplate(w, "register.html", "Такое имя пользователя уже есть")
	}
	fmt.Fprint(w, "congrats, your account has been successfully created")

}

func AddUser(name, surname, login, password string) bool {
	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", "root:FuFa2020@tcp(127.0.0.1:3306)/HistoryProject")
	if err != nil {
		fmt.Println("Openning DB Error:")
		panic(err.Error())
	}

	// провреяем есть ли пользователь
	stmt := "SELECT UserID FROM UserInfo WHERE login = ?"
	row := db.QueryRow(stmt, login)
	var uID string
	err = row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
		return false
	}

	var hash []byte

	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		return false
	}

	// вставка
	var insertStmt *sql.Stmt
	insertStmt, err = db.Prepare("INSERT INTO UserInfo (Username, Usersurname, Login, Password) VALUES (?, ?, ?, ?);")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return false
	}
	defer insertStmt.Close()
	var result sql.Result

	result, err = insertStmt.Exec(name, surname, login, hash)
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("lastIns:", lastIns)
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("error inserting new user")
		return false
	}
	return true
}
