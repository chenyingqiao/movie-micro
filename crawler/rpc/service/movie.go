package service

import (
	"context"
	"crawler/db"
	"crawler/logic"
	"crawler/rpc/protos"
	"crawler/utils"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MovieService 电影服务
type MovieService struct {
}

//Detail 电影详细信息
func (m *MovieService) Detail(ctx context.Context, movieRequest *protos.MovieRequest) (*protos.MovieResponse, error) {
	movie := db.NewMovie()
	movieInfo, err := movie.FindByID(movieRequest.GetObjId())
	if err != nil {
		return nil, errors.Wrap(err, "未找到文档")
	}
	movieResponse := &protos.MovieResponse{}
	movieInfo.FillObj(movieResponse)
	logrus.WithField("info", movieRequest).Info("电影详细信息")
	return movieResponse, nil
}

//List 电影列表
func (m *MovieService) List(movieRequest *protos.MovieRequest, movieListServer protos.Movie_ListServer) error {
	movie := db.NewMovie()
	var err error
	var objID primitive.ObjectID
	if movieRequest.GetObjId() != "" {
		objID, err = primitive.ObjectIDFromHex(movieRequest.GetObjId())
		if err != nil {
			objID = primitive.NewObjectIDFromTimestamp(time.Now().Add(3600 * time.Second))
		}
	} else {
		objID = primitive.NewObjectIDFromTimestamp(time.Now().Add(3600 * time.Second))
	}
	types := movieRequest.GetType()
	filter := bson.M{
		"_id": bson.M{
			"$lt": objID,
		},
		"types": bson.M{
			"$nin": bson.A{
				"福利片",
				"伦理片",
			},
		},
	}
	if types == "ALL" {
		filter = bson.M{
			"_id": bson.M{
				"$lt": objID,
			},
		}
	}
	if types != "ALL" && types != "" {
		filter = bson.M{
			"_id": bson.M{
				"$lt": objID,
			},
			"types": types,
		}
	}

	sort := bson.M{
		"_id": -1,
	}
	movies, err := movie.GetPageData(filter, sort, movieRequest.Limit)
	for _, v := range movies {
		movieResponse := &protos.MovieResponse{}
		v.FillObj(movieResponse)
		movieListServer.Send(movieResponse)
	}
	logrus.WithField("info", movieRequest).Info("电影列表")
	return err
}

//Search 查找电影
func (*MovieService) Search(request *protos.MovieSearchRequest, searchServer protos.Movie_SearchServer) error {
	host := os.Getenv("ES_ADDRESS")
	if host == "" {
		return mongoSearch(request, searchServer)
	}
	return esSerach(request, searchServer)
}

//mongodbSearch
func mongoSearch(request *protos.MovieSearchRequest, searchServer protos.Movie_SearchServer) error {
	movie := db.NewMovie()
	var err error
	var objID primitive.ObjectID
	if request.GetObjId() != "" {
		objID, err = primitive.ObjectIDFromHex(request.GetObjId())
		if err != nil {
			objID = primitive.NewObjectIDFromTimestamp(time.Now().Add(3600 * time.Second))
		}
	} else {
		objID = primitive.NewObjectIDFromTimestamp(time.Now().Add(3600 * time.Second))
	}
	filter := bson.M{
		"_id": bson.M{
			"$lt": objID,
		},
		"title": bson.M{
			"$regex": primitive.Regex{
				Pattern: request.GetKeyword(),
				Options: "i",
			},
		},
	}
	sort := bson.M{
		"_id": -1,
	}
	movies, err := movie.GetPageData(filter, sort, 60)
	for _, v := range movies {
		movieResponse := &protos.MovieResponse{}
		v.FillObj(movieResponse)
		searchServer.Send(movieResponse)
	}

	logrus.WithField("info", request).Info("查找电影")
	return err
}

//elasticsearch搜索
func esSerach(request *protos.MovieSearchRequest, searchServer protos.Movie_SearchServer) error {
	movie := db.NewMovie()
	movies, err := movie.EsGetPageData(request, 120)
	for _, v := range movies {
		movieResponse := &protos.MovieResponse{}
		v.FillObj(movieResponse)
		searchServer.Send(movieResponse)
	}

	logrus.WithField("info", request).Info("查找电影es")
	return err
}

//Delete 删除电影
func (m *MovieService) Delete(ctx context.Context, movieRequest *protos.MovieRequest) (*protos.MovieDeleteResponse, error) {
	movie := db.NewMovie()
	isSuccess, err := movie.DeleteByHash(movieRequest.GetHash())
	if err != nil {
		return nil, err
	}
	status := "success"
	if !isSuccess {
		status = "error"
	}
	movieDeleteResponse := &protos.MovieDeleteResponse{
		Status: status,
	}
	return movieDeleteResponse, nil
}

//ReCrawler 重爬电影信息
func (m *MovieService) ReCrawler(ctx context.Context, movieRequest *protos.MovieRequest) (*protos.MovieResponse, error) {
	opts := utils.NewCrawlerOptionNotCmd()
	crawlerCtx, cancel := logic.GetOptionsDeadlineContext(opts)
	defer cancel()

	//从数据库中获取文档
	movie := db.NewMovie()
	movieInfo, err := movie.FindByHash(movieRequest.GetHash())
	if err != nil {
		return nil, err
	}

	rule := db.NewRule()
	ruleParse := utils.NewRuleParseQuery()
	logic := logic.NewCrawlerLogic(crawlerCtx, ruleParse, opts)
	movieRule, err := rule.GetBySource(movieInfo.Source)
	if err != nil {
		return nil, err
	}
	crawlerMovie, err := logic.CrawlerDetail(movieInfo.DetailURL, movieRule)
	if err != nil {
		return nil, err
	}

	movieResponse := &protos.MovieResponse{}
	crawlerMovie.FillObj(movieResponse)
	return movieResponse, nil
}
