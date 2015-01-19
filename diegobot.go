package main

import (
	irc "github.com/fluffle/goirc/client"
	"github.com/thescouser89/diegobot/handlers"
	"log"
	"os"
)

const (
	IRC_SERVER = "irc.freenode.net"
	BOT_NAME   = "diegobot"
)

func main() {
	c := irc.SimpleClient(BOT_NAME, BOT_NAME)

	connectToServer(c)
	registerToChannels(c)
	c.HandleFunc("privmsg", handlers.MessageHandle)

	handleDisconnectedEvent(c)
}

func connectToServer(c *irc.Conn) {
	if err := c.ConnectTo(IRC_SERVER); err != nil {
		log.Fatal("Connection Error to: " + IRC_SERVER)
		os.Exit(1)
	}
	log.Println("Connected to IRC Server: " + IRC_SERVER)
}

func registerToChannels(c *irc.Conn) {
	channels := os.Args[1:]

	c.HandleFunc("connected", func(conn *irc.Conn, line *irc.Line) {

		for _, channel := range channels {
			conn.Join("#" + channel)
			log.Println("Joined channel: #" + channel)
		}
	})
}

func handleDisconnectedEvent(c *irc.Conn) {
	quit := make(chan bool)

	// what happens when we are disconnected?
	c.HandleFunc("disconnected", func(conn *irc.Conn, line *irc.Line) {
		log.Println("Disconnected to: " + IRC_SERVER)
		quit <- true
	})

	// wait for disconnect
	<-quit
}
