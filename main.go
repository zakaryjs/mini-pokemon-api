package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID       string `json:"id"`
	Species  string `json:"species"`
	Nickname string `json:"nickname"`
	Level    int    `json:"level"`
}

var PokemonStruct = []Pokemon{
	{ID: "1", Species: "Samurott", Nickname: "ASH", Level: 100},
	{ID: "2", Species: "Xerneas", Nickname: "Xerneas", Level: 100},
	{ID: "3", Species: "Pikachu", Nickname: "Bolt", Level: 25},
}

func main() {
	router := gin.Default()
	router.GET("/pokemon", GetPokemon)
	router.GET("/pokemon/:id", GetPokemonByID)
	router.POST("/pokemon", CreatePokemon)

	router.Run("localhost:8080")
}

func GetPokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, PokemonStruct)
}

func CreatePokemon(c *gin.Context) {
	var newPokemon Pokemon

	if err := c.BindJSON(&newPokemon); err != nil {
		return
	}

	PokemonStruct = append(PokemonStruct, newPokemon)
	c.IndentedJSON(http.StatusCreated, newPokemon)
}

func GetPokemonByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range PokemonStruct {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "There was no Pokemon found with that ID."})
}
