package service

import (
	"context"
	"movie/rpc/protos"
)

//MovieService 电影服务
type MovieService struct {
}

//Detail 电影详细信息
func (m *MovieService) Detail(context.Context, *protos.MovieRequest) (*protos.MovieResponse, error) {
	return nil, nil
}

//List 电影列表
func (m *MovieService) List(*protos.MovieRequest, protos.Movie_ListServer) error {
	return nil
}

//Delete 删除电影
func (m *MovieService) Delete(context.Context, *protos.MovieRequest) (*protos.MovieDeleteRequest, error) {
	return nil, nil
}

//ReCrawler 重爬电影信息
func (m *MovieService) ReCrawler(context.Context, *protos.MovieRequest) (*protos.MovieResponse, error) {
	return nil, nil
}
