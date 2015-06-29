package main

import (
	irc "github.com/fluffle/goirc/client"
	"github.com/thescouser89/diegobot/handlers"
	"log"
	"os"
	"time"
)

const (
	IRC_SERVER = "irc.freenode.net"
	BOT_NAME   = "diegobot"
)

func main() {
	c := irc.SimpleClient(BOT_NAME, BOT_NAME)

        // we only need to register the handler and specify which channel
        // to join once. after a reconnect we don't need to do that again,
        // if we do it again, for every message the bot will reply twice
        already_signed_in := false

	for {
		connectToServer(c)
                if !already_signed_in {
                    registerToChannels(c)
                    c.HandleFunc("privmsg", handlers.MessageHandle)
                    already_signed_in = true
                }

		handleDisconnectedEvent(c)
		// give it a delay before we try to reconnect
		time.Sleep(time.Minute)
		log.Println("Reconnecting now...")
	}
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

	if len(channels) == 0 {
		log.Fatal("You need to specify a channel! :: Usage: diegobot mcgillece")
		os.Exit(1)
	}

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
