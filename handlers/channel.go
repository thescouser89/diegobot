package handlers

import (
	irc "github.com/fluffle/goirc/client"
	"log"
)

func ChannelsToJoin(conn *irc.Conn, line *irc.Line) {
	log.Println("Connected to IRC Server")
	conn.Join("#mcgillece")
	conn.Join("#izverifier")
	log.Println("Joined channel")
}
