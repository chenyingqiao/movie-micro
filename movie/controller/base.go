package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(engin gin.Engine)
}
