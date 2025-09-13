package models

type Team struct {
	Id      int64
	Name    string
	Players []Player
	Coaches []Coach
}

// NewTeam is the constructor function for the Team struct.
// It initializes a new Team and returns a pointer to it.
func NewTeam(id int64, name string) *Team {
	return &Team{
		Id:      id,
		Name:    name,
		Players: make([]Player, 0),
		Coaches: make([]Coach, 0),
	}
}
