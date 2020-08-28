package model

import (
	"auth/errs"
	"auth/utils"
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User user
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Power    []string           `bson:"power"`
	Active   bool               `bson:"active"`
}

//NewUser 实例化新用户
func NewUser(username string, password string) User {
	return User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: password,
		Power:    []string{},
		Active:   true,
	}
}

//NewEmptyUser 获取未实例化的User
func NewEmptyUser() User {
	return User{}
}

//Add 添加新的用户记录
func (u *User) Add() error {
	col, err := utils.GetMongoDb("user")
	if err != nil {
		return errs.NewDbError("mongodb 初始化错误", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	filter := bson.M{
		"username": u.Username,
	}
	user := NewEmptyUser()
	userDecodeErr := col.FindOne(ctx, filter).Decode(&user)
	if userDecodeErr == nil || !user.ID.IsZero() {
		return errs.NewDbError("用户已经存在", userDecodeErr)
	}

	result, err := col.InsertOne(ctx, u)
	if err == nil && result.InsertedID != nil {
		return nil
	}
	return errs.NewDbError("用户插入失败", userDecodeErr)
}

//Info 获取用户信息
func (u *User) Info(username string) (User, error) {
	col, err := utils.GetMongoDb("user")
	if err != nil {
		return User{}, errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	filter := bson.M{
		"username": username,
	}
	user := NewEmptyUser()
	userDecodeErr := col.FindOne(ctx, filter).Decode(&user)
	if userDecodeErr != nil {
		return NewEmptyUser(), errors.Wrap(userDecodeErr, "用户查找错误")
	}
	return user, nil
}
