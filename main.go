package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/Tesohh/minini-client/connection"
	"github.com/Tesohh/minini-client/message"
	"github.com/Tesohh/minini-client/rp"
	"github.com/Tesohh/minini-client/view/login"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	file, err := os.OpenFile(".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	slogger := slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{}))
	slog.SetDefault(slogger)

	c := &connection.Client{Quitch: make(chan struct{}), Actions: connection.Actions}

	s := &connection.ServerConn{Addr: "localhost:8080"}
	err = s.Connect(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Conn.Close()

	go func() {
		username := os.Getenv("USERNAME")
		password := os.Getenv("PASSWORD")
		if username == "" || password == "" {
			return
		}

		s.Send(message.Msg{Action: "login", Data: map[string]any{"username": username, "password": password}})
	}()

	model := login.InitialModel(s, "login")
	rp.Global.TeaProg = tea.NewProgram(model, tea.WithAltScreen())
	if _, err := rp.Global.TeaProg.Run(); err != nil {
		log.Fatal("Couldn't start login screen")
	}

	<-c.Quitch
}
