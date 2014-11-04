package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	SEARCH_QUERY = "http://ajax.googleapis.com/ajax/services/search/images?v=1.0&q="
)

type SearchResult struct {
	responseData ResponseData `json:"responseData"`
}
type ResponseData struct {
	results []Results `json:"results"`
}

type Results struct {
	UnescapedUrl string `json:"unescapedUrl"`
}

func MemeHandler(msg string) string {
	var meme string
	text := strings.Replace(msg, "!meme", "", 1)
	trimmed_text := strings.Trim(text, " ")

	if trimmed_text == "" {
		return ""
	} else {
		meme = url.QueryEscape(trimmed_text)
		resp, err := http.Get(SEARCH_QUERY + meme)

		if err != nil {
			log.Fatal(err)
			return "Booboo :("
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		reply := new(SearchResult)
		decoder.Decode(reply)
		return reply.responseData.results[0].UnescapedUrl
	}
}
