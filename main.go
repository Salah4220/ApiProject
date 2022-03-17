package main

import (
	"ApiProject/config"
	"ApiProject/models"
	"log"
	"net/http"
	"testing"
	"os"
)

func TestMain(m *testing.M) {
  config.TestDatabaseInit()

  ret := m.Run()

  config.TestDatabaseDestroy()
  os.Exit(ret)
}

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	models.NewOeuvre(&models.Oeuvre{Title: "Bal du moulin de la", Artist: "Pierre-Auguste Renoir", Year: "1876"})
	models.NewOeuvre(&models.Oeuvre{Title: "Portrait du Docteur", Artist: "Vincent Van Gogh", Year: "1890"})
	models.NewOeuvre(&models.Oeuvre{Title: "Salvator Mundi", Artist: "Léonard de Vinci", Year: "1500"})
	
	log.Println(models.AllOeuvre())

	log.Println("Insertion avec succes")
	log.Println(models.AllOeuvre())
	models.DeleteOeuvreById(12)
	log.Println("Delete avec succes")
	log.Println(models.AllOeuvre())
	
	log.Fatal(http.ListenAndServe(":8080", router))
}