package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"rest-mock/router"
)

func init() {
	log.Info("rest mock server starting...")
}

func main() {
	engine := gin.Default()
	router.InitRouteVersion1(engine.RouterGroup)
	engine.Run(":7070")
}
