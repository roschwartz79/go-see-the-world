package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Player struct represents a hockey player with a name and position.
type Player struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

// Coach struct represents a hockey coach with a name and position.
type Coach struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

// Team struct holds the players and coaches for a single team.
type Team struct {
	Players []Player `json:"players"`
	Coaches []Coach  `json:"coaches"`
}

// In-memory storage for our team data.
var team = Team{}
var mu sync.Mutex // Mutex to prevent race conditions when modifying `team`.

func main() {
	// Initialize a new Gin router
	router := gin.Default()

	// Define the routes for the application
	router.GET("/team", getTeamHandler)
	router.POST("/players", createPlayerHandler)
	router.POST("/coaches", createCoachHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}

func getTeamHandler(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	c.JSON(http.StatusOK, team)
}

func createPlayerHandler(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var newPlayer Player
	// c.BindJSON is a helper function that binds the JSON body to the struct
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team.Players = append(team.Players, newPlayer)
	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Player %s added successfully", newPlayer.Name)})
}

func createCoachHandler(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var newCoach Coach
	if err := c.BindJSON(&newCoach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team.Coaches = append(team.Coaches, newCoach)
	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Coach %s added successfully", newCoach.Name)})
}
