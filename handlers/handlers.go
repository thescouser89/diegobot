package handlers

import (
	irc "github.com/fluffle/goirc/client"
	"log"
)

func Connected(conn *irc.Conn, line *irc.Line) {
	log.Println("Connected to IRC Server")
	conn.Join("#mcgillece")
	log.Println("Joined channel")
}

func Message(conn *irc.Conn, line *irc.Line) {
	MessageHandle(conn, line)
}

func Handle(c *irc.Conn) {
	c.HandleFunc("connected", Connected)
	c.HandleFunc("privmsg", Message)
}
