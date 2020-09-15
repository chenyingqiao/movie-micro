package db

import (
	"context"
	"crawler/rpc/protos"
	"crawler/utils"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Movie 电影数据
type Movie struct {
	ID                 primitive.ObjectID `bson:"_id, omitempty"`
	Source             string             `bson:"source"`
	DetailURL          string             `bson:"detail_url"`
	Title              string             `bson:"title"`
	Alias              string             `bson:"alias"`
	Director           string             `bson:"director"`
	Actor              string             `bson:"actor"`
	Types              string             `bson:"types"`
	Location           string             `bson:"location"`
	Language           string             `bson:"language"`
	ShowingTime        string             `bson:"showingt_ime"`
	Long               string             `bson:"long"`
	UpdateTime         string             `bson:"update_time"`
	UpdateTimeUnix     int64              `bson:"update_time_unix"`
	Introduce          string             `bson:"introduct"`
	VideoM3U8Source    []string           `bson:"video_m3u8_source"`
	VideoZuidallSource []string           `bson:"video_zuidall_source"`
	VideoMp4Source     []string           `bson:"video_map_source"`
	ImageURL           string             `bson:"image_url"`
	Hash               string             `bson:"hash"`
}

//NewMovie 实例化Movie
func NewMovie() Movie {
	movie := Movie{}
	return movie
}

//Fill 填充结构体
func (m *Movie) Fill(data map[string][]string) {
	reflectValue := reflect.ValueOf(m)
	for k, v := range data {
		field := reflectValue.Elem().FieldByName(k)
		isValid := field.IsValid()
		if !isValid {
			continue
		}
		typeName := field.Type().String()
		if isValid && typeName == "string" && len(v) > 0 {
			field.SetString(v[0])
		} else if isValid && typeName == "[]string" {
			field.Set(reflect.ValueOf(v))
		}
	}
	m.ID = primitive.NewObjectID()
	m.Hash, _ = m.hashField()
	m.Source = "www.zuidazy5.com"
	updatetime, err := time.Parse("2006-01-02 15:04:05", m.UpdateTime)
	if err != nil {
		updatetime = time.Now()
	}
	m.UpdateTimeUnix = updatetime.Unix()
}

//FillObj 填充其他结构体
func (m *Movie) FillObj(obj *protos.MovieResponse) {
	reflectValue := reflect.ValueOf(obj)
	movieReflectValue := reflect.ValueOf(*m)
	movieReflectType := movieReflectValue.Type()
	for i := 0; i < movieReflectValue.NumField(); i++ {
		movieField := movieReflectValue.Field(i)
		k := movieReflectType.Field(i).Name
		v := movieField.Interface()

		field := reflectValue.Elem().FieldByName(k)
		isValid := field.IsValid()
		if !isValid {
			continue
		}
		typeName := field.Type().String()
		if isValid && typeName == "string" {
			if k == "ID" {
				field.SetString(v.(primitive.ObjectID).Hex())
			} else {
				field.SetString(v.(string))
			}
		} else if isValid && typeName == "[]string" {
			field.Set(reflect.ValueOf(v))
		}
	}
}

func (m *Movie) hashField() (string, error) {
	m.Hash = ""
	var err error
	var buf []byte
	if buf, err = json.Marshal(m); err == nil {
		md5H := md5.New()
		md5H.Write(buf)
		return hex.EncodeToString(md5H.Sum(nil)), nil
	}
	return "", err
}

//InsertWhenNotExsist 进行更新或者删除
func (m *Movie) InsertWhenNotExsist() error {
	col, err := utils.GetMongoDb("movie")
	if err != nil {
		return errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	filter := bson.M{
		"title":  m.Title,
		"source": m.Source,
	}
	movie := Movie{}
	err = col.FindOne(ctx, filter).Decode(&movie)
	if err == nil {
		if m.Hash == movie.Hash {
			return nil
		}
		m.ID = movie.ID
		_, replaceErr := col.ReplaceOne(ctx, filter, m)
		if replaceErr == nil {
			return nil
		}
		return errors.Wrap(replaceErr, "mongodb 替换错误")
	}
	_, insertErr := col.InsertOne(ctx, m)
	if insertErr == nil {
		return nil
	}
	return errors.Wrap(insertErr, "mongodb 插入错误")
}

//GetPageData 获取分页数据
func (m *Movie) GetPageData(filter interface{}, sort interface{}, limit int64) ([]Movie, error) {
	col, err := utils.GetMongoDb(utils.MongoCol)
	if err != nil {
		return nil, errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(utils.MongoQueryTimeout)*time.Second)
	defer cancel()

	if limit == 0 {
		limit = 60
	}
	opts := options.Find().SetSort(sort).SetLimit(limit)
	cursor, err := col.Find(ctx, filter, opts)
	movies := []Movie{}
	for cursor.Next(context.TODO()) {
		movie := NewMovie()
		err := cursor.Decode(&movie)
		if err != nil {
			return nil, errors.Wrap(err, "Decode 错误")
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

//FindByID id查找
func (m *Movie) FindByID(id string) (Movie, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return NewMovie(), err
	}
	col, err := utils.GetMongoDb(utils.MongoCol)
	if err != nil {
		return NewMovie(), errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(utils.MongoQueryTimeout)*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": objID,
	}
	movie := NewMovie()
	err = col.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		return NewMovie(), err
	}
	return movie, nil
}

//FindByHash 通过hash查找movie
func (m *Movie) FindByHash(hash string) (Movie, error) {
	col, err := utils.GetMongoDb(utils.MongoCol)
	if err != nil {
		return NewMovie(), errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(utils.MongoQueryTimeout)*time.Second)
	defer cancel()

	filter := bson.M{
		"hash": hash,
	}
	movie := NewMovie()
	err = col.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		return NewMovie(), err
	}
	return movie, nil
}

//DeleteByHash 删除对应电影数据
func (m *Movie) DeleteByHash(hash string) (bool, error) {
	col, err := utils.GetMongoDb(utils.MongoCol)
	if err != nil {
		return false, errors.Wrap(err, "mongodb 错误")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(utils.MongoQueryTimeout)*time.Second)
	defer cancel()

	filter := bson.M{
		"hash": hash,
	}
	_, err = col.DeleteOne(ctx, filter)
	if err != nil {
		return false, nil
	}
	return true, nil

}

//Watch 监听movie集合的变化
func (m *Movie) Watch(handler func(movie bson.M)) error {
	col, err := utils.GetMongoDb(utils.MongoCol)
	if err != nil {
		return errors.Wrap(err, "mongodb 错误")
	}

	stream, err := col.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		return errors.Wrap(err, "mongodb 错误")
	}
	for stream.Next(context.TODO()) {
		var data bson.M
		if err := stream.Decode(&data); err != nil {
			panic("mongodb watch任务解析错误")
		}
		handler(data)
	}
	return nil
}
