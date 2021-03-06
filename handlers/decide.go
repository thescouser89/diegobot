package handlers

import (
	"math/rand"
	"strings"
)

func DecideHandler(msg string) string {
	// remove the !decide prefix
	text := RemoveCommandFromString(msg, "!decide")

	// change all '?' to '!'
	affirmative_text := strings.Replace(text, "?", "!", -1)
	options := strings.Split(affirmative_text, "|")

	// choose a random option
	choice := rand.Int() % len(options)
	return strings.Trim(options[choice], " ")
}
