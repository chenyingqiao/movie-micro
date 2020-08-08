package main

import (
	"context"
	"fmt"
	"movie/rpc/protos"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("error")
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	err := protos.RegisterMovieHandlerFromEndpoint(ctx, mux, "127.0.0.1:50059", opts)
	if err != nil {
		return errors.Wrap(err, "创建endopint失败")
	}
	return http.ListenAndServe(":8090", mux)
}
