package handlers

import (
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

// GetTeam handles retrieving all teams from the repository.
func (h *TeamHandler) GetTeam(c *gin.Context) {
	team := h.TeamRepo.GetTeam()
	c.JSON(http.StatusOK, team)
}
