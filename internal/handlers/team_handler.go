package handlers

import (
	"fmt"
	"go-see-the-world/internal/models"
	"go-see-the-world/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeamHandler contains the logic for handling team-related requests.
// It holds a reference to the shared TeamRepository.
type TeamHandler struct {
	TeamRepo *repository.TeamRepository
}

// NewTeamHandler is the constructor for the handler.
func NewTeamHandler(repo *repository.TeamRepository) *TeamHandler {
	return &TeamHandler{
		TeamRepo: repo,
	}
}

// GetTeams handles retrieving all teams from the repository.
func (h *TeamHandler) GetTeams(c *gin.Context) {
	teams := h.TeamRepo.GetTeams()
	c.JSON(http.StatusOK, teams)
}

func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var newTeam models.Team
	// c.BindJSON is a helper function that binds the JSON body to the struct
	if err := c.BindJSON(&newTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.TeamRepo.AddTeam(&newTeam)

	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf(
		"Team %s with id %d added successfully", newTeam.Name, newTeam.Id)},
	)
}
