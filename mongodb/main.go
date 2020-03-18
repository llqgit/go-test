package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

var DB *mongo.Client

type Test struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func init() {
	var err error
	uri := "mongodb://localhost:27017"
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	DB, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}

func GetId() int64 {
	ret := Rdb.Incr("mongo-test-id")
	return ret.Val()
}

// 计算函数消耗时间
func TimeCost(start time.Time, name ...string) {
	terminal := time.Since(start)
	if len(name) > 0 {
		fmt.Printf("%vtime cost: %v\n", name, terminal)
	} else {
		fmt.Println("time cost:", terminal)
	}
}

func main() {
	//var err error

	collection := DB.Database("testing1").Collection("numbers")

	ctx, _ := context.WithCancel(context.Background())
	//ctx, _ = context.WithTimeout(context.Background(), 200*time.Millisecond)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 1; i <= 100000; i++ {
				s := Test{
					Name: fmt.Sprintf("hello-%d", GetId()),
					Age:  rand.Int31(),
				}
				res, err := collection.InsertOne(ctx, s)
				if err != nil {
					panic(err)
				}
				id := res.InsertedID
				fmt.Println(i, id)
			}
		}()
	}

	// 查询
	func() {
		defer TimeCost(time.Now())
		collection.FindOne(ctx, bson.M{"name": "hello-1000000"})
	}()

	// 删除索引
	dropIndexRet, err := collection.Indexes().DropOne(ctx, "name_1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dropIndexRet)
	}

	// 查询
	func() {
		defer TimeCost(time.Now())
		collection.FindOne(ctx, bson.M{"name": "hello-1000000"})
	}()

	// 创建索引
	indexName, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{"name", 1}},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(indexName)
	}

	// 查询
	func() {
		defer TimeCost(time.Now())
		collection.FindOne(ctx, bson.M{"name": "hello-1000000"})
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals)

	<-signals
}
