package huser

import (
	"crypto/sha512"
	"encoding/base64"
)

func GetPasswordHash(s string) []byte {
	var data = []byte(s)
	var sum = sha512.Sum512(data)
	return sum[:]
}

func GetPasswordHashString(s string) string {
	var hashData = GetPasswordHash(s)
	var hashText = base64.StdEncoding.EncodeToString(hashData)
	return hashText
}