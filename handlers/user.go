package handlers

import (
	"encoding/json"
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

func (c *User) SignUp(rw http.ResponseWriter, r *http.Request) {
	c.log.Info("Handle User | signup")
	body := AuthStruct{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		c.log.Error("Unable to decode JSON", "error", err)
		http.Error(rw, "Unable to parse request body", http.StatusInternalServerError)
		return
	}
}
func (c *User) SignIn(rw http.ResponseWriter, r *http.Request) {

}
func (c *User) SignOut(rw http.ResponseWriter, r *http.Request) {

}
