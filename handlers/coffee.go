package handlers

import "github.com/hashicorp/go-hclog"

type Coffee struct {
	log hclog.Logger
}

func NewCoffee(l hclog.Logger) *Coffee {
	return &Coffee{l}
}
