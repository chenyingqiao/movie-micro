package utils

import (
	"os"
	"sync"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
)

var (
	//AuthGrpcAddress 授权服务
	AuthGrpcAddress = "AUTH"
	//MovieGrpcAddress 电影服务
	MovieGrpcAddress = "MOVIE"

	grpcClientConnect = map[string]*grpc.ClientConn{}
	grpcClientMap     = map[string]string{
		"AUTH":  os.Getenv("ATUH_SERVICE_HOST") + ":" + os.Getenv("ATUH_SERVICE_PORT"),
		"MOVIE": os.Getenv("MOVIE_SERVICE_HOST") + ":" + os.Getenv("MOVIE_SERVICE_PORT"),
	}
	lock = sync.Mutex{}
)

//OpenGrpcClientConnect 打开grpc链接，不重复开启链接
func OpenGrpcClientConnect(flag string) (*grpc.ClientConn, error) {
	if _, ok := grpcClientConnect[flag]; ok {
		return grpcClientConnect[flag], nil
	}
	lock.Lock()
	defer lock.Unlock()
	address := grpcClientMap[flag]
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(2))))
	if err != nil {
		return nil, err
	}
	grpcClientConnect[flag] = conn
	return conn, nil
}

//CloseAllGrpcClientConnect 关闭grpc
func CloseAllGrpcClientConnect() {
	for _, v := range grpcClientConnect {
		v.Close()
	}
	grpcClientConnect = map[string]*grpc.ClientConn{}
}
