package models

import (
	"ApiProject/config"
	"log"
)

type Oeuvre struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   string  `json:"year"`
}


type Oeuvres []Oeuvre

func NewOeuvre(c *Oeuvre) {
	if c == nil {
		log.Fatal(c)
	}

	err := config.Db().QueryRow("INSERT INTO Oeuvre (TitleOeuvre, Artist, Year) VALUES ($1,$2,$3) RETURNING id;", c.Title, c.Artist, c.Year).Scan(&c.ID)

	if err != nil {
		log.Fatal(err)
	}
}

func FindOeuvreById(id int) *Oeuvre {
	var oeuvre Oeuvre

	row := config.Db().QueryRow("SELECT * FROM Oeuvre WHERE id = $1;", id)
	err := row.Scan(&oeuvre.ID, &oeuvre.Artist, &oeuvre.Title, &oeuvre.Year)

	if err != nil {
		log.Fatal(err)
	}

	return &oeuvre
}

func AllOeuvre() *Oeuvres {
	var oeuvres Oeuvres

	rows, err := config.Db().Query("SELECT * FROM Oeuvre")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var o Oeuvre

		err := rows.Scan(&o.ID, &o.Artist, &o.Title, &o.Year)

		if err != nil {
			log.Fatal(err)
		}

		oeuvres = append(oeuvres, o)
	}

	return &oeuvres
}

func UpdateOeuvre(oeuvre *Oeuvre) {


	stmt, err := config.Db().Prepare("UPDATE oeuvre SET TitleOeuvre=$1, Artist=$2, Year=$3, WHERE id=$4;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(oeuvre.Title, oeuvre.Artist, oeuvre.Year, oeuvre.ID)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteOeuvreById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM oeuvre WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}