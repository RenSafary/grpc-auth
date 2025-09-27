package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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

func CreateToken(username string) (string, error) {
	secret := "2281337"

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

	hm := hmac.New(sha256.New, []byte(secret))
	hm.Write([]byte(message))
	signature := hm.Sum(nil)

	signatureEncoded := base64UrlEncode(signature)

	token := message + "." + signatureEncoded

	return token, nil
}

func GetPayload(token string) {

}
