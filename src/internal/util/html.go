package util

import (
	"html"
)

func HtmlEncode(value string) string {
	return html.EscapeString(value)
}

func HtmlDecode(value string) string {
	return html.UnescapeString(value)
}
