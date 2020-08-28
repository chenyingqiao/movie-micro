package logic

import (
	"talk/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TalkLogic struct {
}

//LoadPage 获取分页数据
func (t *TalkLogic) LoadPage(lastID string) []db.Talk {
	talk := db.NewEmptyTalk()
	var objID primitive.ObjectID
	var err error
	if lastID != "" {
		objID, err = primitive.ObjectIDFromHex(lastID)
		if err != nil {
			objID = primitive.NewObjectIDFromTimestamp(time.Now().Add(3600 * time.Second))
		}
	} else {
		objID = primitive.NewObjectIDFromTimestamp(time.Now().Add(3600 * time.Second))
	}

	filter := bson.M{
		"_id": bson.M{
			"$lt": objID,
		},
	}

	sort := bson.M{
		"_id": 1,
	}

	list, err := talk.GetList(filter, sort, 300)
	if err != nil {
		return []db.Talk{}
	}
	return list
}

//Submit 添加聊天记录
func (t *TalkLogic) Submit(username string, content string) bool {
	talk := db.NewTalk(username, content, time.Now().Unix())
	err := talk.Add()
	if err != nil {
		return false
	}
	return true
}