package handlers

import (
	"fmt"
	"go-see-the-world/internal/models"
	"go-see-the-world/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoachHandler struct {
	TeamRepository *repository.TeamRepository
}

// NewCoachHandler is the constructor for the handler.
// It takes a pointer to the shared Team instance.
func NewCoachHandler(teamRepo *repository.TeamRepository) *CoachHandler {
	return &CoachHandler{
		TeamRepository: teamRepo,
	}
}

// CreateCoach is a method on the CoachHandler struct.
func (h *CoachHandler) CreateCoach(c *gin.Context) {
	var newCoach models.Coach
	// c.BindJSON is a helper function that binds the JSON body to the struct
	if err := c.BindJSON(&newCoach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call a method on the shared Team instance to add the player.
	h.TeamRepository.AddCoach(newCoach)

	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Coach %s added successfully", newCoach.Name)})
}
