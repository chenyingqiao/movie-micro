package controller

import (
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"talk/logic"
	"talk/middle"
	"talk/room"
	"talk/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoomController struct{}

func NewRoomController() *RoomController {
	return &RoomController{}
}

//Register Register
func (rc *RoomController) Register(engin *gin.Engine) {
	engin.SetHTMLTemplate(rc.testTmpl())
	engin.GET("/", rc.hello)
	engin.GET("/test/:roomid", rc.test)
	engin.POST("/room/:roomid", middle.AuthMiddle, rc.message)
	engin.GET("/stream/:roomid", rc.estream)
	engin.GET("/talk/:roomid/:id", rc.page)
	engin.GET("/talk/:roomid", rc.page)
}

func (m *RoomController) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
	return
}

func (rc RoomController) test(c *gin.Context) {

	roomid := c.Param("roomid")
	userid := fmt.Sprint(rand.Int31())
	c.HTML(http.StatusOK, "chat_room", gin.H{
		"roomid": roomid,
		"userid": userid,
	})
}

//Message 判断是否加入房间并且发送消息
func (rc *RoomController) message(c *gin.Context) {
	jwt := c.GetHeader("Authorization")
	roomID := c.Param("roomid")
	message := c.PostForm("message")
	authLogic := logic.NewAuthLogic()
	userInfo, err := authLogic.GetUserInfo(jwt)
	if err != nil {
		logrus.WithField("err", err).Info("用户信息获取失败")
		c.JSON(http.StatusOK, utils.JSONResult("用户信息获取失败", nil, 500))
	}
	username := userInfo.Info.Username

	chatGuardian := room.NewChatGuardianSingleton()
	mesgInfo := room.NewMessage(username, roomID, message)
	chatGuardian.Talk(mesgInfo)

	//添加发言数据到数据库中
	talkLogic := logic.NewTalkLogic()
	talkLogic.Submit(roomID, username, message)
	c.JSON(http.StatusOK, utils.JSONResult("success", nil, 200))
}

//EventStream 推送
func (rc *RoomController) estream(c *gin.Context) {
	jwt := c.Query("token")
	roomID := c.Param("roomid")
	authLogic := logic.NewAuthLogic()
	userInfo, err := authLogic.GetUserInfo(jwt)
	if err != nil {
		logrus.WithField("err", err).Info("用户信息获取失败")
		c.JSON(http.StatusOK, utils.JSONResult("用户信息获取失败", nil, 500))
	}
	username := userInfo.Info.Username

	var people room.People
	people = room.NewTalkPeople(username)
	op := room.NewRoomOperator(roomID, &people)
	chatGuardian := room.NewChatGuardianSingleton()
	if !chatGuardian.IsInRoom(roomID, &people) {
		chatGuardian.Join(op)
	}
	defer chatGuardian.Out(op)

	clientGone := c.Writer.CloseNotify()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			chatGuardian.Out(op)
			return false
		case message := <-*people.GetChan():
			messageInfo := message.(room.Message)
			if messageInfo.IsHeartbeat() {
				c.SSEvent("message", gin.H{
					"heartbeat": true,
				})
				return true
			}
			c.SSEvent("message", utils.JSONResult("success", gin.H{
				"message": fmt.Sprintf("<span>%s</span>: %s", messageInfo.GetUsername(), messageInfo.GetMessage()),
			}, 200))
			return true
		}
	})

}

func (rc *RoomController) page(c *gin.Context) {
	lastID := c.Params.ByName("id")
	roomID := c.Params.ByName("roomid")
	talkLogic := logic.NewTalkLogic()
	talks := talkLogic.LoadPage(lastID, roomID)
	c.JSON(http.StatusOK, utils.JSONResultArr("success", talks, 200))
}

func (rc *RoomController) testTmpl() *template.Template {
	var html = template.Must(template.New("chat_room").Parse(`
	<html> 
	<head> 
		<title>{{.roomid}}</title>
		<link rel="stylesheet" type="text/css" href="http://meyerweb.com/eric/tools/css/reset/reset.css">
		<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.7/jquery.js"></script> 
		<script src="http://malsup.github.com/jquery.form.js"></script> 
		<script> 
			$('#message_form').focus();
			$(document).ready(function() { 
				// bind 'myForm' and provide a simple callback function 
				$('#myForm').ajaxForm(function() {
					$('#message_form').val('');
					$('#message_form').focus();
				});

				if (!!window.EventSource) {
					var source = new EventSource('/stream/{{.roomid}}/{{.userid}}');
					source.addEventListener('message', function(e) {
						$('#messages').append(e.data + "</br>");
						$('html, body').animate({scrollTop:$(document).height()}, 'slow');
					}, false);
				} else {
					alert("NOT SUPPORTED");
				}
			});
		</script> 
		</head>
		<body>
		<h1>Welcome to {{.roomid}} room</h1>
		<div id="messages"></div>
		<form id="myForm" action="/room/{{.roomid}}" method="post"> 
		User: <input id="user_form" name="username" value="{{.userid}}">
		Message: <input id="message_form" name="message">
		<input type="submit" value="Submit"> 
		</form>
	</body>
	</html>
	`))
	return html
}
