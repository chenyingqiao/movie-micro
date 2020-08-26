package main

import (
	"talk/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	controller.RegisterController(router, controller.NewRoomController())
	router.Run(":8084")
}
