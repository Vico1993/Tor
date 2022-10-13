package util

import "strings"

func ToCamelCase(str string) string {
	words := strings.Split(str, " ")
	key := strings.ToLower(words[0])

	for _, word := range words[1:] {
		key += strings.Title(word)
	}

	return key
}