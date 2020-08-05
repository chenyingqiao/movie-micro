package utils

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	// CrawlerStart 开始爬虫
	CrawlerStart = "start"
	// CrawlerRestart 重启爬虫
	CrawlerRestart = "restart"
	// CrawlerStop 停止爬虫
	CrawlerStop = "stop"
	// CrawlerStatus 爬虫状态
	CrawlerStatus = "status"
)

// CrawlerOption 爬虫终端
type CrawlerOption struct {
	operator        string
	goroutineNumber int
	timeout         time.Duration
	help            bool
}

//CrawlerHandle 回调函数
type CrawlerHandle func(crawlerOption *CrawlerOption) error

// NewCrawlerOption 获取一个Terminal提示
func NewCrawlerOption() CrawlerOption {
	crawlerT := CrawlerOption{}
	crawlerT.Init()
	return crawlerT
}

// Init 初始化参数
func (ct *CrawlerOption) Init() {
	if len(os.Args) <= 1 {
		ct.operator = "start"
	} else {
		ct.operator = os.Args[1]
	}
	flag.IntVar(&ct.goroutineNumber, "G", 5, "Set goroutine number")
	flag.DurationVar(&ct.timeout, "t", 180, "Set crawler timeout")
	flag.BoolVar(&ct.help, "h", true, "See help")
	flag.Parse()
}

//IsNeedShowHelp 判断是否需要显示help
func (ct *CrawlerOption) IsNeedShowHelp() bool {
	return ct.help && ct.operator == ""
}

// PrintHelper 打印help
func (ct *CrawlerOption) PrintHelper() {
	if !ct.help {
		return
	}
	fmt.Println("crawler [start|status|stop|restart] [-hGt]")
	flag.PrintDefaults()
}

//GetOptions 获取配置参数
func (ct *CrawlerOption) GetOptions() map[string]interface{} {
	return map[string]interface{}{
		"goroutineNumber": ct.goroutineNumber,
		"timeout":         ct.timeout,
	}
}

// Action 执行对应命令的处理
func (ct *CrawlerOption) Action(handles ...CrawlerHandle) error {
	for _, handle := range handles {
		err := handle(ct)
		if err != nil {
			return err
		}
	}
	return nil
}
