package db

import (
	"context"
	"talk/errs"
	"talk/utils"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Talk struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	Username   string             `bson:"username"`
	Content    string             `bson:"content"`
	Createtime int64              `bson:"createtime"`
}

//NewTalk 实例化talk
func NewTalk(username string, content string, createtime int64) Talk {
	return Talk{
		Username:   username,
		Content:    content,
		Createtime: createtime,
	}
}

//NewEmptyTalk 实例化talk
func NewEmptyTalk() Talk {
	return Talk{}
}

//Add 插入聊天数据
func (t *Talk) Add() error {
	col, err := utils.GetMongoDb(utils.MongoCol)
	if err != nil {
		return errors.Wrap(err, "mongodb 初始化错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(utils.MongoQueryTimeout)*time.Second)
	defer cancel()

	_, insertErr := col.InsertOne(ctx, t)
	if err != nil {
		return errors.Wrap(insertErr, "mongodb 插入错误")
	}
	return nil
}

//GetList 获取发言列表
func (t *Talk) GetList(filter interface{}, sort interface{}, limit int64) ([]Talk, error) {
	col, err := utils.GetMongoDb(utils.MongoCol)
	if err != nil {
		return nil, errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(utils.MongoQueryTimeout)*time.Second)
	defer cancel()

	if limit == 0 {
		limit = 300
	}

	talks := []Talk{}

	opts := options.Find().SetSort(sort).SetLimit(limit)
	cursor, err := col.Find(ctx, filter, opts)
	if err != nil {
		return nil, errs.NewDbError("mongodb 查询错误", err)
	}
	for cursor.Next(context.TODO()) {
		talk := NewEmptyTalk()
		err := cursor.Decode(&talk)
		if err != nil {
			return nil, errs.NewDbError("mongodb 查询 Decode 错误", err)
		}
		talks = append(talks, talk)
	}
	return talks, nil
}
