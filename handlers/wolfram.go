package handlers

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	WOLFRAM_APP_ID   = "PAHYHY-V23QR5658J"
	WOLFRAM_LINK     = "http://api.wolframalpha.com/v2/query?"
	WOLFRAM_ENDPOINT = WOLFRAM_LINK + "appid=" + WOLFRAM_APP_ID + "&input="
)

type WolframAns struct {
	Pods []Pod `xml:"pod"`
}

type Pod struct {
	PlainText string `xml:"subpod>plaintext"`
}

func WolframHandler(msg string) string {
	text := strings.Replace(msg, "!wolfram", "", 1)
	trimmed_text := strings.Trim(text, " ")
	search := url.QueryEscape(trimmed_text)

	resp, err := http.Get(WOLFRAM_ENDPOINT + search)

	if err != nil {
		log.Fatal(err)
		return "Booboo :("
	}
	defer resp.Body.Close()

	// Decode XML response
	decoder := xml.NewDecoder(resp.Body)
	reply := new(WolframAns)
	decoder.Decode(reply)

	if len(reply.Pods) > 1 {
		return reply.Pods[1].PlainText
	}
	return "Nothing found :("
}
