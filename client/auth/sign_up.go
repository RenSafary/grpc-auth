package auth

import (
	"AuthService/client/grpc"
	encryption "AuthService/client/utils"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("./templates/sign_up.html")
	if err != nil {
		log.Println(err)
	}

	tmpl.Execute(w, nil)
}

func GetUserDataSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	hash_pass, err := encryption.EncryptPass(password)
	if err != nil {
		log.Println(err)
		return
	}

	status, message := grpc.SignUpGRPC(email, username, hash_pass)

	fmt.Printf("GRPC said: %t, %s\n", status, message)

	http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)

}
