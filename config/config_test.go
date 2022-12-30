package config

import (
	"os"
	"testing"
)

func setupTests(t *testing.T) string {
	conf := "/tmp/config.json"
	os.Remove(conf)
	f, err := os.OpenFile(conf, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	f.WriteString(`{"Name": "Mostafa"}`)
	return conf
}
