package db

import (
	"log"
	"testing"
)

func TestOpen(t *testing.T) {
	path := "./foo.db"
	_ = Open(path)

}

func TestInsert(t *testing.T) {
	path := "./foo.db"
	sts := `
	DROP TABLE IF EXISTS cars;
	CREATE TABLE cars(id INTEGER PRIMARY KEY, name TEXT, price INT);
	INSERT INTO cars(name, price) VALUES('Audi',52642);
	`
	db := Open(path)
	err := Insert(sts, db)
	if err != nil {
		log.Fatal(err)
	}
}
