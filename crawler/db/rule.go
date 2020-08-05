package db

// Rule 电影信息抓取的规则
type Rule struct {
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
	PageNumberXpath         string
}

//GetRules 获取规则列表
func GetRules() ([]Rule, error) {
	rules := []Rule{
		Rule{
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
			PageNumberXpath:         "div.pages|&include尾页|Attr:href",
		},
	}
	return rules, nil
}

//GetList rule
// func (rule *Rule) GetList() ([]Movie, error) {

// }
