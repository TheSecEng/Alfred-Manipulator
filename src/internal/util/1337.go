package util

import "strings"

var characters map[string]string

func init() {
	characters = map[string]string{
		"a": "4",
		"b": "8",
		"c": "c",
		"d": "d",
		"e": "3",
		"f": "f",
		"g": "9",
		"h": "h",
		"i": "1",
		"j": "j",
		"k": "k",
		"l": "1",
		"m": "m",
		"n": "n",
		"o": "0",
		"p": "p",
		"q": "q",
		"r": "2",
		"s": "5",
		"t": "7",
		"u": "u",
		"v": "v",
		"w": "w",
		"x": "x",
		"y": "y",
		"z": "2",
	}
}

func LeetEncode(value string) string {
	leetString := ""

	for _, char := range strings.ToLower(value) {
		if _, ok := characters[string(char)]; ok {
			substitutions := characters[string(char)]
			leetString += substitutions
		} else {
			leetString += string(char)
		}
	}

	return leetString
}
