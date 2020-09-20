package db_test

import (
	"crawler/db"
	"crawler/rpc/protos"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestEsInsert(t *testing.T) {
	a := assert.New(t)
	os.Setenv("ES_ADDRESS", "elasticsearch:9200")
	os.Setenv("ES_TIMEOUT", "10")

	movie := db.NewMovie()
	movie.ID = primitive.NewObjectID()
	movie.Source = "www.zuidazy5.com"
	movie.DetailURL = "http://www.zuidazy5.com/?m=vod-detail-id-93756.html"
	movie.Title = "黑袍纠察队第二季"
	movie.Alias = "男孩帮/黑衣小队/七痞郎 / 英雄克星"
	movie.Director = "菲利浦·斯格里西亚,塞巴斯蒂安·席尔瓦"
	movie.Actor = "卡尔·厄本,杰克·奎德,安东尼·斯塔尔,艾琳·莫里亚蒂,多米妮克·麦克艾丽戈特,杰西·厄舍,拉兹·阿隆索,切斯·克劳福,托默·卡蓬,凯伦·福原,内森·米切尔,蔻碧·米纳菲,珊特尔·范圣滕,妮卡·埃利奥特,阿雅·卡什,莱拉·罗宾斯,乔丹娜·拉茹瓦,吉姆·比弗,周豪,卡梅伦·克罗维蒂,克劳迪娅·多米特,贾森·加里-斯坦福德,肖恩·阿什莫"
	movie.Types = "欧美剧"
	movie.Hash = "9cdf9d2aa95ff0e3dfc9a3ae72e5c2f0"
	err := movie.EsInsertWhenNotExsist()

	a.Equal(nil, err)
}

func TestEsSearch(t *testing.T) {
	a := assert.New(t)
	os.Setenv("ES_ADDRESS", "elasticsearch:9200")
	os.Setenv("ES_TIMEOUT", "10")

	movie := db.NewMovie()

	request := &protos.MovieSearchRequest{
		Keyword: "黑袍",
		ObjId:   "1",
	}
	res, err := movie.EsGetPageData(request, nil, 60)
	a.Equal(nil, err)
	a.IsType([]db.Movie{}, res)
}
