package connection

import (
	"encoding/json"
	"log/slog"
	"net"

	"github.com/Tesohh/minini-client/message"
)

type ServerConn struct {
	Addr string
	Conn net.Conn
}

func (s *ServerConn) Connect(c *Client) error {
	conn, err := net.Dial("tcp", s.Addr)
	if err != nil {
		return err
	}

	s.Conn = conn

	go s.ReadIncomingMessages(c)

	return nil
}

func (s *ServerConn) ReadIncomingMessages(c *Client) {
	buf := make([]byte, 2048)
	for {
		length, err := s.Conn.Read(buf)
		if err != nil {
			slog.Error("ServerConn.ReadIncomingMessages error while reading:", "error", err.Error())
		}

		var msg message.Msg
		err = json.Unmarshal(buf[:length], &msg)
		if err != nil {
			slog.Error("ServerConn.ReadIncomingMessages error while unmarshaling:", "error", err.Error())
		}

		go c.HandleMessage(msg)
	}
}

func (s *ServerConn) Send(m message.Msg) error {
	marsh, err := json.Marshal(m)
	if err != nil {
		return err
	}
	_, err = s.Conn.Write(marsh)
	return err
}
