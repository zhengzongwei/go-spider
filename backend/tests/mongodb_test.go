package tests

import (
	"backend/administrator/controller"
	mongo_helper "backend/mongo-helper"
	"backend/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_connect_mongodb(t *testing.T) {
	mongoHelper := mongo_helper.GetMongoHelper().GetClientHelper()
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

func Test_Pools(t *testing.T) {
	mongoHelper := mongo_helper.GetMongoHelper().GetClientHelper()
	defer mongoHelper.Disconnect()

	client := mongoHelper.Client

	wg := sync.WaitGroup{}
	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			defer wg.Done()
			users := make([]interface{}, 1000000)
			for i := 0; i < 1000000; i++ {
				users[i] = bson.M{"uuid": uuid.New().String(), "test-index": string(rune(j)) + "-" + string(rune(i))}
			}
			time.Sleep(time.Second * 20)
			start := utils.TimeStamp()
			_, _ = client.Database("administrator").Collection("test_01").InsertMany(context.Background(), users)
			log.Printf("插入耗时: %d\n", utils.TimeStamp()-start)
		}()
	}

	wg.Wait()
}
