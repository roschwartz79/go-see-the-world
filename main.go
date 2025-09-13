package main

import (
	"fmt"
	"go-see-the-world/internal/handlers"
	"go-see-the-world/internal/models"
	"go-see-the-world/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a new Gin router
	router := gin.Default()

	// Initialize a Team
	team := &models.Team{
		Players: make([]models.Player, 0),
		Coaches: make([]models.Coach, 0),
	}

	teamRepo := repository.NewTeamRepository(team)

	teamHandler := handlers.NewTeamHandler(teamRepo)
	playerHandler := handlers.NewPlayerHandler(teamRepo)
	coachHandler := handlers.NewCoachHandler(teamRepo)

	// Define the routes for the application
	router.GET("/team", teamHandler.GetTeam)
	router.POST("/players", playerHandler.CreatePlayer)
	router.POST("/coaches", coachHandler.CreateCoach)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}
