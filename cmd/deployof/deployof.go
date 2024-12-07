package main

import (
	"flag"
	"fmt"
	"github.com/Juminiy/kube/cmd/deployof/handlers"
	_ "github.com/Juminiy/kube/cmd/deployof/sdk"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {}

func init() {
	flag.StringVar(&_host, "host", "0.0.0.0", "host to bind")
	flag.IntVar(&_port, "port", 7899, "port to listen")
	flag.Parse()

	app := gin.Default()
	app.Use(gin.Recovery())
	apiv1 := app.Group("/api/v1")
	apiv1.POST("/image/upload", handlers.ImageUpload)
	apiv1.GET("/image/download/tmp-url", handlers.ImageDownloadURL)
	log.Fatal(app.Run(fmt.Sprintf("%s:%d", _host, _port)))
}

var (
	_host string
	_port int
)
