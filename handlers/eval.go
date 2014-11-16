package handlers

// TODO: Improve code readability. Super hacky
import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

const (
	LANGUAGE = "ruby/mri-2.1"
	WEBSITE  = "https://eval.in/"
)

func EvalHandler(msg string) string {
	trimmed_text := RemoveCommandFromString(msg, "!eval")
	resp, err := http.PostForm(WEBSITE,
		url.Values{"utf8": {"λ"},
			"code":    {"puts " + trimmed_text},
			"execute": {"on"},
			"lang":    {LANGUAGE},
			"input":   {""}})
	if err != nil {
		return "err1"
	}
	doc, err1 := goquery.NewDocumentFromResponse(resp)

	if err1 != nil {
		return "err2"
	}
	eval := ""
	doc.Find("html body div.paste").Each(func(i int, s *goquery.Selection) {
		eval = s.Find("pre").Text()
	})
	return strings.Split(eval, "\n")[1]
}
