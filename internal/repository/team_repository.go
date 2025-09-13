package repository

import (
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
func (r *TeamRepository) AddPlayer(player models.Player) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Team.Players = append(r.Team.Players, player)
}

// AddCoach adds a new coach to the team in a thread-safe way.
func (r *TeamRepository) AddCoach(coach models.Coach) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Team.Coaches = append(r.Team.Coaches, coach)
}
