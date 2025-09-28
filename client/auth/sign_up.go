package auth

import (
	"AuthService/client/grpc"
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

	username := r.FormValue("username")
	password := r.FormValue("password")

	response, token := grpc.SignUpGRPC(username, []byte(password))

	tmpl, err := template.ParseFiles("templates/response.html")
	if err != nil {
		log.Fatal(err)
	}

	data := &ResponseData{Response: response, Token: token}
	tmpl.Execute(w, data)
}
