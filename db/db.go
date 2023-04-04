package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Open(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("reading db error: %s", err)
	}
	// defer db.Close()
	// var version string
	// err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	// if err != nil {
	//     log.Fatal(err)
	// }

	// fmt.Println(version)
	return db
}

func Insert(query string, db *sql.DB) error {
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
