package service

import (
	"context"
	"crawler/db"
	"crawler/logic"
	"crawler/rpc/protos"
	"crawler/utils"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//MovieService 电影服务
type MovieService struct {
}

//Detail 电影详细信息
func (m *MovieService) Detail(ctx context.Context, movieRequest *protos.MovieRequest) (*protos.MovieResponse, error) {
	movie := db.NewMovie()
	movieInfo, err := movie.FindByHash(movieRequest.GetHash())
	if err != nil {
		return nil, errors.Wrap(err, "未找到文档")
	}
	movieResponse := &protos.MovieResponse{}
	movieInfo.FillObj(movieResponse)
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
		// "types": bson.M{
		// 	"$nin": bson.A{
		// 		"福利片",
		// 		"伦理片",
		// 	},
		// },
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
	return err
}

//Search 查找电影
func (*MovieService) Search(request *protos.MovieSearchRequest, searchServer protos.Movie_SearchServer) error {

	// movie := db.NewMovie()

	return status.Errorf(codes.Unimplemented, "method Search not implemented")
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
