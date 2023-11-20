package database

import (
	"log"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
