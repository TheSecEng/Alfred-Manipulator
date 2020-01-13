package util

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func ToSha1(value string) string {
	rValue := sha1.Sum([]byte(value))
	return fmt.Sprintf("%x", string(rValue[:]))
}

func ToSha256(value string) string {
	rValue := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", string(rValue[:]))
}

func ToSha512(value string) string {
	rValue := sha512.Sum512([]byte(value))
	return fmt.Sprintf("%x", string(rValue[:]))
}
