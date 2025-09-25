package main

import (
	"AuthService/client/auth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// sign-in
	r.HandleFunc("/sign-in/", auth.SignIn).Methods("GET")
	r.HandleFunc("/sign-in/proccess", auth.GetUserDataSignIn).Methods("POST")

	// sign-up
	r.HandleFunc("/sign-up/", auth.SignUp).Methods("GET")
	r.HandleFunc("/sign-up/proccess", auth.GetUserDataSignUp).Methods("POST")

	log.Println("Server started on port :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}
