package handlers

import (
	urban "github.com/dpatrie/urbandictionary"
)

func UrbanHandler(msg string) string {
	trimmed_text := RemoveCommandFromString(msg, "!urban")
	search, err := urban.Query(trimmed_text)

	if err == nil {
		return search.Results[0].Definition
	} else {
		return "Nothing found :("
	}
}
