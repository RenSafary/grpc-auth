package db

import (
	"AuthService/client/utils"
	encryption "AuthService/client/utils"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DbUsers struct {
	DB *sql.DB
}

func GetEnvVariablesDB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", user, password, name, host, port, sslmode)
	return connStr
}

func Conn() (*DbUsers, error) {
	connStr := GetEnvVariablesDB()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DbUsers{DB: db}, nil
}

func (d *DbUsers) AddUser(username, password string) (string, error) {
	// does user exist or not
	var exist bool
	err := d.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)",
		username,
	).Scan(&exist)
	if err != nil {
		return "", err
	}

	if exist {
		return "", fmt.Errorf("user already exists")
	}

	// hash password
	hashPass, err := encryption.EncryptPass(password)
	if err != nil {
		return "", err
	}

	// add record
	_, err = d.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashPass)
	if err != nil {
		return "", err
	}

	token, err := utils.CreateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (d *DbUsers) CheckUser(username, password string) (string, error) {
	var hash_pass string
	err := d.DB.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&hash_pass)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		return "", err
	}

	err = encryption.DecryptPass(hash_pass, password)
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token, err := utils.CreateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}
