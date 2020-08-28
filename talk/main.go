package main

import (
	"talk/controller"
	"talk/middle"
	"talk/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	defer utils.CloseAllGrpcClientConnect()

	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middle.Cors)
	controller.RegisterController(router, controller.NewRoomController())
	router.Run(":8084")
}
