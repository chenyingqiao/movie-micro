package utils

import (
	"sync"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var (
	//AuthGrpcAddress 授权服务
	AuthGrpcAddress = "AUTH"
	//MovieGrpcAddress 电影服务
	MovieGrpcAddress = "MOVIE"

	grpcClientConnect = map[string]*grpc.ClientConn{}
	grpcClientMap     = map[string]string{
		"AUTH":  "auth:50060",
		"MOVIE": "server:30059",
	}
	lock = sync.Mutex{}
)

//OpenGrpcClientConnect 打开grpc链接，不重复开启链接
func OpenGrpcClientConnect(flag string) (*grpc.ClientConn, error) {
	if _, ok := grpcClientConnect[flag]; ok {
		conn := grpcClientConnect[flag]
		if conn.GetState() == connectivity.Ready {
			return grpcClientConnect[flag], nil
		}
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
