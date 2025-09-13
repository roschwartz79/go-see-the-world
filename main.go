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

	// Initialize a Team Slice
	team := []*models.Team{}

	teamRepo := repository.NewTeamRepository(team)

	teamHandler := handlers.NewTeamHandler(teamRepo)
	playerHandler := handlers.NewPlayerHandler(teamRepo)
	coachHandler := handlers.NewCoachHandler(teamRepo)

	// Define the routes for the application
	router.GET("/teams", teamHandler.GetTeams)
	router.POST("/teams", teamHandler.CreateTeam)
	router.POST("/players", playerHandler.CreatePlayer)
	router.POST("/coaches", coachHandler.CreateCoach)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}
