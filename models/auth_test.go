package models

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/go-ini/ini"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conf_location = "conf/app.ini"

func setup() {
	cwd, _ := os.Getwd()
	conf_full_path := path.Join(path.Dir(cwd), conf_location)
	file, err := ini.Load(conf_full_path)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	db_name := file.Section("database").Key("Path")
	db_path := path.Join(path.Join(cwd), "..", db_name.String())

	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		log.Printf("error opening database: %v", err)
	}

	auth := Auth{
		Username: "mostafa",
		Password: "pass123",
	}
	rows := db.Create(&auth).RowsAffected
	log.Printf("rows: %v", rows)
	// res := new(Auth)
	// rows1 := db.First(&res, "mostafa")
	// log.Printf("rows: %v", rows1)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}
