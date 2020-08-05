package cmd_test

import (
	"os"
	"testing"
)

func TestCrawler(t *testing.T) {
	os.Setenv("MONGODB_USERNAME", "root")
	os.Setenv("MONGODB_PWD", "root")
	os.Setenv("MONGODB_ADDR", "10.100.175.216:27017")
	os.Setenv("MONGODB_DB", "test")
	os.Setenv("MONGODB_TIMEOUT", "10")
	os.Setenv("MONGODB_POOLSIZE", "30")

	// crawlerOption := utils.NewCrawlerOption()
	// ctx, cancel := cmd.GetOptionsDeadlineContext(crawlerOption)
	// ruleParse := utils.NewRuleParseQuery()
	// defer cancel()
	// rules, err := db.GetRules()
	// if err != nil {
	// 	t.Errorf("%s", err)
	// }
	// if len(rules) < 0 {
	// 	return
	// }
	// movies := cmd.NewCrawlerLogic(ctx, ruleParse, crawlerOption).Handle("1", rules[0])
	// for movie := range movies {
	// 	movie.InsertWhenNotExsist()
	// }
}
