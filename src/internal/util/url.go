package util

import "net/url"

func UrlEncode(value string) string {
	return url.QueryEscape(value)
}

func UrlDecode(value string) string {
	decodedValue, err := url.QueryUnescape(value)
	if err != nil {
		return ""
	}
	return decodedValue
}
