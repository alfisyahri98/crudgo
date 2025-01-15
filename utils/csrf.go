package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateCSRFToken() (string, error) {
	token := make([]byte, 32)

	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(token), nil
}
