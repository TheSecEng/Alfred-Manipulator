package util

import (
	b64 "encoding/base64"
)

func Base64Encode(value string) string {
	return b64.StdEncoding.EncodeToString([]byte(value))
}

func Base64Decode(value string) string {
	sDec, err := b64.StdEncoding.DecodeString(value)
	if err != nil {
		return ""
	}
	return string(sDec)
}
