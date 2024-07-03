package mongo_helper

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var mongoHelper *MongoHelper = nil
var mongoClient *MongoClientHelper = nil

func GetMongoHelper() *MongoHelper {
	if mongoHelper == nil {
		mongoHelper = &MongoHelper{}
		mongoHelper.init()
	}
	return mongoHelper
}

type MongoHelper struct {
	clientOption *options.ClientOptions
}

func (mh *MongoHelper) init() {
	mh.clientOption = options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	// 配置连接池选项
	// 连接池上限10
	mh.clientOption.SetMinPoolSize(1)
	mh.clientOption.SetMaxPoolSize(1)
	//mh.clientOption.SetPoolMonitor(&event.PoolMonitor{Event: func(pe *event.PoolEvent) {
	//	msg, _ := json.Marshal(pe)
	//	log.Println(string(msg))
	//}})
	mh.clientOption.SetMaxConnIdleTime(time.Minute * 90)
}

func (mh MongoHelper) GetClientHelper() *MongoClientHelper {
	if mongoClient == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, mh.clientOption)
		if err != nil {
			log.Fatal(err)
		}
		mc := &MongoClientHelper{}
		mc.Client = client
		mongoClient = mc
	}
	return mongoClient
}

type MongoClientHelper struct {
	Client *mongo.Client
}

func (mc *MongoClientHelper) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := mc.Client.Disconnect(ctx); err != nil {
		log.Fatalln("数据库链接关闭失败!" + err.Error())
	}
}

func (mc MongoClientHelper) KeepHeart() {
	for {
		if err := mc.Client.Ping(context.Background(), readpref.Primary()); err != nil {
			log.Println(err.Error())
		}
		time.Sleep(time.Second * 60)
	}
}
