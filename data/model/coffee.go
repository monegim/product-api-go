package model

import (
	"encoding/json"
	"io"
)

type Coffees []Coffee

func (c *Coffees) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(c)
}

func (c *Coffees) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

type Coffee struct {
	ID          int                `db:"id" json:"id"`
	Name        string             `db:"name" json:"name"`
	Ingredients []CoffeeIngredient `json:"ingredients"`
}

func (c *Coffee) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(c)
}

func (c *Coffee) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

type CoffeeIngredient struct {
}
