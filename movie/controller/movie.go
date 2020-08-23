package controller

import (
	"movie/errs"
	"movie/rpc/client"
	"movie/rpc/protos"
	"movie/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//MovieController 电影信息控制器
type MovieController struct {
	movieRPCClient client.MovieClient
}

//RegisterMovieController 注册MovieController
func RegisterMovieController(engine *gin.Engine) {
	controller := MovieController{
		movieRPCClient: client.NewMovieClient(),
	}
	controller.Register(engine)
}

//Register 注册控制器中的方法
func (m *MovieController) Register(engine *gin.Engine) {
	engine.GET("/", m.hello)
	engine.GET("/detail/:hash", m.detail)
	engine.GET("/list/:id", m.list)
	engine.GET("/list", m.list)
}

func (m *MovieController) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
	return
}

//detail
func (m *MovieController) detail(c *gin.Context) {
	hash := c.Params.ByName("hash")
	if hash == "" {
		c.JSON(http.StatusOK, utils.JSONResult("error", nil, 200))
	}
	request := &protos.MovieRequest{
		Hash: hash,
	}
	response, err := m.movieRPCClient.Detail(request)
	if errs.ValidataGrpcErrorError(err) {
		c.JSON(http.StatusOK, utils.JSONResult(err.Error(), nil, 200))
		return
	}
	c.JSON(http.StatusOK, utils.JSONResult("success", response, 200))
	return
}

func (m *MovieController) list(c *gin.Context) {
	objID := c.Params.ByName("id")
	request := &protos.MovieRequest{
		ObjId: objID,
	}

	response, err := m.movieRPCClient.List(request)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "/tmpl/list.html", response)
	}
	c.HTML(http.StatusOK, "/tmpl/list.html", response)
	return
}
