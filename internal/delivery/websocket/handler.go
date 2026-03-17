package websocket

import (
	"log"
	"net/http"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
	"github.com/gin-gonic/gin"
	gorillaWs "github.com/gorilla/websocket"
)

var allowedOrigins = map[string]bool{
	"http://localhost:3000": true,
	// "https://rawve.live":    true,   <-- this domain prodcution
}

var upgrader = gorillaWs.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if allowedOrigins[origin] {
			return true
		}
		log.Println("Origin not allowed:", origin)
		return false
	},
}

func ServeWS(hub *Hub, c *gin.Context) {
	streamID := c.Param("stream_id")
	userID, exist := c.Get("user_id")
	if !exist {
		userID = c.Query("user_id")
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user unathorized"})
			return
		}
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Fail upgrade to websocket", err)
		return
	}
	client := &Client{
		hub:      hub,
		conn:     conn,
		Send:     make(chan *domain.Message, 256),
		StreamID: streamID,
		UserID:   userID.(string),
	}
	client.hub.register <- client
	go client.ReadPump()
	go client.WritePump()
}
