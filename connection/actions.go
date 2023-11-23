package connection

import (
	"fmt"
	"log/slog"

	"github.com/Tesohh/minini-client/message"
	"github.com/Tesohh/minini-client/rp"
)

type ActionFunc func(c *Client, m message.Msg) error

var Actions = map[string]ActionFunc{
	"error": func(c *Client, m message.Msg) error {
		rp.Global.TeaProg.Send(OkMsg(false))
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
		rp.Global.TeaProg.Send(OkMsg(true))
		return nil
	},
}

type OkMsg bool
