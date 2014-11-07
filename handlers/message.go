package handlers

import (
	"github.com/PuerkitoBio/goquery"
	irc "github.com/fluffle/goirc/client"
	"strings"
)

// sender info is like this: <nick>!<user>@<ip>
// we want to extract <nick> only
func GetSenderNick(line *irc.Line) string {
	sender_info := line.Src
	index_bang := strings.Index(sender_info, "!")
	return sender_info[0:index_bang]
}

func ExtractLink(text string) string {
	start := strings.Index(text, "http")
	start_of_link := text[start:]
	link := strings.Split(start_of_link, " ")[0]
	return link
}

func GetTitle(link string) string {
	doc, err := goquery.NewDocument(link)
	if err != nil {
		return ""
	}

	title := ""
	doc.Find("html head").Each(func(i int, s *goquery.Selection) {
		title = s.Find("title").Text()
	})
	return title
}

// Properly handle newlines in string when sending back reply
func SendIRCSanitized(conn *irc.Conn, target string, msg string) {
	for _, s := range strings.Split(msg, "\n") {
		conn.Privmsg(target, s)
	}
}

func GetTwitterTitle(link string) string {
	title := GetTitle(link)
	return strings.Replace(title, " on Twitter", "", 1)
}

func PrintUrlTitle(conn *irc.Conn, msg string, target string) {
	if strings.Contains(msg, "http://") || strings.Contains(msg, "https://") {
		link := ExtractLink(msg)
		if strings.Contains(link, "https://twitter.com") {
			SendIRCSanitized(conn, target, GetTwitterTitle(link))
		} else {
			SendIRCSanitized(conn, target, GetTitle(link))
		}
	}
}

func HelpHandle() string {
	return "!ping\n" +
		"!decide <choice1> | <choice2> | ...\n" +
		"!urban <words> :: searches Urban Dictionnary for answers to life\n" +
		"!wolfram <words> :: Asks Wolfram for answers to life\n" +
		"!weather {place} :: Default location is Toronto. You can specify your own location\n" +
		"!pic <words> :: Returns a link to a picture\n" +
		"!video <words> :: Returns a link to a video"
}

func MessageHandle(conn *irc.Conn, line *irc.Line) {
	target := line.Target()
	msg := strings.Trim(line.Text(), " ")
	sender_nick := GetSenderNick(line)

	go PrintUrlTitle(conn, msg, target)

	switch {
	case msg == "!ping":
		conn.Privmsg(target, sender_nick+": pong!")

	case strings.HasPrefix(msg, "!decide"):
		answer := DecideHandler(msg)
		if answer != "" {
			conn.Privmsg(target, sender_nick+": "+answer)
		}

	case strings.HasPrefix(msg, "!urban"):
		SendIRCSanitized(conn, target, UrbanHandler(msg))

	case strings.HasPrefix(msg, "!wolfram"):
		SendIRCSanitized(conn, target, WolframHandler(msg))

	case strings.HasPrefix(msg, "!weather"):
		SendIRCSanitized(conn, target, WeatherHandler(msg))

	case strings.HasPrefix(msg, "!pic"):
		SendIRCSanitized(conn, target, PicHandler(msg))

	case strings.HasPrefix(msg, "!video"):
		SendIRCSanitized(conn, target, VideoHandler(msg))

	case strings.HasPrefix(msg, "!help"):
		SendIRCSanitized(conn, sender_nick, HelpHandle())
	}
}
