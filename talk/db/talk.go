package db

import (
	"context"
	"talk/utils"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Talk struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	Username   string             `bson:"username"`
	Content    string             `bson:"content"`
	Createtime string             `bson:"createtime"`
}

//NewTalk 实例化talk
func NewTalk(username string, content string, createtime string) Talk {
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
