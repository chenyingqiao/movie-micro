package db

import "errors"

// Rule 电影信息抓取的规则
type Rule struct {
	Source                  string
	ListURL                 string
	ListXpath               string
	TitleXpath              string
	AliasXpath              string
	Scroe                   string
	DirectorXpath           string
	ActorXpath              string
	TypesXpath              string
	LocationXpath           string
	LanguageXpath           string
	ShowingTimeXpath        string
	LongXpath               string
	UpdateTimeXpath         string
	IntroduceXpath          string
	VideoM3u8SourceXpath    string
	VideoZuidallSourceXpath string
	VideoMp4SourceXpath     string
	ImageURLXpath           string
	PageNumberXpath         string
	PageReg                 string
	PageInfoURL             string
}

//NewRule 获取新的rule
func NewRule() Rule {
	return Rule{}
}

//GetRules 获取规则列表
func GetRules() ([]Rule, error) {
	rules := []Rule{
		{
			Source:                  "www.zuidazy5.com",
			ListURL:                 "http://www.zuidazy5.com/?m=vod-index-pg-{$page}.html",
			ListXpath:               "span.xing_vb4>a|&All:Attr:href",
			TitleXpath:              ".vodh>h2|&Text",
			AliasXpath:              ".vodinfobox>ul>li|&include别名|span|&Text",
			Scroe:                   ".vodh>label|&Text",
			DirectorXpath:           ".vodinfobox>ul>li|&include导演|span|&Text",
			ActorXpath:              ".vodinfobox>ul>li|&include主演|span|&Text",
			TypesXpath:              ".vodinfobox>ul>li|&include类型|span|&Text",
			LocationXpath:           ".vodinfobox>ul>li|&include地区|span|&Text",
			LanguageXpath:           ".vodinfobox>ul>li|&include语言|span|&Text",
			ShowingTimeXpath:        ".vodinfobox>ul>li|&include上映|span|&Text",
			LongXpath:               ".vodinfobox>ul>li|&include片长|span|&Text",
			UpdateTimeXpath:         ".vodinfobox>ul>li|&include更新|span|&Text",
			IntroduceXpath:          "div.jjText>span.more|&Text",
			VideoM3u8SourceXpath:    "#play_1>ul>li|&All:Text",
			VideoZuidallSourceXpath: "#play_2>ul>li|&All:Text",
			VideoMp4SourceXpath:     "#play_3>ul>li|&All:Text",
			ImageURLXpath:           "div.vodImg>img.lazy|&Attr:src",
			PageNumberXpath:         "a.pagelink_a|&All:Attr:href",
			PageReg:                 "vod-index-pg-(\\d+).html",
			PageInfoURL:             "http://www.zuidazy5.com/",
		},
	}
	return rules, nil
}

//GetBySource 通过source获取rule
func (rule *Rule) GetBySource(source string) (Rule, error) {
	rules, err := GetRules()
	if err != nil {
		return Rule{}, err
	}
	for _, v := range rules {
		if v.Source == source {
			return v, err
		}
	}
	return Rule{}, errors.New("no found rule")
}

//GetList rule
// func (rule *Rule) GetList() ([]Movie, error) {

// }
