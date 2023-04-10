package main

import (
	"fmt"
	"io"
	"os"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monegim/product-api-go/models"
	"github.com/monegim/product-api-go/pkg/setting"
	"github.com/monegim/product-api-go/routers"
	log "github.com/sirupsen/logrus"
)

func init() {
	setting.Setup()
	models.Setup()
}
func main() {
	fmt.Println("Running main")
	gin.SetMode(setting.ServerSetting.RunMode)

	log := &log.Logger{
		Out:   io.MultiWriter(os.Stdout),
		Level: log.TraceLevel,
		Formatter: &log.TextFormatter{
			FullTimestamp: true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	}

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout * time.Millisecond,
		WriteTimeout:   writeTimeout * time.Millisecond,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
