package middle

import (
	"fmt"
	"net/http"
	"talk/rpc/client"
	"talk/rpc/protos"
	"talk/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddle(c *gin.Context) {
	jwt := c.GetHeader("Authorization")
	authClient := client.NewAuthClient()
	response, err := authClient.Validate(&protos.ValidateRequest{
		Token: jwt,
	})
	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.JSON(http.StatusOK, utils.JSONResult("身份验证失败，服务器错误", nil, 500))
		return
	}
	if !response.Status || err != nil {
		c.Abort()
		c.JSON(http.StatusOK, utils.JSONResult("身份验证失败，请登录", nil, 500))
		return
	}
	c.Next()
}

func Cors(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
