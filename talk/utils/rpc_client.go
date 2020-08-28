package utils

import (
	"os"
	"sync"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
)

var (
	grpcClientConnect *grpc.ClientConn
	once              sync.Once
)

//OpenGrpcClientConnect 获取grpc链接
func OpenGrpcClientConnect() (*grpc.ClientConn, error) {
	authGrpcAddress := os.Getenv("ATUH_SERVICE_HOST") + ":" + os.Getenv("ATUH_SERVICE_PORT")
	if grpcClientConnect != nil {
		return grpcClientConnect, nil
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	once.Do(func() {
		defer wg.Done()
		conn, err := grpc.Dial(authGrpcAddress, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(2))))
		if err != nil {
			return
		}
		grpcClientConnect = conn
	})
	wg.Wait()
	return grpcClientConnect, nil

}

//CloseGrpcClientConnect 关闭grpc
func CloseGrpcClientConnect() {
	if grpcClientConnect != nil {
		grpcClientConnect.Close()
	}
}
