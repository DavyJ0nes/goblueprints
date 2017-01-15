package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	// socket is the websocket for this client
	socket *websocket.Conn
	// send is the channel from where messages are sent
	send chan []byte
	// room is the room this client is chatting in
	room *room
}

// read takes message from room.forward channel
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

// write sends message to client send channel
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
