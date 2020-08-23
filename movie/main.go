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
	defer utils.CloseGrpcClientConnect()

	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		fmt.Println("模板文件载入失败")
	}
	r.SetHTMLTemplate(t)
	controller.RegisterMovieController(r)
	r.Run(":8081")
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
