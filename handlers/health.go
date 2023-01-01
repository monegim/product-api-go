package handlers

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type Health struct {
	logger hclog.Logger
}

func NewHealth(l hclog.Logger) *Health {
	return &Health{l}
}

func (h *Health) Liveness(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%s", "ok")
}
