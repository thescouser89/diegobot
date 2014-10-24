package handlers

import (
	urban "github.com/dpatrie/urbandictionary"
	irc "github.com/fluffle/goirc/client"
	"math/rand"
	"strings"
)

// sender info is like this: <nick>!<user>@<ip>
// we want to extract <nick> only
func GetSenderNick(line *irc.Line) string {
	sender_info := line.Src
	index_bang := strings.Index(sender_info, "!")
	return sender_info[0:index_bang]
}

func DecideHandler(msg string) string {
	// remove the !decide prefix
	text := strings.Replace(msg, "!decide", "", 1)

	// change all '?' to '!'
	affirmative_text := strings.Replace(text, "?", "!", -1)
	options := strings.Split(affirmative_text, "|")

	// choose a random option
	choice := rand.Int() % len(options)
	return strings.Trim(options[choice], " ")
}

func UrbanHandler(msg string) string {
	text := strings.Replace(msg, "!urban", "", 1)
	trimmed_text := strings.Trim(text, " ")
	search, err := urban.Query(trimmed_text)

	if err == nil {
		return search.Results[0].Definition
	} else {
		return "Nothing found :("
	}
}

func MessageHandle(conn *irc.Conn, line *irc.Line) {
	target := line.Target()
	msg := strings.Trim(line.Text(), " ")
	sender_nick := GetSenderNick(line)

	switch {
	case msg == "!ping":
		conn.Privmsg(target, sender_nick+": pong!")

	case strings.HasPrefix(msg, "!decide"):
		answer := DecideHandler(msg)
		if answer != "" {
			conn.Privmsg(target, sender_nick+": "+answer)
		}

	case strings.HasPrefix(msg, "!urban"):
		conn.Privmsg(target, UrbanHandler(msg))
	}
}
