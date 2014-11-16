package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

const (
	IMAGE_QUERY    = "http://ajax.googleapis.com/ajax/services/search/images?v=1.0&q="
	VIDEO_QUERY    = "http://ajax.googleapis.com/ajax/services/search/video?v=1.0&q="
	SEARCH_QUERY_1 = "https://www.google.com/search?q="
	SEARCH_QUERY_2 = "&btnI="
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
	trimmed_text := RemoveCommandFromString(msg, "!pic")

	var meme string
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
	trimmed_text := RemoveCommandFromString(msg, "!video")

	var search string
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

		unescaped, err := url.QueryUnescape(reply.Data.Results[0].Url)

		if err != nil {
			return ""
		}

		title := reply.Data.Results[0].Title
		unescaped_title, err_title := url.QueryUnescape(title)

		if err_title == nil {
			return string(unescaped + "\n" + unescaped_title)
		}

		return ""
	}
}

func SearchHandler(msg string) string {
	var search string
	trimmed_text := RemoveCommandFromString(msg, "!search")
	search = url.QueryEscape(trimmed_text)

	search_url := SEARCH_QUERY_1 + search + SEARCH_QUERY_2
	resp, err := http.Get(search_url)

	if err != nil {
		return ""
	}
	finalURL := resp.Request.URL.String()
	return finalURL + "\n" + GetTitle(finalURL)
}
