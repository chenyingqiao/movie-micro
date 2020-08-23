cd ./rpc/protos/protobuffer

protoc -I/usr/local/include -I.  -I/home/lerko/go/src -I/home/lerko/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis --go_out=plugins=grpc:../ ./*.proto

protoc -I/usr/local/include -I.  -I/home/lerko/go/src -I/home/lerko/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis --grpc-gateway_out=logtostderr=true:../ .//*.proto


cd ../../../
go-assets-builder tmpl -o ./utils/assets.go