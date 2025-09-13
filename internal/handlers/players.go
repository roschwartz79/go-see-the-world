package handlers

import (
	"fmt"
	"go-see-the-world/internal/models"
	"go-see-the-world/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	TeamRepository *repository.TeamRepository
}

// NewPlayerHandler is the constructor for the handler.
// It takes a pointer to the shared Team instance.
func NewPlayerHandler(repo *repository.TeamRepository) *PlayerHandler {
	return &PlayerHandler{
		TeamRepository: repo,
	}
}

// CreatePlayer is a method on the TeamHandler struct.
func (h *PlayerHandler) CreatePlayer(c *gin.Context) {
	var newPlayer models.Player
	// c.BindJSON is a helper function that binds the JSON body to the struct
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call a method on the shared Team instance to add the player.
	h.TeamRepository.AddPlayer(newPlayer)

	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Player %s added successfully", newPlayer.Name)})
}
