package handlers

import (
	"github.com/PuerkitoBio/goquery"
	irc "github.com/fluffle/goirc/client"
	"strings"
)

// Helper function to remove the command from the text of string from IRC
func RemoveCommandFromString(text, command string) string {
	text_no_cmd := strings.Replace(text, command, "", 1)
	return strings.Trim(text_no_cmd, " ")
}

// sender info is like this: <nick>!<user>@<ip>
// we want to extract <nick> only
func GetSenderNick(line *irc.Line) string {
	sender_info := line.Src
	index_bang := strings.Index(sender_info, "!")
	return sender_info[0:index_bang]
}

// extract the http(s) link from the text and return the link only
func ExtractLink(text string) string {
	start := strings.Index(text, "http")
	start_of_link := text[start:]
	link := strings.Split(start_of_link, " ")[0]
	return link
}

// get html title of a link
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

// Helper method to return the title of html link of a twitter link
// Function will remove the "on Twitter" in the title
func GetTwitterTitle(link string) string {
	title := GetTitle(link)
	return strings.Replace(title, " on Twitter", "", 1)
}

// Properly handle newlines in string when sending back reply
func SendIRCSanitized(conn *irc.Conn, target string, msg string) {
	for _, s := range strings.Split(msg, "\n") {
		conn.Privmsg(target, s)
	}
}

// bot prints on IRC the title of any html link posted on IRC
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
		"!holidays :: Next 3 holidays!\n" +
		"!pic <words> :: Returns a link to a picture\n" +
		"!video <words> :: Returns a link to a video\n" +
		"!search <words> :: \"I'm feeling Lucky\""
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

	case strings.HasPrefix(msg, "!holidays"):
		SendIRCSanitized(conn, target, Holidays())

	case strings.HasPrefix(msg, "!search"):
		SendIRCSanitized(conn, target, SearchHandler(msg))
	case strings.HasPrefix(msg, "!eval"):
		SendIRCSanitized(conn, target, EvalHandler(msg))
	case strings.HasPrefix(msg, "!save"):
		SaveHandler(msg)
	case strings.HasPrefix(msg, "!retrieve"):
		SendIRCSanitized(conn, target, RetrieveHandler(msg))
	}
}
