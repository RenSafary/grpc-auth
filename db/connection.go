package db

import (
	encryption "AuthService/client/utils"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DbUsers struct {
	DB *sql.DB
}

func Conn() (*DbUsers, error) {
	connStr := "user=ilya password=123 dbname=AuthGRPC host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DbUsers{DB: db}, nil
}

func (d *DbUsers) AddUser(email, username, password string) error {

	// does user exist or not
	var exist bool
	err := d.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE username=$1 OR email=$2)",
		username, email,
	).Scan(&exist)
	if err != nil {
		return err
	}

	if exist == false {
		return err
	}

	// add record
	_, err = d.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *DbUsers) CheckUser(username, password string) (bool, error) {
	var hash_pass string
	err := d.DB.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&hash_pass)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	err = encryption.DecryptPass(hash_pass, password)
	if err != nil {
		return false, nil
	}

	return true, nil
}
