package main

import (
	"log"
	"net/http"

	"goblueprints/trace"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// room represents a room that users chat in
type room struct {
	// forward is a channel that holds messages
	// that should be forwarded to other clients
	forward chan []byte
	// using 2 chans to avoid locking issues with concurrent access
	// join is a chanel for clients wanting to join the room
	join chan *client
	// leave is a chanel for clients wanting to leave the room
	leave chan *client
	// clients holds all current clients in this room
	clients map[*client]bool
	// tracer receives trace info of activity in the room
	tracer trace.Tracer
}

// newRoom is a helper for creating a new room object
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
	}
}

// run is a constantly running function that handles
// client interaction and message routing
func (r *room) run() {
	for {
		select {
		// joining
		case client := <-r.join:
			// add client to clients map
			r.clients[client] = true
			r.tracer.Trace("New Client Joined")
		// leaving
		case client := <-r.leave:
			// remove client from clients map
			delete(r.clients, client)
			// close client channel
			close(client.send)
			r.tracer.Trace("Client Left")
		// forward messages to all clients
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", string(msg))
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- sent to client")
			}
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

// ServeHTTP allows a room to act like an http.Handler
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// need to upgrade HTTP connection to allow websockets
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("room ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}

	r.join <- client
	defer func() { r.leave <- client }()
	// run write in seperate thread
	go client.write()
	client.read()
}
