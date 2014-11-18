package handlers

import (
	"github.com/thescouser89/diegobot/memory"
	"strings"
)

func SaveHandler(msg string) {
	trimmed_text := RemoveCommandFromString(msg, "!save")
	key := strings.Split(trimmed_text, " ")[0]
	value := strings.Replace(trimmed_text, key, "", 1)
	value_trimmed := strings.Trim(value, " ")
	memory.PutKey(key, value_trimmed)
}

func RetrieveHandler(msg string) string {
	key := RemoveCommandFromString(msg, "!retrieve")
	return memory.GetKey(key)
}
