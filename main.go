package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	Species  string `json:"species"`
	Nickname string `json:"nickname"`
	Level    int    `json:"level"`
}

var PokemonStruct = []Pokemon{
	{Species: "Samurott", Nickname: "ASH", Level: 100},
	{Species: "Xerneas", Nickname: "Xerneas", Level: 100},
	{Species: "Pikachu", Nickname: "Bolt", Level: 25},
}

func main() {
	router := gin.Default()
	router.GET("/pokemon", getPokemon)

	router.Run("localhost:8080")
}

func getPokemon(c *gin.Context) {
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
