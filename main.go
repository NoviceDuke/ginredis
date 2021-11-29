package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.TODO()
var rclient *redis.Client

func main() {
	rclient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	pong, err := rclient.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(pong)
	executeStringCommand()
}
func executeStringCommand() {
	fmt.Println("-------Command start-----")

	// set a key hello and value is "world" without expired time
	err := rclient.Set(ctx, "hello", "world", 0).Err()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	val, err := rclient.Get(ctx, "hello").Result()
	handleResult(val, err)
	err = rclient.Set(ctx, "hello", "world", time.Minute).Err()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("-----Command End")
}

func handleResult(result interface{}, err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(result)
}
