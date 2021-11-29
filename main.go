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
	executeListCommand()
	executeHashCommand()
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
	times, err := rclient.TTL(ctx, "hello").Result()
	handleResult(times, err)
	err = rclient.Del(ctx, "hello").Err()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	val, err = rclient.Get(ctx, "hello").Result()
	handleResult(val, err)
	fmt.Println("-----Command End----")
}
func executeListCommand() {
	fmt.Println("-------List Command start-----")
	err := rclient.LPush(ctx, "test", "妳好").Err()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	list, err := rclient.LRange(ctx, "test", 0, -1).Result()
	handleResult(list, err)
	val, err := rclient.LPop(ctx, "test").Result()
	handleResult(val, err)
	list, err = rclient.LRange(ctx, "test", 0, -1).Result()
	handleResult(list, err)
	err = rclient.Del(ctx, "test").Err()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	list, err = rclient.LRange(ctx, "test", 0, -1).Result()
	handleResult(list, err)
	fmt.Println("-------List Command End-----")
}
func executeHashCommand() {
	fmt.Println("-------Hash Command start-----")
	// k := []string{"1", "2", "3"}
	m := map[string]string{
		"a": "b",
		"c": "d",
		"e": "f",
	}
	err := rclient.HSet(ctx, "k", m).Err()
	if err != nil {
		fmt.Println("ERROR123: ", err)
		return
	}
	val, err := rclient.HGet(ctx, "k", "a").Result()
	handleResult(val, err)
	fmt.Println("-------Hash Command End-----")
}
func handleResult(result interface{}, err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(result)
}
