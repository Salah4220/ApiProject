package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Oeuvre struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int64  `json:"year"`
}

var Oeuvres = []Oeuvre{
	{ID: "1", Title: "Bal du moulin de la Galette ", Artist: "Pierre-Auguste Renoir", Year: 1876},
	{ID: "2", Title: "Portrait du Docteur Gachet", Artist: "Vincent Van Gogh", Year: 1890},
	{ID: "3", Title: "Salvator Mundi ", Artist: "Léonard de Vinci", Year: 1500},
}

// renvoi la list des oeuvre en json
func getOeuvres(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Oeuvres)
}

func main() {
	router := gin.Default()
	router.GET("/Oeuvres", getOeuvres)

	router.Run("localhost:8080")
}
