package utils

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoClient *mongo.Client
	//MongoCol 默认文档集合
	MongoCol = "talk"
	//MongoQueryTimeout 查询超时
	MongoQueryTimeout = 10
)

//GetMongoDb 获取mongodb的链接
func GetMongoDb(col string) (*mongo.Collection, error) {
	//从环境变量中获取mongodb的用户名和密码数据
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PWD")
	addr := os.Getenv("MONGODB_ADDR")
	db := os.Getenv("MONGODB_DB")
	timeout, err := strconv.Atoi(os.Getenv("MONGODB_TIMEOUT"))
	if err != nil {
		return nil, errors.Wrap(err, "MONGODB_TIMEOUT_NOT_INT")
	}
	poolSize, err := strconv.ParseUint(os.Getenv("MONGODB_POOLSIZE"), 20, 64)
	if err != nil {
		return nil, errors.Wrap(err, "MONGODB_POOLSIZE_NOT_INT")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s&&ssl=false", username, password, addr, "admin")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if mongoClient != nil {
		pingErr := mongoClient.Ping(ctx, readpref.Primary())
		if pingErr != nil {
			panic("无法连接数据库")
		}
		return mongoClient.Database(db).Collection(col), nil
	}
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(uri).SetMaxPoolSize(poolSize),
	)
	if err != nil {
		return nil, err
	}
	mongoClient = client
	pingErr := client.Ping(ctx, readpref.Primary())
	if pingErr != nil {
		panic("无法连接数据库")
	}

	return client.Database(db).Collection(col), nil
}
