package handlers

import (
	"net/http"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
	"github.com/gin-gonic/gin"
)

type StreamHandler struct {
	ChatUsecase domain.ChatUseCase
}

func NewStreamHandler(r *gin.RouterGroup, cu domain.ChatUseCase) {
	handler := &StreamHandler{
		ChatUsecase: cu,
	}
	streamGroup := r.Group("/streams")
	{
		streamGroup.POST("/start", handler.StartStream)
		streamGroup.POST("/end", handler.EndStream)
	}
}

// start stream handler
func (h *StreamHandler) StartStream(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user unathorized"})
		return
	}
	var payload struct {
		Title        string `json:"title" binding:"required"`
		Category     string `json:"category"`
		ThumbnailURL string `json:"thumbnail_url"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong format or title empty"})
		return
	}
	stream := &domain.Stream{
		ID:           userID.(string),
		Title:        payload.Title,
		Category:     payload.Category,
		ThumbnailURL: payload.ThumbnailURL,
	}
	if err := h.ChatUsecase.StartStream(stream); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start stream"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stream succes, your stream now !",
		"data": stream,
	})
}

// end stream handler
func (h *StreamHandler) EndStream(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if err := h.ChatUsecase.EndStream(userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to end stream"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stream ended"})
}

// get active stream
func (h *StreamHandler) GetLiveStream(c *gin.Context) {
	streams, err := h.ChatUsecase.GetActiveStreams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stream list"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success get list livestream",
		"data": streams,
	})
}
