package repository

import (
	"fmt"
	"go-see-the-world/internal/models"
	"sync"
)

// TeamRepository handles all data access for the Team.
// It contains the shared data and its mutex.
type TeamRepository struct {
	Teams []*models.Team
	mu    sync.Mutex
}

// NewTeamRepository creates and returns a new repository instance.
func NewTeamRepository(teams []*models.Team) *TeamRepository {
	return &TeamRepository{
		Teams: teams,
	}
}

// GetTeams retrieves a safe copy of the teams
func (r *TeamRepository) GetTeams() []models.Team {
	r.mu.Lock()
	defer r.mu.Unlock()

	teams := make([]models.Team, len(r.Teams))

	for i, team := range r.Teams {
		teams[i] = *team
	}

	return teams
}

func (r *TeamRepository) AddTeam(team *models.Team) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Teams = append(r.Teams, team)
}

// AddPlayer adds a new player to the team in a thread-safe way.
func (r *TeamRepository) AddPlayer(player models.Player, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, team := range r.Teams {
		if team.Id == id {
			// Found the team, now add the player.
			team.Players = append(team.Players, player)
			return nil // Success, return no error
		}
	}

	return fmt.Errorf("team with id %d not found", id)
}

// AddCoach adds a new coach to the team in a thread-safe way.
func (r *TeamRepository) AddCoach(coach models.Coach, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, team := range r.Teams {
		if team.Id == id {
			team.Coaches = append(team.Coaches, coach)
			return nil
		}
	}

	return fmt.Errorf("team with id %d not found", id)
}
