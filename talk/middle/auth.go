package middle

import (
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
	if !response.Status || err != nil {
		c.Abort()
		c.JSON(http.StatusOK, utils.JSONResult("身份验证失败，请登录", nil, 500))
		return
	}
}
