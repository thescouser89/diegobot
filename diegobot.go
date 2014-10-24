package main

import (
	"github.com/diegobot/handlers"
	irc "github.com/fluffle/goirc/client"
	"log"
)

func main() {
	c := irc.SimpleClient("diegobot", "diegobot")

	quit := make(chan bool)

	c.HandleFunc("disconnected",
		func(conn *irc.Conn, line *irc.Line) {
			log.Println("Disconnected")
			quit <- true
		})

	handlers.Handle(c)

	if err := c.ConnectTo("irc.freenode.net"); err != nil {
		log.Fatal("Connection Error")
	}

	// wait for disconnect
	<-quit
}
