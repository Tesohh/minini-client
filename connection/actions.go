package connection

import (
	"fmt"
	"log/slog"

	"github.com/Tesohh/minini-client/data"
	"github.com/Tesohh/minini-client/message"
	"github.com/Tesohh/minini-client/rp"
	tea "github.com/charmbracelet/bubbletea"
)

func sendToGui(msg tea.Msg) {
	if rp.Global.TeaProg != nil {
		rp.Global.TeaProg.Send(msg)
	}
}

type ActionFunc func(c *Client, m message.Msg) error

var Actions = map[string]ActionFunc{
	"error": func(c *Client, m message.Msg) error {
		sendToGui(OkMsg(false))
		slog.Error("Error incoming from server", "error", m.Data["error"])
		return nil
	},
	"login.ok": func(c *Client, m message.Msg) error {
		type loginOKMsg struct {
			PlayerID string `json:"playerid"`
		}
		data, err := message.Data[loginOKMsg](m)
		if err != nil {
			return err
		}
		c.Authenticated = true
		c.PlayerID = data.PlayerID
		fmt.Println("successfully logged in")
		sendToGui(OkMsg(true))
		return nil
	},
	"me.state": func(c *Client, m message.Msg) error {
		data, err := message.Data[data.User](m)
		if err != nil {
			return err
		}
		c.State = *data
		return nil
	},
}

type OkMsg bool
