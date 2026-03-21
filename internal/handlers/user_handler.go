package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUseCase
}

func NewUserHandler(r *gin.RouterGroup, us domain.UserUseCase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	r.POST("/webhooks/clerk", handler.HandlerClerkWebhook)
}

func generateStreamKey() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return "sk_rawve_" + hex.EncodeToString(bytes)
}

// HandlerClerkWebhook godoc
// @Summary      Clerk Webhook Handler
// @Description  Receives user lifecycle events from Clerk authentication service
// @Tags         webhooks
// @Accept       json
// @Produce      json
// @Param        payload  body      object  true  "Clerk webhook payload"
// @Success      200      {object}  map[string]interface{}  "message"
// @Failure      400      {object}  map[string]interface{}  "Error"
// @Failure      500      {object}  map[string]interface{}  "error: Failed to syncron user from auth"
// @Router       /webhooks/clerk [post]
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
	if payload.Type == "user.created" || payload.Type == "user.updated" {
		user := &domain.User{
			ID:        payload.Data.ID,
			Username:  payload.Data.Username,
			Email:     payload.Data.EmailAddress,
			AvatarURL: payload.Data.ImageURL,
		}
		if payload.Type == "user.created" {
			user.StreamKey = generateStreamKey()
		}
		if err := h.UserUsecase.SyncUserFromAuth(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to syncron user from auth"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "webhook successfully processed"})
}

// update profile
// SetupProfile godoc
// @Summary      Setup user profile
// @Description  Updates user profile with display name, bio, and category
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      object{display_name=string,bio=string,category=string}  true  "Profile details"
// @Success      200      {object}  map[string]interface{}  "message"
// @Failure      400      {object}  map[string]interface{}  "error: Wrong format or Display Name Empty"
// @Failure      401      {object}  map[string]interface{}  "error: Unathorized Access"
// @Failure      500      {object}  map[string]interface{}  "error: Failed to update profile"
// @Security     BearerAuth
// @Router       /profile/setup [put]
func (h *UserHandler) SetupProfile(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized Access"})
		return
	}
	var payload struct {
		DisplayName string `json:"display_name" binding:"required"`
		Bio         string `json:"bio"`
		Category    string `json:"category"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong format or Display Name Empty"})
		return
	}
	if err := h.UserUsecase.UpdateProfile(userID.(string), payload.DisplayName, payload.Bio, payload.Category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully, WELCOME TO RAWVE"})
}
