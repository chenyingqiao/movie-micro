package utils

import (
	netURL "net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

//RuleParse 规则解析
type RuleParse interface {
	Parse(rule string, doc *goquery.Document) ([]string, error)
	GetDoc(url string) (*goquery.Document, error)
	GetDomain() string
}

//RuleParseQuery 页面解析
type RuleParseQuery struct {
	domain string
}

//NewRuleParseQuery 返回一个RuleParse对象
func NewRuleParseQuery() *RuleParseQuery {
	return &RuleParseQuery{}
}

//Parse 执行parse
func (r *RuleParseQuery) Parse(rule string, doc *goquery.Document) ([]string, error) {
	ruleArr := strings.Split(rule, "|")
	var selection *goquery.Selection
	result := []string{}
	var docCp *goquery.Document
	for i, v := range ruleArr {
		if strings.HasPrefix(v, "&") && selection != nil {
			v = strings.Replace(v, "&", "", 1)
			if newDocNodePoint := r.includeRule(v, selection); newDocNodePoint != -1 && len(selection.Nodes) > 0 {
				docCp = goquery.NewDocumentFromNode(selection.Nodes[newDocNodePoint])
				continue
			}

			result = r.final(v, selection)

		} else {
			if docCp != nil {
				selection = docCp.Find(v)
			} else {
				selection = doc.Find(v)
			}
			isNode := len(selection.Nodes) > 0
			if i < len(ruleArr)-1 && !isNode {
				return []string{""}, nil
			}
		}
	}
	return result, nil
}

//includeRule 从上一个selection中找到内容中包含对应字符的节点
func (r *RuleParseQuery) includeRule(parteRule string, selection *goquery.Selection) int {
	nodePosition := -1
	if strings.HasPrefix(parteRule, "include") {
		includeString := strings.Replace(parteRule, "include", "", 1)
		selection.Each(func(i int, s *goquery.Selection) {
			content := s.Text()
			if strings.Contains(content, includeString) {
				nodePosition = i
				return
			}
		})
		return nodePosition
	}
	return nodePosition
}

func (r *RuleParseQuery) final(parteRule string, selection *goquery.Selection) []string {
	if strings.HasPrefix(parteRule, "All:") {
		return r.finalRuleAtAll(
			strings.Replace(parteRule, "All:", "", 1),
			selection,
		)
	} else {
		return []string{
			r.finalRule(parteRule, selection),
		}
	}
}

func (r *RuleParseQuery) finalRuleAtAll(parteRule string, selection *goquery.Selection) []string {
	result := []string{}
	if parteRule == "Text" {
		selection.Each(func(i int, item *goquery.Selection) {
			result = append(result, item.Text())
		})
	} else if strings.HasPrefix(parteRule, "Attr:") {
		selection.Each(func(i int, item *goquery.Selection) {
			attrStr, isExsist := item.Attr(strings.Replace(parteRule, "Attr:", "", 1))
			if isExsist {
				result = append(result, attrStr)
			}
		})
	}
	return result
}

//finalTextRule 获取最终节点中信息
func (r *RuleParseQuery) finalRule(parteRule string, selection *goquery.Selection) string {
	if parteRule == "Text" {
		return selection.Text()
	} else if strings.HasPrefix(parteRule, "Attr:") {
		attrStr, isExsist := selection.Attr(strings.Replace(parteRule, "Attr:", "", 1))
		if !isExsist {
			return ""
		} else {
			return attrStr
		}
	}
	return ""
}

//GetDoc 获取页面内容
func (r *RuleParseQuery) GetDoc(url string) (*goquery.Document, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	urlParseInfo, err := netURL.Parse(url)
	if err != nil {
		return nil, errors.Wrap(err, "传入的url无法解析")
	}
	r.domain = "http://" + urlParseInfo.Host
	return doc, err
}

//GetDomain 获取domain
func (r *RuleParseQuery) GetDomain() string {
	return r.domain
}
