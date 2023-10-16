package handler

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"log"
)

type Message struct {
	From    int    `json:"from"`
	To      int    `json:"to"`
	Message string `json:"message"`
}

type ChatRoom struct {
	ChatID     uuid.UUID `json:"chatID"`
	clients    map[*websocket.Conn]bool
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	broadcast  chan Message
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		ChatID:     uuid.New(),
		clients:    make(map[*websocket.Conn]bool),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		broadcast:  make(chan Message),
	}
}

var chatRoom = NewChatRoom()

func (h *Handler) sendMessage(conn *websocket.Conn) {
	go chatRoom.Run()

	chatRoom.register <- conn
	defer func() {
		chatRoom.unregister <- conn
		conn.Close()
	}()

	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println("Error reading message from client:", err)
			break
		}
		chatRoom.broadcast <- message
	}
}

func (c *ChatRoom) Run() {
	for {
		select {
		case conn := <-c.register:
			c.clients[conn] = true
		case conn := <-c.unregister:
			delete(c.clients, conn)
		case message := <-c.broadcast:
			for conn := range c.clients {
				err := conn.WriteJSON(message)
				if err != nil {
					log.Println("Error sending message to client:", err)
					conn.Close()
					delete(c.clients, conn)
				}
			}
		}
	}
}
