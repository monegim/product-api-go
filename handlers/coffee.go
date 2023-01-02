package handlers

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type Coffee struct {
	log hclog.Logger
}

func NewCoffee(l hclog.Logger) *Coffee {
	return &Coffee{l}
}

func (c *Coffee) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle Coffee")
	fmt.Fprintf(rw, "coffee")

}

func (c *Coffee) CreateCoffee(_ int, rw http.ResponseWriter, r *http.Request) {

}
