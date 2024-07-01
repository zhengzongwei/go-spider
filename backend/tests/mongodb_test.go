package tests

import (
	"backend/administrator/controller"
	mongo_helper "backend/mongo-helper"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_connect_mongodb(t *testing.T) {
	mongoHelper := mongo_helper.GetMongoHelper().GetConnection()
	defer mongoHelper.Disconnect()

	collection := mongoHelper.Client.Database("administrator").Collection("sys_user")
	result := collection.FindOne(context.TODO(), bson.M{"username": "go-spider@golang.com"})

	user := &controller.Users{}
	if err := result.Decode(user); err != nil {
		panic(err)
	}
	// 时间戳读不出来
	fmt.Println(user.LoginTime)
}
