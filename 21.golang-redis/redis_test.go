package golangredis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	// err := client.Close()

	// fmt.Println(client)

	// assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, result)
	fmt.Println(result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Muhamad Isro", time.Second*3)

	result, err := client.Get(ctx, "name").Result()

	assert.Nil(t, err)
	assert.Equal(t, "Muhamad Isro", result)

	time.Sleep(time.Second * 5)

	_, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "Muhamad")
	client.RPush(ctx, "names", "Isro")
	client.RPush(ctx, "names", "Sabanur")

	assert.Equal(t, "Muhamad", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Isro", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Sabanur", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "Muhamad")
	client.SAdd(ctx, "students", "Muhamad")
	client.SAdd(ctx, "students", "Isro")
	client.SAdd(ctx, "students", "Isro")
	client.SAdd(ctx, "students", "Sabanur")
	client.SAdd(ctx, "students", "Sabanur")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"Muhamad", "Isro", "Sabanur"}, client.SMembers(ctx, "students").Val())
}

func TestSortedList(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Muhamad"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Isro"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 75, Member: "Sabanur"})

	assert.Equal(t, []string{"Sabanur", "Isro", "Muhamad"}, client.ZRange(ctx, "scores", 0, 2).Val())
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "Muhamad")
	client.HSet(ctx, "user:1", "email", "isro@gmail.com")

	user := client.HGetAll(ctx, "user:1").Val()

	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Muhamad", user["name"])
	assert.Equal(t, "isro@gmail.com", user["email"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko A",
		Longitude: 106.8227,
		Latitude:  -6.177590,
	})
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko B",
		Longitude: 106.8227,
		Latitude:  -6.177590,
	})

	distance := client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val()
	assert.Equal(t, 0.35423, distance)

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  106.8218128,
		Latitude:   -6.12323,
		Radius:     5,
		RadiusUnit: "km",
	})

	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "muhamad", "isro", "sabanur")
	client.PFAdd(ctx, "visitors", "muhamad", "reza", "sabanur")
	client.PFAdd(ctx, "visitors", "muhamad", "mama", "sabanur")

	assert.Equal(t, int64(4), client.PFCount(ctx, "visitors").Val())
}

func TestPipeline(t *testing.T) {
	_, err := client.Pipelined(ctx, func(pipeline redis.Pipeliner) error {

		pipeline.SetEx(ctx, "name", "roozy", time.Second*5)
		pipeline.SetEx(ctx, "address", "indonesia", time.Second*5)

		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, "roozy", client.Get(ctx, "name").Val())
	assert.Equal(t, "indonesia", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(pipeline redis.Pipeliner) error {
		pipeline.SetEx(ctx, "name", "roozy", time.Second*5)
		pipeline.SetEx(ctx, "address", "indonesia", time.Second*5)

		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, "roozy", client.Get(ctx, "name").Val())
	assert.Equal(t, "indonesia", client.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":    "Muhamad",
				"address": "Indonesia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-2", "consumer-2")
}

func TestGetStream(t *testing.T) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    time.Second * 5,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.Values)
		}
	}
}

func TestSubscriber(t *testing.T) {
	pubSub := client.Subscribe(ctx, "subscribe-1")
	defer pubSub.Close()
	for i := 0; i < 10; i++ {
		message, _ := pubSub.ReceiveMessage(ctx)
		fmt.Println(message)
	}
}

func TestPubl(t *testing.T) {
	for i := 0; i < 10; i++ {
		client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i))
	}
}
