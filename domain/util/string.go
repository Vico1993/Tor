package util

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToCamelCase(str string) string {
	words := strings.Split(str, " ")
	key := words[0]

	for _, word := range words[1:] {
		key += cases.Title(language.English, cases.NoLower).String(word)
	}

	return key
}