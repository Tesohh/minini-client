package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/Tesohh/minini-client/connection"
	"github.com/Tesohh/minini-client/message"
	"github.com/Tesohh/minini-client/rp"
	"github.com/Tesohh/minini-client/view/hud"
	"github.com/Tesohh/minini-client/view/login"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// m, err := render.MapFromFile("assets/maps/world1.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(m)
	// return
	file, err := os.OpenFile(".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Truncate(0)
	slogger := slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{}))
	slog.SetDefault(slogger)

	c := &connection.Client{Quitch: make(chan struct{}), Actions: connection.Actions}

	s := &connection.ServerConn{Addr: "localhost:8080"}
	err = s.Connect(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Conn.Close()

	// autologin
	go func() {
		username := os.Getenv("USERNAME")
		password := os.Getenv("PASSWORD")
		if username == "" || password == "" {
			return
		}

		c.Username = username
		s.Send(message.Msg{Action: "login", Data: map[string]any{"username": username, "password": password}})
	}()

	model := login.InitialModel(c, s, "login")
	rp.Global.TeaProg = tea.NewProgram(model, tea.WithAltScreen())
	if _, err := rp.Global.TeaProg.Run(); err != nil {
		log.Fatal("Couldn't start login screen")
	}

	if !c.Authenticated {
		os.Exit(0)
	}

	// fetch state from db
	s.Send(message.Msg{Action: "me.state"})

	g := hud.InitialModel(c, s)
	rp.Global.TeaProg = tea.NewProgram(g, tea.WithAltScreen())
	if _, err := rp.Global.TeaProg.Run(); err != nil {
		log.Fatal("Couldn't start game screen")
	}
}
