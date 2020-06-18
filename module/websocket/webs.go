package webs

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type Conn struct {
	Ws *websocket.Conn
}

func (t *Conn) Open(host, path string) error {
	var err error
	u := url.URL{Scheme: "ws", Host: host, Path: path}
	t.Ws, _, err = websocket.DefaultDialer.Dial(u.String(), nil)

	return err
}
func (t *Conn) Write(buf []byte) (int, error) {
	err := t.Ws.WriteMessage(websocket.TextMessage, buf)
	return len(buf), err
}

func (t *Conn) Read(buf []byte) (int, error) {
	_, message, err := t.Ws.ReadMessage()
	if err != nil {
		return 0, err
	}
	copy(buf, message)
	return len(message), nil
}

func (t *Conn) Close() error {
	return t.Ws.Close()
}
