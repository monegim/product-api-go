package handlers

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type User struct {
	log hclog.Logger
}

type AuthStruct struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthResponse struct {
	UserID   int    `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}

func NewUser(l hclog.Logger) *User {
	return &User{
		l,
	}
}

func (c *User) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle User | unknown", "path", r.URL.Path)
	http.NotFound(rw, r)
}
