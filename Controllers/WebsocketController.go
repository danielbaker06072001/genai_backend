package Controllers

import (
	"log"
	"net/http"

	"genai2025/Worker" // import your Worker package

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade WebSocket:", err)
		return
	}

	// ✅ Register connection
	Worker.Clients[username] = conn
	log.Println("Registered WebSocket for user:", username)

	defer func() {
		conn.Close()
		delete(Worker.Clients, username)
		log.Println("WebSocket closed and removed for user:", username)
	}()

	// Optional: read and echo messages
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		log.Printf("Received from %s: %s", username, msg)

		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println("WebSocket write error:", err)
			break
		}
	}
}
