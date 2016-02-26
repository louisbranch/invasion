package client

import "golang.org/x/net/websocket"

type Client interface {
	Write([]byte) error
	Close() error
}

type WebsocketClient struct {
	conn *websocket.Conn
}

func NewWsClient(conn *websocket.Conn) *WebsocketClient {
	return &WebsocketClient{conn: conn}
}

func (c *WebsocketClient) Write(txt []byte) error {
	//FIXME
	return websocket.Message.Send(c.conn, string(txt))
}

func (c *WebsocketClient) Close() error {
	return c.conn.Close()
}
