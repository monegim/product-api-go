package models

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/go-ini/ini"
)

var conf_location = "conf/app.ini"

func setup() {
	cwd, _ := os.Getwd()
	conf_full_path := path.Join(path.Dir(cwd),conf_location)
	file, err := ini.Load(conf_full_path)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	db_name := file.Section("database").Key("Path")
	db_path := path.Join(path.Join(cwd),"..", db_name.String())
	fmt.Printf("file: %v\n", db_path)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}
