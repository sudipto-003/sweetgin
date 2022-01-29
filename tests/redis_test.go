package tests

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestRedisConnection(t *testing.T) {

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "0.0.0.0:6379",
	// 	DB:   0,
	// })
	opts, err := redis.ParseURL("redis://0.0.0.0:6379/0")
	if err != nil {
		t.Fatalf("%q\n", err)
	}
	rdb := redis.NewClient(opts)
	val := "hello"
	err = rdb.Set(ctx, "key", val, 0).Err()
	if err != nil {
		t.Fatalf("%q\n", err)
	}

	value, err := rdb.Get(ctx, "key").Result()
	if err == redis.Nil {
		t.Errorf("key don't exist in redis")
	}
	if err != nil {
		t.Fatalf("%q\n", err)
	}

	if val != value {
		t.Errorf("Value mismatch. want=%q got=%q", val, value)
	}
}
