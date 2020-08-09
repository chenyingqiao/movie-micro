package utils

import (
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

// NewCrawlerOption 初始化参数
func NewCrawlerOption() CrawlerOption {
	return CrawlerOption{
		batchNumber: 5,
		timeout:     30,
		help:        false,
		pageEnd:     "1",
	}
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
