package handlers

import (
	"net/http"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUseCase
}

func NewUserHandler(r *gin.RouterGroup, us domain.UserUseCase) {
	handler := &UserHandler{
		userUsecase: us,
	}
	r.POST("/webhooks/clerk", handler.HandlerClerkWebhook)
}

func (h *UserHandler) HandlerClerkWebhook(c *gin.Context) {
	var payload struct {
		Data struct {
			ID           string `json:"id"`
			Username     string `json:"username"`
			EmailAddress string `json:"email_address"`
			ImageURL     string `json:"image_url"`
		} `json:"data"`
		Type string `json:"type"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	user := &domain.User{
		ID:        payload.Data.ID,
		Username:  payload.Data.Username,
		Email:     payload.Data.EmailAddress,
		AvatarURL: payload.Data.ImageURL,
	}
	if err := h.userUsecase.SyncUserFromAuth(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if payload.Type == "user.created" || payload.Type == "user.updated" {
		user := &domain.User{
			ID:        payload.Data.ID,
			Username:  payload.Data.Username,
			Email:     payload.Data.EmailAddress,
			AvatarURL: payload.Data.ImageURL,
		}
		if err := h.userUsecase.SyncUserFromAuth(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to syncron user from auth"})
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "webhook successfully processed"})
}
