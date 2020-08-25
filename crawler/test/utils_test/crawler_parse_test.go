package utils_test

import (
	"crawler/utils"
	"testing"
)

func TestParseByRule(t *testing.T) {
	ruleParse := utils.NewRuleParseQuery()
	doc, err := ruleParse.GetDoc("http://www.zuidazy5.com/?m=vod-index-pg-1.html")
	if err != nil {
		t.Logf("获取文档错误")
	}
	_, errParse := ruleParse.Parse("span.xing_vb4>a#&All:Attr:href", doc)
	if errParse != nil {
		t.Errorf("失败%s", errParse)
	}
}

func TestParseByRuleInclude(t *testing.T) {
	ruleParse := utils.NewRuleParseQuery()
	doc, err := ruleParse.GetDoc("http://www.zuidazy5.com/?m=vod-detail-id-93012.html")
	if err != nil {
		t.Logf("获取文档错误")
	}
	rule := ".vodinfobox>ul>li|&include类型|span|&Text"
	_, errParse2 := ruleParse.Parse(rule, doc)
	if errParse2 != nil {
		t.Errorf("失败%s", errParse2)
	}
}
