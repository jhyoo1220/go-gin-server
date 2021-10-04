package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jhyoo1220/go-gin-server/internal/app/sample"
	"log"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/log/ping"},
	}))
	r.Use(gin.Recovery())

	r.GET("/log/do-nothing", sample.ReadLogAndDoNothing)
	r.GET("/log/kinesis/:stream_name", sample.SendToKinesis)
	r.GET("/log/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Failed to listening at 0.0.0.0:8080!")
	}
}
