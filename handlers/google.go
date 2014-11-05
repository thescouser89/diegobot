package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	IMAGE_QUERY = "http://ajax.googleapis.com/ajax/services/search/images?v=1.0&q="
	VIDEO_QUERY = "http://ajax.googleapis.com/ajax/services/search/video?v=1.0&q="
)

type SearchResult struct {
	ResponseStatus int          `json:"responseStatus"`
	Data           Responsedata `json:"responseData"`
}

type VideoSearchResult struct {
	Data ResponseDataVideo `json:"responseData"`
}

type Responsedata struct {
	Results []Result `json:"results"`
}

type ResponseDataVideo struct {
	Results []VideoResult `json:"results"`
}

type Result struct {
	UnescapedUrl string `json:"unescapedUrl"`
}

type VideoResult struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

func PicHandler(msg string) string {
	var meme string
	text := strings.Replace(msg, "!pic", "", 1)
	trimmed_text := strings.Trim(text, " ")

	if trimmed_text == "" {
		return ""
	} else {
		meme = url.QueryEscape(trimmed_text)
		resp, err := http.Get(IMAGE_QUERY + meme)

		if err != nil {
			log.Fatal(err)
			return "Booboo :("
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		reply := new(SearchResult)
		decoder.Decode(reply)
		return string(reply.Data.Results[0].UnescapedUrl)
	}
}

func VideoHandler(msg string) string {
	var search string
	text := strings.Replace(msg, "!video", "", 1)
	trimmed_text := strings.Trim(text, " ")

	if trimmed_text == "" {
		return ""
	} else {
		search = url.QueryEscape(trimmed_text)
		resp, err := http.Get(VIDEO_QUERY + search)

		if err != nil {
			log.Fatal(err)
			return "Booboo :("
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		reply := new(VideoSearchResult)
		decoder.Decode(reply)

		title := reply.Data.Results[0].Title
		unescaped, err := url.QueryUnescape(reply.Data.Results[0].Url)

		if err == nil {
			return string(unescaped + "\n" + title)
		}
		return ""
	}
}
