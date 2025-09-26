package auth

import (
	"AuthService/client/grpc"
	encryption "AuthService/client/utils"
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

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	hash_pass, err := encryption.EncryptPass(password)
	if err != nil {
		log.Println(err)
		return
	}

	response, token := grpc.SignInGRPC(username, string(hash_pass))

	tmpl, err := template.ParseFiles("templates/response.html")
	if err != nil {
		log.Fatal(err)
	}

	data := &ResponseData{Response: response, Token: token}

	tmpl.Execute(w, data)
}
