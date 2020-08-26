package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(engin *gin.Engine)
}

func RegisterController(engin *gin.Engine, controller interface{}) {
	c := controller.(Controller)
	c.Register(engin)
}