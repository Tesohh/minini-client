package connection

import (
	"log/slog"

	"github.com/Tesohh/minini-client/data"
	"github.com/Tesohh/minini-client/message"
)

type Client struct {
	Authenticated bool
	Actions       map[string]ActionFunc
	PlayerID      string
	Username      string
	State         data.User
	Quitch        chan struct{}
}

func (c *Client) HandleMessage(msg message.Msg) {
	act, ok := c.Actions[msg.Action]
	if !ok {
		slog.Warn("Client.HandleMessage error: action doesn't exist", "action", msg.Action)
		return
	}
	err := act(c, msg)
	if err != nil {
		slog.Error("Error while executing action", "action", msg.Action, "error", err)
	}
}
