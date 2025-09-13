package repository

import (
	"go-see-the-world/internal/models"
	"sync"
)

// TeamRepository handles all data access for the Team.
// It contains the shared data and its mutex.
type TeamRepository struct {
	Team *models.Team
	mu   sync.Mutex
}

// NewTeamRepository creates and returns a new repository instance.
func NewTeamRepository(team *models.Team) *TeamRepository {
	return &TeamRepository{
		Team: team,
	}
}

// GetTeam retrieves a safe copy of the team
func (r *TeamRepository) GetTeam() models.Team {
	r.mu.Lock()
	defer r.mu.Unlock()

	players := make([]models.Player, len(r.Team.Players))
	copy(players, r.Team.Players)

	coaches := make([]models.Coach, len(r.Team.Coaches))
	copy(coaches, r.Team.Coaches)

	return models.Team{
		Players: players,
		Coaches: coaches,
	}
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
