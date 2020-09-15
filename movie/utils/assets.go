package utils

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets1ae98f10d3ac183c7640a24264c7f219a35c8e17 = "<!doctype html>\n<html lang=\"en\">\n  <script>\n    var _hmt = _hmt || [];\n    (function() {\n      var hm = document.createElement(\"script\");\n      hm.src = \"https://hm.baidu.com/hm.js?e7056a58ebcf4780dadc9ce4b8cd7fd6\";\n      var s = document.getElementsByTagName(\"script\")[0]; \n      s.parentNode.insertBefore(hm, s);\n    })();\n  </script>\n  <head>\n    <link rel=\"shortcut icon\" type=\"image/x-icon\" href=\"/static/favicon.ico\" />\n    <title>乐扣电影</title>\n    <!-- Required meta tags -->\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n    <meta name=\"description\" content=\"乐扣电影\" />\n    <meta name=\"keywords\" content=\"乐扣电影 电影网 电影下载 电影在线 在线观看 美剧 釜山行 半岛\" />\n\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\" integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">\n  </head>\n\n  <link rel=\"stylesheet\" href=\"/static/css/index.css\">\n  <body>\n\n    <div id=\"side\">\n        <div class=\"list-group\">\n        </div>\n        <div class=\"list-group menu\">\n            <a href=\"#\" class=\"list-group-item\" data=\"\">全部</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"战争片\">战争片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"音乐片\">音乐片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"国产剧\">国产剧</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"韩国剧\">韩国剧</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"记录片\">记录片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"动漫片\">动漫片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"日本剧\">日本剧</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"恐怖片\">恐怖片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"台湾剧\">台湾剧</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"剧情片\">剧情片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"动作片\">动作片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"喜剧片\">喜剧片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"综艺片\">综艺片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"欧美剧\">欧美剧</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"海外剧\">海外剧</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"科幻片\">科幻片</a>\n            <a href=\"#\" class=\"list-group-item\" data=\"香港剧\">香港剧</a>\n        </div>\n    </div>\n\n    <div class=\"container-full\">\n        <input type=\"hidden\" id=\"hid_isload\" value=\"0\">\n        <input type=\"hidden\" id=\"last_id\" value=\"0\">\n        <div class=\"row\">\n            <input type=\"text\"\n                class=\"form-control search_input\" name=\"\" id=\"search\" aria-describedby=\"helpId\" placeholder=\"输入要搜索的影片，回车搜索\">\n        </div>\n        <div class=\"row\">\n            <div class=\"col-xs-1-12 movie-content\">\n            </div>\n        </div>\n    </div>\n    <a href=\"/list/\"></a>\n      \n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    <script src=\"https://libs.baidu.com/jquery/2.1.1/jquery.min.js\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js\" integrity=\"sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1\" crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js\" integrity=\"sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM\" crossorigin=\"anonymous\"></script>\n  </body>\n  <script src=\"/static/js/index.js\"></script>\n</html>"
var _Assets61a6c052b04233124da80ed6652a1cb70f38c9b8 = "<!doctype html>\n<html lang=\"en\">\n  <script>\n    var _hmt = _hmt || [];\n    (function() {\n      var hm = document.createElement(\"script\");\n      hm.src = \"https://hm.baidu.com/hm.js?e7056a58ebcf4780dadc9ce4b8cd7fd6\";\n      var s = document.getElementsByTagName(\"script\")[0]; \n      s.parentNode.insertBefore(hm, s);\n    })();\n  </script>\n  <head>\n    <link rel=\"shortcut icon\" type=\"image/x-icon\" href=\"/static/favicon.ico\" />\n    <title>{{.data.Title}}</title>\n    <!-- Required meta tags -->\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n    <meta name=\"description\" content=\"乐扣电影\" />\n    <meta name=\"keywords\" content=\"乐扣电影 电影网 电影下载 电影在线 在线观看 美剧 釜山行 半岛\" />\n\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\" integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">\n    <link href=\"/static/video-js/video-js.css\" rel=\"stylesheet\">\n    <script src=\"/static/video-js/video.js\"></script>\n    <script src=\"/static/video-js/videojs-contrib-hls.min.js\"></script>\n    <script>\n        videojs.options.flash.swf = \"video-js/video-js.swf\"\n        var movieHash = \"{{.data.ID}}\"\n    </script>\n    <link rel=\"stylesheet\" href=\"/static/css/detail.css\">\n  </head>\n  <body>\n    <div id=\"video-panel\" style=\"float:left\">\n        <video id=\"live-stream\" class=\"video-js vjs-default-skin vjs-big-play-centered\"\n        controls preload=\"auto\"\n        data-setup='{\"example_option\":true}'>\n        <source src=\"{{.default.url}}\" type=\"application/x-mpegURL\">\n        </video>\n    </div>\n    <ul class=\"nav nav-tabs\" id=\"my-tab\">\n        <li class=\"nav-item\">\n            <a href=\"#choose\" class=\"nav-link active\">选集</a>\n        </li>\n        <li class=\"nav-item\">\n            <a href=\"#talk\" class=\"nav-link \">聊天</a>\n        </li>\n        <li class=\"nav-item\">\n          <a href=\"/\" class=\"nav-link\" id=\"redrect_home\" style=\"color:red\">🏠︎<span style=\"font-size:0.3rem\">首页</span></a>\n        </li>\n    </ul>\n\n    <div class=\"tab-content\" id=\"my-tab-p\"> \n        <div class=\"tab-pane\" id=\"talk\">\n            <span id=\"lastID\" style=\"display:none;\" data=\"\"></span>\n            <ul class=\"list-group list-group-flush\" id=\"comment_p\">\n            </ul>\n            <input type=\"text\" id=\"message_text\" placeholder=\"输入并回车\">\n        </div>\n        <div class=\"tab-pane active\" id=\"choose\">\n            <div>\n                <div style=\"display: flex;justify-content: center;align-items: center;\">\n                    <img src=\"{{.data.ImageURL}}\" alt=\"\" width=\"290px\">\n                </div>\n                <div>\n                    <div>标题： {{.data.Title}}</div>\n                    <div>产地： {{.data.Location}}</div>\n                    <div>类型： {{.data.Types}}</div>\n                    <div>演员： {{.data.Actor}}</div>\n                    <div>更新： {{.data.UpdateTime}}</div>\n                    <div>简介:  {{.data.Introduce}}</div>\n                </div>\n            </div>\n            <div class=\"list-group\">\n                {{range .m3u3}}\n                <button type=\"button\" class=\"list-group-item list-group-item-action choose-button\" src=\"{{.url}}\">{{.title}}</button>\n                {{end}}\n            </div>\n        </div>\n    </div>\n\n    <!--登录模态框-->\n      <div class=\"modal fade\" id=\"login-mode\" tabindex=\"-1\" role=\"dialog\" aria-labelledby=\"login-mode\" aria-hidden=\"true\">\n        <div class=\"modal-dialog\">\n          <div class=\"modal-content\">\n            <div class=\"modal-header\">\n              <h5 class=\"modal-title\" id=\"login-mode\">注册</h5>\n              <button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-label=\"Close\">\n                <span aria-hidden=\"true\">&times;</span>\n              </button>\n            </div>\n            <div class=\"modal-body\">\n                <form class=\"form-inline center\">\n                    <div class=\"form-group \">\n                        <label for=\"\">用户名:</label>\n                        <input type=\"text\" name=\"username\" id=\"username\" class=\"form-control\" placeholder=\"\" aria-describedby=\"helpId\">\n                    </div>\n                    <div class=\"form-group \">\n                        <label for=\"\">密码:</label>\n                        <input type=\"password\" name=\"password\" id=\"password\" class=\"form-control\" placeholder=\"\" aria-describedby=\"helpId\">\n                    </div>\n                    <div class=\"form-group \">\n                      <label for=\"\">验证码:</label>\n                      <input type=\"text\" name=\"answer\" id=\"login_answer\" class=\"form-control\" placeholder=\"\" aria-describedby=\"helpId\">\n                      <input type=\"hidden\" name=\"cap_id\" id=\"login_captcha_id\">\n                  </div>\n                  <div class=\"form-group \">\n                    <label for=\"\"></label>\n                    <audio controls=\"controls\" id=\"audio_captcha_login\" src=\"\"/>\n                    </div>\n                </form>\n            </div>\n            <div class=\"modal-footer\">\n              <button type=\"button\" class=\"btn btn-secondary\" data-dismiss=\"modal\">Close</button>\n              <button type=\"button\" id=\"login_action\" class=\"btn btn-primary\">登录</button>\n              <button type=\"button\" id=\"register_tigger\" class=\"btn btn-primary\">注册</button>\n            </div>\n          </div>\n        </div>\n      </div>\n\n    <!--注册模态框-->\n    <div class=\"modal fade\" id=\"register-mode\" tabindex=\"-1\" role=\"dialog\" aria-labelledby=\"register-mode\" aria-hidden=\"true\">\n        <div class=\"modal-dialog\">\n          <div class=\"modal-content\">\n            <div class=\"modal-header\">\n              <h5 class=\"modal-title\" id=\"register-mode\">注册</h5>\n              <button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-label=\"Close\">\n                <span aria-hidden=\"true\">&times;</span>\n              </button>\n            </div>\n            <div class=\"modal-body\">\n                <form class=\"form-inline center\">\n                    <div class=\"form-group \">\n                        <label for=\"\">用户名:</label>\n                        <input type=\"text\" name=\"username\" id=\"reg_username\" class=\"form-control\" placeholder=\"\" aria-describedby=\"helpId\">\n                    </div>\n                    <div class=\"form-group \">\n                        <label for=\"\">密码:</label>\n                        <input type=\"password\" name=\"password\" id=\"reg_password\" class=\"form-control\" placeholder=\"\" aria-describedby=\"helpId\">\n                    </div>\n                    <div class=\"form-group \">\n                        <label for=\"\">确认密码:</label>\n                        <input type=\"password\" name=\"password_re\" id=\"reg_password_re\" class=\"form-control\" placeholder=\"\" aria-describedby=\"helpId\">\n                    </div>\n                    <div class=\"form-group \">\n                        <label for=\"\">验证码:</label>\n                        <input type=\"text\" name=\"answer\" id=\"register_answer\" class=\"form-control\" placeholder=\"\" aria-describedby=\"helpId\">\n                        <input type=\"hidden\" name=\"cap_id\" id=\"register_captcha_id\">\n                        <br />\n                    </div>\n                  <div class=\"form-group \">\n                    <label for=\"\"></label>\n                    <audio controls=\"controls\" id=\"audio_captcha_register\" src=\"\"/>\n                  </div>\n                </form>\n            </div>\n            <div class=\"modal-footer\">\n              <button type=\"button\" class=\"btn btn-secondary\" data-dismiss=\"modal\">Close</button>\n              <button type=\"button\" id=\"register_login\" class=\"btn btn-primary\">注册</button>\n            </div>\n          </div>\n        </div>\n    </div>\n      \n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    <script src=\"http://libs.baidu.com/jquery/2.0.0/jquery.min.js\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js\" integrity=\"sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1\" crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js\" integrity=\"sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM\" crossorigin=\"anonymous\"></script>\n    <script src=\"/static/js/detail.js?_v=1.0.1\"></script>\n      \n  </body>\n</html>"
var _Assetse258cf2577c5233e6e430cb062e105ee4411ae82 = "{{range .data}}\n<div class=\"movie\" href=\"/detail/{{.ID}}\">\n    <img src=\"{{.ImageURL}}\" alt=\"\" onerror=\"this.src='http://img.chenyingqiao.top/blog/20200824202651.png'\">\n    <div class=\"cover\" oid=\"{{.ID}}\"><span>{{.Title}}</span> </div>\n</div>\n{{else}}\nnothing\n{{end}}\n\n{{if .last_hash}}\n<a href=\"/list/{{.last_hash.ID}}\"></a>\n{{end}}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"tmpl"}, "/tmpl": []string{"detail.html", "list.html", "index.html"}}, map[string]*assets.File{
	"/tmpl/detail.html": &assets.File{
		Path:     "/tmpl/detail.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1600172383, 1600172383518745199),
		Data:     []byte(_Assets61a6c052b04233124da80ed6652a1cb70f38c9b8),
	}, "/tmpl/list.html": &assets.File{
		Path:     "/tmpl/list.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1600169691, 1600169691098858962),
		Data:     []byte(_Assetse258cf2577c5233e6e430cb062e105ee4411ae82),
	}, "/tmpl/index.html": &assets.File{
		Path:     "/tmpl/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1600068008, 1600068008432449294),
		Data:     []byte(_Assets1ae98f10d3ac183c7640a24264c7f219a35c8e17),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1600170692, 1600170692647704148),
		Data:     nil,
	}, "/tmpl": &assets.File{
		Path:     "/tmpl",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1598519129, 1598519129908119504),
		Data:     nil,
	}}, "")
