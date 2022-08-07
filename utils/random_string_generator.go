package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomStringGenerator(strlen int) (string, error){
	b := make([]byte, strlen)
	_, err := rand.Read(b)
	if err != nil {
        return "", err
    }
	return base64.URLEncoding.EncodeToString(b), err
}