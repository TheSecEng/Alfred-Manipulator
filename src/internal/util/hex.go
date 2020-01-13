package util

import (
	"encoding/hex"
)

func HexEncode(value string) string {
	return hex.EncodeToString([]byte(value))
}

func HexDecode(value string) string {
	sDec, err := hex.DecodeString(value)
	if err != nil {
		return ""
	}
	return string(sDec)
}
