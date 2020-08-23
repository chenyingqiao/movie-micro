package client

import (
	"context"
	"io"
	"movie/errs"
	"movie/rpc/protos"
	"movie/utils"
	"time"
)

//MovieClient 电影信息请求
type MovieClient struct {
}

//NewMovieClient movieclient
func NewMovieClient() MovieClient {
	return MovieClient{}
}

//Detail 电影的详细信息
func (m *MovieClient) Detail(movieRequest *protos.MovieRequest) (protos.MovieResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := utils.OpenGrpcClientConnect()
	if err != nil {
		return protos.MovieResponse{}, errs.NewGrpcError("获取grpc链接失败", err)
	}

	movieClient := protos.NewMovieClient(conn)
	response, err := movieClient.Detail(ctx, movieRequest)
	if err != nil {
		return protos.MovieResponse{}, errs.NewGrpcError("grpc:Movie Detail获取失败", err)
	}
	return *response, nil
}

//List 获取电影列表
func (m *MovieClient) List(movieRequest *protos.MovieRequest) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := utils.OpenGrpcClientConnect()
	if err != nil {
		return nil, errs.NewGrpcError("获取grpc链接失败", err)
	}
	movieClient := protos.NewMovieClient(conn)
	stream, err := movieClient.List(ctx, movieRequest)
	if err != nil {
		return nil, errs.NewGrpcError("数据获取失败", err)
	}
	responses := []interface{}{}
	var item *protos.MovieResponse
	for {
		item, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return nil, err
		}

		responses = append(responses, *item)
	}
	return responses, nil
}

//Delete 删除电影
func (m *MovieClient) Delete(movieRequest *protos.MovieRequest) (*protos.MovieDeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := utils.OpenGrpcClientConnect()
	if err != nil {
		return nil, errs.NewGrpcError("获取grpc链接失败", err)
	}
	movieClient := protos.NewMovieClient(conn)

	movieResponse, err := movieClient.Delete(ctx, movieRequest)
	if err != nil {
		return nil, errs.NewGrpcError("grpc:Movie 删除失败", err)
	}
	return movieResponse, nil
}
