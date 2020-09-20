package main

import (
	"crawler/db"
	"crawler/logic"
	"crawler/utils"
	"fmt"
)

func main() {
	crawlerOption := utils.NewCrawlerOption()
	if crawlerOption.IsNeedShowHelp() {
		crawlerOption.PrintHelper()
		return
	}
	ctx, cancel := logic.GetOptionsDeadlineContext(crawlerOption)
	ruleParse := utils.NewRuleParseQuery()
	defer cancel()
	rules, err := db.GetRules()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	if len(rules) < 0 {
		fmt.Println("未找到规则")
	}
	for _, rule := range rules {
		logic := logic.NewCrawlerLogic(ctx, ruleParse, crawlerOption)
		page := "1"
		logic.Handle(page, rule, func(movies chan db.Movie, finish chan bool) {
			for !logic.IsFinished {
				movie, ok := <-movies
				if !ok {
					finish <- true
					return
				}
				err = movie.InsertWhenNotExsist()
				if err != nil {
					fmt.Println("错误")
					fmt.Printf("%+v\n", err)
					finish <- true
					return
				}
				err = movie.EsInsertWhenNotExsist()
				if err != nil {
					fmt.Printf("%+v\n", err)
				}
				fmt.Println(movie.Title + "=from:" + movie.Source)
			}
			finish <- true //完成后通过chan通知完成
		})
	}

}
