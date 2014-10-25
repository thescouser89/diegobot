package handlers

import (
	urban "github.com/dpatrie/urbandictionary"
	"strings"
)

func UrbanHandler(msg string) string {
	text := strings.Replace(msg, "!urban", "", 1)
	trimmed_text := strings.Trim(text, " ")
	search, err := urban.Query(trimmed_text)

	if err == nil {
		return search.Results[0].Definition
	} else {
		return "Nothing found :("
	}
}
