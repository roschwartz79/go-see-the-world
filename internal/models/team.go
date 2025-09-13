package models

type Team struct {
	Name    string
	Players []Player
	Coaches []Coach
}

// NewTeam is the constructor function for the Team struct.
// It initializes a new Team and returns a pointer to it.
func NewTeam() *Team {
	return &Team{
		Players: make([]Player, 0),
		Coaches: make([]Coach, 0),
	}
}
