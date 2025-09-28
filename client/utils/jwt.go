package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Header struct {
	Alg  string `json:"alg"`
	Type string `json:"typ`
}

type Payload struct {
	Sub string `json:"username"`
	Exp int    `json:"exp"`
}

func base64UrlEncode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func base64UrlDecode(data string) (string, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func GetSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	secret_key := os.Getenv("SECRET_KEY")
	return secret_key
}

func CreateToken(username string) (string, error) {
	secret_key := GetSecretKey()

	h := Header{Alg: "sha256", Type: "jwt"}
	p := Payload{Sub: username, Exp: 6}

	header_json, err := json.Marshal(h)
	if err != nil {
		return "", err
	}

	payload_json, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	encoded_header := base64UrlEncode(header_json)
	encoded_payload := base64UrlEncode(payload_json)

	message := encoded_header + "." + encoded_payload

	hm := hmac.New(sha256.New, []byte(secret_key))
	hm.Write([]byte(message))
	signature := hm.Sum(nil)

	signatureEncoded := base64UrlEncode(signature)

	token := message + "." + signatureEncoded

	return token, nil
}

func GetPayload(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token format")
	}

	payload, err := base64UrlDecode(parts[1])
	if err != nil {
		return "", err
	}

	return payload, nil
}
