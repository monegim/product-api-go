package model

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var coffeesData = `
[
	{
		"id": 1,
		"name": "Latte",
		"price": 50.0
	},
	{
		"id": 2,
		"name": "Americano",
		"price": 30.0
	}
]
`

func TestCoffeesDeserializeFromJSON(t *testing.T) {
	c := Coffees{}
	err := c.FromJSON(bytes.NewReader([]byte(coffeesData)))
	assert.NoError(t, err)
}
