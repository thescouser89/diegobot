package handlers

import (
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

	case strings.HasPrefix(msg, "!wolfram"):
		conn.Privmsg(target, WolframHandler(msg))

	case strings.HasPrefix(msg, "!weather"):
		conn.Privmsg(target, WeatherHandler(msg))
	}
}
