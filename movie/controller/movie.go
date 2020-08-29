package controller

import (
	"movie/errs"
	"movie/rpc/client"
	"movie/rpc/protos"
	"movie/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//MovieController 电影信息控制器
type MovieController struct {
	movieRPCClient client.MovieClient
}

//NewMovieController 注册MovieController
func NewMovieController() *MovieController {
	controller := MovieController{
		movieRPCClient: client.NewMovieClient(),
	}
	return &controller
}

//Register 注册控制器中的方法
func (m *MovieController) Register(engin *gin.Engine) {
	engin.GET("/", m.index)
	engin.GET("/detail/:hash", m.detail)
	engin.GET("/list/:id", m.list)
	engin.GET("/list", m.list)
}

func (m *MovieController) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
	return
}

func (m *MovieController) index(c *gin.Context) {
	c.HTML(http.StatusOK, "/tmpl/index.html", nil)
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
	m3u8Data := []map[string]string{}
	for _, m3u8URl := range response.VideoM3U8Source {
		if !strings.Contains(m3u8URl, "$") {
			continue
		}
		item := strings.Split(m3u8URl, "$")
		m3u8Data = append(m3u8Data, map[string]string{
			"title": item[0],
			"url":   item[1],
		})
	}

	zuidaIsM3u8 := false
	zuidallData := []map[string]string{}
	for _, ZuidallURL := range response.VideoZuidallSource {
		if !strings.Contains(ZuidallURL, "$") {
			continue
		}
		item := strings.Split(ZuidallURL, "$")
		if strings.Contains(item[1], "m3u8") {
			zuidaIsM3u8 = true
		}
		zuidallData = append(zuidallData, map[string]string{
			"title": item[0],
			"url":   item[1],
		})
	}

	if zuidaIsM3u8 {
		m3u8Data = zuidallData
	}

	htmlData := map[string]interface{}{
		"data":    response,
		"m3u3":    m3u8Data,
		"zuidall": zuidallData,
		"default": m3u8Data[0],
	}
	c.HTML(http.StatusOK, "/tmpl/detail.html", htmlData)
	return
}

func (m *MovieController) list(c *gin.Context) {
	objID := c.Params.ByName("id")
	types := c.Query("type")
	request := &protos.MovieRequest{
		ObjId: objID,
		Type:  types,
	}

	response, err := m.movieRPCClient.List(request)
	if err != nil || len(response) == 0 {
		c.HTML(http.StatusInternalServerError, "/tmpl/list.html", map[string]interface{}{
			"data":      response,
			"last_hash": nil,
		})
		return
	}
	c.HTML(http.StatusOK, "/tmpl/list.html", map[string]interface{}{
		"data":      response,
		"last_hash": response[len(response)-1],
	})
	return
}
