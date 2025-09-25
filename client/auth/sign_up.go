package auth

import (
	"AuthService/client/grpc"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
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

	hash_pass, err := EncryptPass(password)
	if err != nil {
		log.Println(err)
		return
	}

	status, message := grpc.SignUpGRPC(email, username, hash_pass)

	fmt.Printf("GRPC said: %t, %s\n", status, message)

	http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)

}

func EncryptPass(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}
