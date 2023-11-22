package main

import (
	"github.com/Tesohh/minini-client/connection"
	"github.com/Tesohh/minini-client/message"
)

func main() {
	c := &connection.Client{Quitch: make(chan struct{}), Actions: connection.Actions}

	s := &connection.ServerConn{Addr: "localhost:8080"}
	s.Connect(c)
	defer s.Conn.Close()

	s.Send(message.Msg{Action: "login", Data: map[string]any{"username": "tesohh", "password": "zestfest"}})

	<-c.Quitch
}
