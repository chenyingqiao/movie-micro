package utils

import (
	"flag"
	"fmt"
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
	batchNumber int
	timeout     time.Duration
	help        bool
	pageEnd     string
}

//CrawlerHandle 回调函数
type CrawlerHandle func(crawlerOption *CrawlerOption) error

//NewCrawlerOptionNotCmd 非名命令行环境下进行初始化
func NewCrawlerOptionNotCmd() CrawlerOption {
	return CrawlerOption{
		batchNumber: 5,
		timeout:     30,
		help:        false,
		pageEnd:     "1",
	}
}

// NewCrawlerOption 获取一个Terminal提示
func NewCrawlerOption() CrawlerOption {
	crawlerT := CrawlerOption{}
	crawlerT.Init()
	return crawlerT
}

// Init 初始化参数
func (ct *CrawlerOption) Init() {
	flag.IntVar(&ct.batchNumber, "G", 5, "Set per batch crawler number")
	flag.BoolVar(&ct.help, "h", false, "See help")
	flag.StringVar(&ct.pageEnd, "p", "", "End of page")
	flag.Parse()
}

//IsNeedShowHelp 判断是否需要显示help
func (ct *CrawlerOption) IsNeedShowHelp() bool {
	return ct.help
}

// PrintHelper 打印help
func (ct *CrawlerOption) PrintHelper() {
	if !ct.help {
		return
	}
	fmt.Println("crawler [start|status|stop|restart] [-hGtp]")
	flag.PrintDefaults()
}

//GetPageNumber 获取设置的最小页面
func (ct *CrawlerOption) GetPageNumber() string {
	return ct.pageEnd
}

//GetBatchNumber 设置批量数量
func (ct *CrawlerOption) GetBatchNumber() int {
	return ct.batchNumber
}

//HasPageNumber 是否有pagenumber设置项目
func (ct *CrawlerOption) HasPageNumber() bool {
	return ct.pageEnd != ""
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
