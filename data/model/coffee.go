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
}
