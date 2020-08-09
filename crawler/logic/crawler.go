package logic

import (
	"context"
	"crawler/db"
	"crawler/utils"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//CrawlerLogic 爬虫逻辑
type CrawlerLogic struct {
	detailLock   sync.WaitGroup
	ruleParse    utils.RuleParse
	option       utils.CrawlerOption
	movieChan    chan db.Movie
	ctx          *context.Context
	IsFinished   bool
	maxPageReady map[string]string
}

//HandleFunc 爬虫数据通道获取handle
type HandleFunc func(movies chan db.Movie, isFinishSign chan bool)

//NewCrawlerLogic 创建CrawlerLogic
func NewCrawlerLogic(ctx *context.Context, ruleParse utils.RuleParse, co utils.CrawlerOption) *CrawlerLogic {
	crawlerLogic := &CrawlerLogic{
		detailLock:   sync.WaitGroup{},
		ruleParse:    ruleParse,
		option:       co,
		movieChan:    make(chan db.Movie, 10),
		ctx:          ctx,
		IsFinished:   false,
		maxPageReady: make(map[string]string),
	}
	return crawlerLogic
}

//GetOptionsDeadlineContext 获取通过配置设置的context
func GetOptionsDeadlineContext(crawlerOption utils.CrawlerOption) (*context.Context, context.CancelFunc) {
	timeout := 10 * time.Second
	detailContext, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	return &detailContext, cancel
}

//Handle 逻辑入口
func (cl *CrawlerLogic) Handle(page string, rule db.Rule, handle HandleFunc) {
	cl.IsFinished = false
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return
	}
	isFinish := make(chan bool)
	go handle(cl.movieChan, isFinish)
	for i := 1; i <= pageNumber; i++ {
		cl.crawler(strconv.Itoa(i), rule)
		cl.detailLock.Wait() //等上一页都采集完成
	}
	cl.IsFinished = true
	<-isFinish
}

//crawler 爬取列表页面并切创建detai页面的goroutine
func (cl *CrawlerLogic) crawler(pageNumber string, rule db.Rule) {
	url := strings.Replace(rule.ListURL, "{$page}", pageNumber, 1)
	doc, err := cl.ruleParse.GetDoc(url)
	urls, _ := cl.ruleParse.Parse(rule.ListXpath, doc)
	if err != nil {
		return
	}

	batchURL := []string{}
	batchURLList := [][]string{}
	for i, url := range urls {
		batchURL = append(batchURL, url)
		if (i+1)%cl.option.GetBatchNumber() == 0 || i+1 == len(urls) {
			batchURLList = append(batchURLList, batchURL)
			batchURL = []string{}
			continue
		}
	}
	cl.detailLock.Add(len(batchURLList))
	for _, batchURLListItem := range batchURLList {
		go cl.syncCrawlerDetailByURLs(rule, batchURLListItem)
	}
}

func (cl *CrawlerLogic) syncCrawlerDetailByURLs(rule db.Rule, batchURL []string) {
	for _, batchURLtem := range batchURL {
		movieDetail, detailErr := cl.crawlerDetail(cl.ruleParse.GetDomain()+batchURLtem, rule)
		if detailErr == nil {
			cl.movieChan <- movieDetail
		}
		select {
		case <-(*cl.ctx).Done():
			cl.detailLock.Done()
			return
		default:
		}
	}
	cl.detailLock.Done()
}

//CrawlerDetail 重新爬去详细信息
func (cl *CrawlerLogic) CrawlerDetail(detailURL string, rule db.Rule) (db.Movie, error) {
	movieDetail, err := cl.crawlerDetail(cl.ruleParse.GetDomain()+detailURL, rule)
	if err != nil {
		return db.NewMovie(), err
	}
	return movieDetail, nil
}

//crawlerDetail 获取详情页面的信息
func (cl *CrawlerLogic) crawlerDetail(detailURL string, rule db.Rule) (db.Movie, error) {
	movie := db.Movie{}
	doc, err := cl.ruleParse.GetDoc(detailURL)
	if err != nil {
		return movie, err
	}
	fieldMap := map[string]string{
		"DetailURL":          detailURL,
		"VideoMp4Sourc":      rule.VideoMp4SourceXpath,
		"Title":              rule.TitleXpath,
		"Alias":              rule.AliasXpath,
		"Director":           rule.DirectorXpath,
		"Actor":              rule.ActorXpath,
		"Location":           rule.LocationXpath,
		"Language":           rule.LanguageXpath,
		"ShowingTime":        rule.ShowingTimeXpath,
		"Long":               rule.LongXpath,
		"UpdateTime":         rule.UpdateTimeXpath,
		"Introduce":          rule.IntroduceXpath,
		"Types":              rule.TypesXpath,
		"VideoM3u8Source":    rule.VideoM3u8SourceXpath,
		"VideoZuidallSource": rule.VideoZuidallSourceXpath,
		"VideoMp4Source":     rule.VideoMp4SourceXpath,
		"ImageURL":           rule.ImageURLXpath,
	}
	data := map[string][]string{}
	for k, v := range fieldMap {
		parseValue := []string{}
		if parseValue, err = cl.ruleParse.Parse(v, doc); err != nil {
			return movie, err
		}
		data[k] = parseValue
	}
	movie.Fill(data)
	return movie, nil
}

//GetMaxPageNumber 获取最大页面
func (cl *CrawlerLogic) GetMaxPageNumber(rule db.Rule) string {
	infoURLMd5 := utils.Md5Simple(rule.PageInfoURL)
	if v, ok := cl.maxPageReady[infoURLMd5]; ok {
		return v
	}
	doc, err := cl.ruleParse.GetDoc(rule.PageInfoURL)
	if err != nil {
		return "1"
	}
	pageURL, err := cl.ruleParse.Parse(rule.PageNumberXpath, doc)
	re := regexp.MustCompile(rule.PageReg)
	if len(pageURL) <= 0 {
		return "1"
	}
	length := len(pageURL)
	lastPageURL := pageURL[length-1]
	matched := re.FindAllStringSubmatch(lastPageURL, -1)
	cl.maxPageReady[infoURLMd5] = matched[0][1]
	return matched[0][1]
}
