package db

import (
	"context"
	"crawler/rpc/protos"
	"crawler/utils"
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

//EsInsertWhenNotExsist 进行更新或者删除
func (m *Movie) EsInsertWhenNotExsist() error {
	esClient, err := utils.GetEsConnect()
	if err != nil {
		return err
	}
	_, err = m.EsFindByHash(m.Hash, nil)
	if err == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	_, err = esClient.Index().Index("movie").Type("movie").BodyJson(m).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

//EsGetPageData 获取分页数据
func (m *Movie) EsGetPageData(filter *protos.MovieSearchRequest, limit int64) ([]Movie, error) {
	esClient, err := utils.GetEsConnect()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	boolQ := elastic.NewBoolQuery()
	boolQ.Should(
		elastic.NewMatchQuery("title", filter.Keyword),
		elastic.NewMatchQuery("alias", filter.Keyword),
		elastic.NewMatchQuery("actor", filter.Keyword),
	)
	pageNumber, err := strconv.Atoi(filter.ObjId)
	if err != nil {
		pageNumber = 1
	}
	pageNumber--
	res, err := esClient.Search("movie").Type("movie").Query(boolQ).From(pageNumber * int(limit)).Size(int(limit)).Do(ctx)

	if err != nil {
		return nil, err
	}

	movies := []Movie{}
	var movie **Movie
	for _, item := range res.Each(reflect.TypeOf(&m)) {
		movie = item.(**Movie)
		movies = append(movies, **movie)
	}
	return movies, nil
}

//EsFindByHash 通过hash查找movie
func (m *Movie) EsFindByHash(hash string, EsId *string) (Movie, error) {
	esClient, err := utils.GetEsConnect()
	if err != nil {
		return NewMovie(), err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	res, err := esClient.Search("movie").Type("movie").Query(
		elastic.NewQueryStringQuery("hash:" + hash),
	).Do(ctx)
	if err != nil {
		return NewMovie(), errors.Wrap(err, "未找到对应文档")
	}

	var movie Movie
	if res.Hits.TotalHits.Value > 0 {
		err = json.Unmarshal(res.Hits.Hits[0].Source, &movie)
		if err != nil {
			return NewMovie(), errors.Wrap(err, "解析返回结果错误")
		}
		EsId = &res.ScrollId
		return movie, nil
	} else {
		return NewMovie(), errors.New("未找到电影")
	}
}

//EsDeleteByHash 删除对应电影数据
func (m *Movie) EsDeleteByHash(hash string) error {
	esClient, err := utils.GetEsConnect()
	if err != nil {
		return err
	}
	_, err = m.EsFindByHash(m.Hash, nil)
	if err == nil {
		return nil
	}

	esClient.Delete().Index("movie").Type("movie")
	return nil
}
