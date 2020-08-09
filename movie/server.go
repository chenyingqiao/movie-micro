package main

import (
	"fmt"
	"movie/rpc/protos"
	rpc_service "movie/rpc/service"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50059"
)

func main() {
	fmt.Println("开始")
	Run()
}

//Run run
func Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "tcp监听错误"+port)
	}
	server := grpc.NewServer()
	protos.RegisterMovieServer(server, &rpc_service.MovieService{})
	if err := server.Serve(lis); err != nil {
		return errors.Wrap(err, "rpc服务启动错误"+port)
	}
	return nil
}