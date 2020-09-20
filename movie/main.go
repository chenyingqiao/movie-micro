package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"movie/controller"
	"movie/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	defer utils.CloseAllGrpcClientConnect()

	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		fmt.Println("模板文件载入失败")
	}
	r.SetHTMLTemplate(t)
	loadStatic(r)
	controller.RegisterController(r, controller.NewUserCotroller())
	controller.RegisterController(r, controller.NewMovieController())
	controller.RegisterController(r, controller.NewCaptchaController())
	r.Run(":8071")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range utils.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func loadStatic(r *gin.Engine) {
	r.Static("/static", "./static/")
}
