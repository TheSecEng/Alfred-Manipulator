package util

import (
	"crypto/md5"
	"fmt"
)

func ToMD5(value string) string {
	rValue := md5.Sum([]byte(value))
	return fmt.Sprintf("%x", string(rValue[:]))
}
