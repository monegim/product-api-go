package model

import (
	"encoding/json"
	"io"
)

type Coffees []Coffee

func (c *Coffee) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(c)
}

type Coffee struct {
}
