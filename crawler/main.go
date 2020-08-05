package main

import (
	"crawler/cmd"
	"crawler/db"
	"crawler/utils"
	"fmt"
)

func main() {
	crawlerOption := utils.NewCrawlerOption()
	if crawlerOption.IsNeedShowHelp() {
		crawlerOption.PrintHelper()
		return
	}
	ctx, cancel := cmd.GetOptionsDeadlineContext(crawlerOption)
	ruleParse := utils.NewRuleParseQuery()
	defer cancel()
	rules, err := db.GetRules()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	if len(rules) < 0 {
		fmt.Println("未找到规则")
	}
	logic := cmd.NewCrawlerLogic(ctx, ruleParse, crawlerOption)
	page := logic.GetMaxPageNumber(rules[0])
	fmt.Println(page)
	logic.Handle("20", rules[0], func(movies chan db.Movie, finish chan bool) {
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
			fmt.Println(movie.Title)
		}
		finish <- true //完成后通过chan通知完成
	})

}
