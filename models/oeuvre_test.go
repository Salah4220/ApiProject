package models

import (
	"ApiProject/config"
	"os"
	"reflect"
	"testing"
)

func TestNewOeuvre(t *testing.T) {

	oeuvres := []Oeuvre{
		{},
		{Title: "Name", Artist: "Name", Year: "Year"},
	}
	for _, o := range oeuvres {
		NewOeuvre(&o)

	}

	rows, _ := config.Db().Query("SELECT * FROM oeuvre")

	var rowsCount int

	for rows.Next() {
		rowsCount += 1
	}

	if rowsCount != 2 {
		t.Errorf("Database does not have 2 records, it has %v records", rowsCount)
	}
}

func TestFindOeuvreById(t *testing.T) {
	oeuvre := &Oeuvre{Title: "Name", Artist: "Name", Year: "Year"}

	NewOeuvre(oeuvre)
	oeuvreFound := FindOeuvreById(1)

	if oeuvre.ID != oeuvreFound.ID {
		t.Error("Couldn't find oeuvre by id")
	}
}

func TestAllOeuvre(t *testing.T) {
	ResetTableOeuvre()

	var oeuvres Oeuvres

	oeuvre := &Oeuvre{Title: "Name", Artist: "Name", Year: "Year"}

	NewOeuvre(oeuvre)

	oeuvres = append(oeuvres, *FindOeuvreById(1))

	if !reflect.DeepEqual(&oeuvres, AllOeuvre()) {
		t.Error("Couldn't find correct oeuvre from All")
	}
}

func TestMain(m *testing.M) {
	config.TestDatabaseInit()

	ret := m.Run()

	config.TestDatabaseDestroy()
	os.Exit(ret)
}
func ResetTableOeuvre() {
	config.Db().Exec("DROP TABLE oeuvre; CREATE TABLE IF NOT EXISTS oeuvre(id serial,TitleOeuvre varchar(200), Artist varchar(200), Year varchar(20), constraint pk primary key(id))")
}