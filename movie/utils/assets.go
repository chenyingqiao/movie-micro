package utils

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetse258cf2577c5233e6e430cb062e105ee4411ae82 = "<div class=\"movie\">\n    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n    <div class=\"cover\"><span>北部丛林</span> </div>\n</div>\n\n{{range .}}\n<div class=\"movie\">\n    <img src=\"{{.ImageURL}}\" alt=\"\" href=\"#\">\n    <div class=\"cover\"><span>{{.Title}}</span> </div>\n</div>\n{{else}}\nnothing\n{{end}}"
var _Assets1ae98f10d3ac183c7640a24264c7f219a35c8e17 = "<!doctype html>\n<html lang=\"en\">\n  <head>\n    <title>Title</title>\n    <!-- Required meta tags -->\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\" integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">\n  </head>\n\n  <style>\n    .movie{\n        width:250px;\n        height:370px;\n        float: left;\n        margin: 5px;\n    }\n    .movie img{\n        width:250px;\n        height:370px;\n    }\n    .movie .cover {\n        width:250px;\n        height:370px;\n        line-height: 20px;\n        background-color: grey;\n        translate: 80%;\n        position:relative;\n        top: -370px;\n        background-color: rgba(255,255,255,0.5);\n        text-align: center;\n        line-height: 370px;\n        font-weight: bold;\n        user-select:none;\n        display: none;\n    }\n\n    .mao {\n        background: inherit;\n        -webkit-filter: blur(5px);\n        -moz-filter: blur(5px);\n        -ms-filter: blur(5px);\n        -o-filter: blur(5px);\n        filter: blur(5px);\n        filter: progid:DXImageTransform.Microsoft.Blur(PixelRadius=9, MakeShadow=false);\n    }\n\n    .movie-content{\n        display:flex;\n        flex-wrap:wrap;\n        align-content:flex-start;\n        align-items: center;\n        justify-content:center;\n    }\n    .search_input{\n        width:500px;\n        margin: 0px auto;\n        margin-top: 30px;\n        margin-bottom: 10px;\n    }\n  </style>\n  <body>\n\n  \n\n    <div class=\"container\">\n        <input type=\"hidden\" id=\"hid_isload\" value=\"0\">\n        <input type=\"hidden\" id=\"last_id\" value=\"0\">\n        <div class=\"row\">\n            <input type=\"text\"\n                class=\"form-control search_input\" name=\"\" id=\"\" aria-describedby=\"helpId\" placeholder=\"输入要搜索的影片\">\n        </div>\n        <div class=\"row\">\n            <div class=\"col-xs-1-12 movie-content\">\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\" href=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <div class=\"movie\">\n                    <img src=\"https://tu.tianzuida.com/pic/upload/vod/2020-08-15/202008151597474441.jpg\" alt=\"\">\n                    <div class=\"cover\"><span>北部丛林</span> </div>\n                </div>\n                <a href=\"#\" style=\"display:none;\">next-link</a>\n            </div>\n        </div>\n    </div>\n      \n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    <script src=\"https://code.jquery.com/jquery-3.3.1.slim.min.js\" integrity=\"sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo\" crossorigin=\"anonymous\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js\" integrity=\"sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1\" crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js\" integrity=\"sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM\" crossorigin=\"anonymous\"></script>\n  </body>\n  <script>\n      $(function(){\n            $(document).on(\"mouseover\",\".movie\",function(env){\n                $(this).find(\".cover\").show(1500)\n            })\n            $(document).on(\"mouseout\",\".movie\",function(env){\n                $(this).find(\".cover\").hide(1500)\n            })\n      });\n      var jlheight = 50;             //距下边界长度/单位px\n      var totalheight = 0;\n      var is_load = 0;\n      var pages = 2;\n      var n_i = 0;\n      $(window).scroll(function () {\n          var srollPos = $(window).scrollTop();    //滚动条距顶部距离(页面超出窗口的高度)\n          totalheight = parseFloat($(window).height()) + parseFloat(srollPos);\n          is_load = $(\"#hid_isload\").val();\n          if (is_load == 0) {\n              if (($(document).height() - jlheight) <= totalheight) {\n                  $(\"#hid_isload\").attr(\"value\", \"1\");//滚动的时候设置成1：不执行滚动\n                  console.log(\"加载....\");\n              }\n          }\n      });\n  </script>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"tmpl"}, "/tmpl": []string{"list.html", "index.html"}}, map[string]*assets.File{
	"/tmpl": &assets.File{
		Path:     "/tmpl",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1598171165, 1598171165990984649),
		Data:     nil,
	}, "/tmpl/list.html": &assets.File{
		Path:     "/tmpl/list.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1598175015, 1598175015998771554),
		Data:     []byte(_Assetse258cf2577c5233e6e430cb062e105ee4411ae82),
	}, "/tmpl/index.html": &assets.File{
		Path:     "/tmpl/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1598174820, 1598174820711497319),
		Data:     []byte(_Assets1ae98f10d3ac183c7640a24264c7f219a35c8e17),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1598175384, 1598175384855232547),
		Data:     nil,
	}}, "")