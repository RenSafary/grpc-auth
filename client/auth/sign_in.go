package auth

import (
	"log"
	"net/http"
	"text/template"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/sign_in.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func GetUserDataSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Content-Type", "text/html")

}
