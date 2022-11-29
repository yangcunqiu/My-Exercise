package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "124.221.123.87:6379",
		Password: "nike5510",
		DB:       0,
	})

	err := rdb.Set(ctx, "name", "ycq", time.Minute*5).Err()
	if err != nil {
		log.Println(err)
	}

	name := rdb.Get(ctx, "name")
	log.Println(name)

}
