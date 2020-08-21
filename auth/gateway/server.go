package main

import (
	"auth/rpc/protos"
	"context"
	"fmt"
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

	err := protos.RegisterAuthHandlerFromEndpoint(ctx, mux, "127.0.0.1:50060", opts)
	if err != nil {
		return errors.Wrap(err, "创建endopint失败")
	}
	protos.RegisterUserHandlerFromEndpoint(ctx, mux, "127.0.0.1:50060", opts)
	if err != nil {
		return errors.Wrap(err, "创建endopint失败")
	}
	return http.ListenAndServe(":8090", mux)
}
