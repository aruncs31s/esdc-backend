package handler

import (
	"esdc-backend/internal/model"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatHub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan model.Message
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	messages   []model.Message
	mu         sync.RWMutex
}

var hub = &ChatHub{
	clients:    make(map[*websocket.Conn]bool),
	broadcast:  make(chan model.Message),
	register:   make(chan *websocket.Conn),
	unregister: make(chan *websocket.Conn),
	messages:   []model.Message{},
}

func init() {
	go hub.run()
	
	// Add mock messages
	hub.messages = []model.Message{
		{
			ID:        uuid.New().String(),
			UserID:    "1",
			Username:  "Alice",
			Text:      "Hey everyone! Welcome to the chat!",
			Timestamp: time.Now().Add(-10 * time.Minute),
		},
		{
			ID:        uuid.New().String(),
			UserID:    "2",
			Username:  "Bob",
			Text:      "Thanks! Excited to be here.",
			Timestamp: time.Now().Add(-8 * time.Minute),
		},
	}
}

func (h *ChatHub) run() {
	for {
		select {
		case conn := <-h.register:
			h.clients[conn] = true
		case conn := <-h.unregister:
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				conn.Close()
			}
		case message := <-h.broadcast:
			h.mu.Lock()
			h.messages = append(h.messages, message)
			if len(h.messages) > 100 {
				h.messages = h.messages[1:]
			}
			h.mu.Unlock()
			
			for conn := range h.clients {
				if err := conn.WriteJSON(message); err != nil {
					conn.Close()
					delete(h.clients, conn)
				}
			}
		}
	}
}

// HandleWebSocket godoc
// @Summary WebSocket chat endpoint
// @Description Establish WebSocket connection for real-time chat
// @Tags chat
// @Accept json
// @Produce json
// @Success 101 {string} string "Switching Protocols"
// @Router /ws [get]
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	
	hub.register <- conn
	
	// Send message history
	hub.mu.RLock()
	for _, msg := range hub.messages {
		conn.WriteJSON(msg)
	}
	hub.mu.RUnlock()
	
	defer func() {
		hub.unregister <- conn
	}()
	
	for {
		var msg model.Message
		if err := conn.ReadJSON(&msg); err != nil {
			break
		}
		
		msg.ID = uuid.New().String()
		msg.Timestamp = time.Now()
		hub.broadcast <- msg
	}
}

// GetMessages godoc
// @Summary Get chat messages
// @Description Retrieve all chat messages history
// @Tags chat
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Messages retrieved successfully"
// @Router /messages [get]
func GetMessages(c *gin.Context) {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	
	c.JSON(200, gin.H{"data": hub.messages})
}
