package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Game represents data about game detail
type game struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Publisher string  `json:"publisher"`
	Platform  string  `json:"platform"`
	Price     float64 `json:"price"`
}

// Games slice to seed record game data.
var games = []game{
	{ID: "1", Name: "Football Manager 2023", Publisher: "Sport Interactive", Platform: "PC", Price: 39.99},
	{ID: "2", Name: "Football Manager Mobile 2023", Publisher: "Sport Interactive", Platform: "Mobile", Price: 19.99},
	{ID: "3", Name: "FIFA 2023", Publisher: "EA Sport", Platform: "PC", Price: 29.99},
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.GET("/games/:id", getGameById)
	router.POST("/games", postGames)

	router.Run("localhost:8080")
}

// getGames responds with the list of all games as JSON.
func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

// getGamesById locates the game whose ID value matches the id
// parameter sent by client, then returns that game as a response
func getGameById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of games, looking for
	// an game whose ID value matches the parameter
	for _, a := range games {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Game not found"})
}

// postGames adds an game from JSON received in the request body.
func postGames(c *gin.Context) {
	var newGame game

	// Call BindJSON to bind the received JSON to newGame.
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	// Add the new game to the slice.
	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}
