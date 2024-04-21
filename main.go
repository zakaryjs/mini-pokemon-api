package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "-"
	port     = 0
	user     = ""
	password = ""
	dbname   = ""
)

type Pokemon struct {
	ID       string `json:"id"`
	Species  string `json:"species"`
	Nickname string `json:"nickname"`
	Level    int    `json:"level"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"string"`
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
	router.GET("/users", GetUsers)
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

func GetUsers(c *gin.Context) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error connecting to the database"})
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM "user"`)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error executing query"})
		return
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error scanning row"})
			return
		}

		users = append(users, user)
	}
	c.IndentedJSON(http.StatusOK, users)
}
