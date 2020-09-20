package utils

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

var (
	esClient *elastic.Client
)

//GetEsConnect
func GetEsConnect() (*elastic.Client, error) {
	host := os.Getenv("ES_ADDRESS")
	host = "http://" + host + "/"
	var err error
	esClient, err = elastic.NewClient(
		elastic.SetURL(host),
	)
	if err != nil {
		return nil, errors.Wrap(err, "新建连接错误")
	}

	timeout, err := strconv.Atoi(os.Getenv("ES_TIMEOUT"))
	if err != nil {
		return nil, errors.Wrap(err, "ES_TIMEOUT_NOT_INT")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	_, _, err = esClient.Ping(host).Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "es 无法连接")
	}
	_, err = esClient.ElasticsearchVersion(host)
	if err != nil {
		return nil, errors.Wrap(err, "获取version失败")
	}

	return esClient, nil
}
