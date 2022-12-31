package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testConfig struct {
	Name string
}

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

func TestLoadsConfigIntoStructOnStart(t *testing.T) {
	conf := setupTests(t)
	tc := &testConfig{}
	f, err := New(conf, tc)
	assert.NoError(t, err)
	t.Log(f)
	assert.Equal(t, "Mostafa", tc.Name)
}
