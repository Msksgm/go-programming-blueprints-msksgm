package main

import "github.com/gorilla/websocket"

type room struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネルです
	forward chan []byte
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err ==nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg);
			err != nil {
				break
			}
	}
	c.socket.Close()
}