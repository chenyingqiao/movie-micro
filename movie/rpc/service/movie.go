package service

import (
	"context"
	"movie/db"
	"movie/rpc/protos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MovieService 电影服务
type MovieService struct {
}

//Detail 电影详细信息
func (m *MovieService) Detail(ctx context.Context, movieRequest *protos.MovieRequest) (*protos.MovieResponse, error) {
	movie := db.NewMovie()
	movieInfo, err := movie.FindByHash(movieRequest.GetHash())
	if err != nil {
		return nil, err
	}
	movieResponse := &protos.MovieResponse{}
	movieInfo.FillObj(movieResponse)
	return movieResponse, nil
}

//List 电影列表
func (m *MovieService) List(movieRequest *protos.MovieRequest, movieListServer protos.Movie_ListServer) error {
	movie := db.NewMovie()
	objID, err := primitive.ObjectIDFromHex(movieRequest.GetObjId())
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": bson.M{
			"$gt": objID,
		},
	}
	sort := bson.M{
		"_id": 1,
	}
	movies, err := movie.GetPageData(filter, sort, movieRequest.Limit)
	for _, v := range movies {
		movieResponse := &protos.MovieResponse{}
		v.FillObj(movieResponse)
		movieListServer.Send(movieResponse)
	}
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
func (m *MovieService) ReCrawler(context.Context, *protos.MovieRequest) (*protos.MovieResponse, error) {
	return nil, nil
}
