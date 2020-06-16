package webs

import (
	"net/url"

	"golang.org/x/net/websocket"
)

type Conn struct {
	Ws *websocket.Conn
}

func (t *Conn) Open(host, path string) error {
	var err error
	u := url.URL{Scheme: "ws", Host: host, Path: path}
	t.Ws, err = websocket.Dial(u.String(), "", "http://"+host+"/")

	return err
}

func (t *Conn) Write(buf []byte) (int, error) {
	return t.Ws.Write(buf)
}

func (t *Conn) Read(buf []byte) (int, error) {
	//TODO:由于网络原因，可能一包数据变成两包
	return t.Ws.Read(buf)
}

func (t *Conn) Close() error {
	return t.Ws.Close()
}
