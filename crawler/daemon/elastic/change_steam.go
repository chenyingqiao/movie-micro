package main

import (
	"crawler/db"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

//mongodb changestream 同步到es
func main() {
	movie := db.NewMovie()
	err := movie.Watch(func(data bson.M) {
		fmt.Println(data)
	})
	if err != nil {
		fmt.Println(err)
	}
}
