package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

)

var db *sql.DB

func TestDatabaseInit() {
	connection, err := sql.Open("postgres", "user=postgres password=123 sslmode=disable")
	_, err = connection.Exec("CREATE DATABASE goapi")

	connection.Close()

	db, err = sql.Open("postgres", "user=postgres password=123 dbname=goapi sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	createOeuvresTable()

}

func TestDatabaseDestroy() {
	db.Close()

	connection, err := sql.Open("postgres", "user=postgres password=123 sslmode=disable")
	_, err = connection.Exec("DROP DATABASE goapi")

	if err != nil {
		log.Fatal(err)
	}
}


func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "user=postgres password=123 dbname=goapi sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	
	createOeuvresTable()
}

func createOeuvresTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS oeuvre(id serial,TitleOeuvre varchar(200), Artist varchar(200), Year varchar(20), constraint pk primary key(id))")

	if err != nil {
		log.Fatal("erreur creation")
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}