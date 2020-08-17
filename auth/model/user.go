package model

import (
	"auth/utils"
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Power    []string           `bson:"power"`
	Active   bool               `bson:"active"`
}

//NewUser 实例化新用户
func NewUser() User {
	return User{}
}

func (u *User) Add() error {
	col, err := utils.GetMongoDb("movie")
	if err != nil {
		return errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	result, err := col.InsertOne(ctx, u)
	if result.InsertedID != nil && err == nil {
		return nil
	}
	return errors.Wrap(err, "用户插入失败")
}

func (u *User) Info(username string) (User, error) {
	col, err := utils.GetMongoDb("movie")
	if err != nil {
		return User{}, errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	filter := bson.M{
		"username": username,
	}
	user := NewUser()
	userDecodeErr := col.FindOne(ctx, filter).Decode(&user)
	if userDecodeErr != nil {
		return NewUser(), errors.Wrap(userDecodeErr, "用户查找错误")
	}
	return user, nil
}
