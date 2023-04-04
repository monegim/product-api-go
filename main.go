package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Coffee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.GET("teacher", GetCoffeeHandler)
	err := r.Run(":3333")
	if err != nil {
		log.Fatalf("could not start the server: %s", err)
	}
}

func GetCoffeeHandler(c *gin.Context) {
	
}
