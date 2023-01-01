package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/monegim/product-api-go/config"
	"github.com/nicholasjackson/env"
)

type Config struct {
	BindAddress string `json:"bind_address"`
}

var conf *Config
var logger hclog.Logger

var bindAddress string
var configFile = env.String("CONFIG_FILE", false, "./conf.json", "Path to JSON encoded file")

// var test = "test"

func main() {
	// var testInt = env.Int("TEST_FILE", false, 1, "Path to JSON encoded file")
	logger = hclog.Default()
	// fmt.Println(*testInt)
	err := env.Parse()
	if err != nil {
		logger.Error("Error parsing flags", "error", err)
		os.Exit(1)
	}
	r := mux.NewRouter()
	// fmt.Println(test)

	conf = &Config{
		BindAddress: bindAddress,
	}
	// r.Use(cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
	// 	AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	// }).Handler)
	if conf.BindAddress == "" {
		c, err := config.New(*configFile, conf)
		if err != nil {
			logger.Error("Unable to load config file", "error", err)
			os.Exit(1)
		}
		defer c.Close()
	}
	err = http.ListenAndServe(conf.BindAddress, r)
	if err != nil {
		logger.Error("Unable to start server", "bind", conf.BindAddress, "error", err)
	}
}
