package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var KeyName = "test:test"

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.110.83:6379",
		Password: "", // no password set
		DB:       1,  // use DB 1 (Default is 0)
	})

	//Push
	if err := rdb.LPush(ctx, KeyName, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10).Err(); err != nil {
		panic(err)
	} else {
		fmt.Println("Pushed finished...............")
	}

	//Search
	if strKeys, err := rdb.Keys(ctx, "*").Result(); err != nil {
		panic(err)
	} else {
		fmt.Println("Key List: ", strKeys)
	}

	if strValues, err := rdb.LRange(ctx, KeyName, 0, -1).Result(); err != nil {
		panic(err)
	} else {
		fmt.Println("Key: ", KeyName, ", Values: ", strValues)
	}

	if nLen, err := rdb.LLen(ctx, KeyName).Result(); err != nil {
		panic(err)
	} else {
		fmt.Println("Key: ", KeyName, ", Len: ", nLen)
	}
	fmt.Println("Search finished...............")

	// Pop
	for {
		strValue, err := rdb.RPop(ctx, KeyName).Result()
		if err == redis.Nil {
			fmt.Println("Key is empty")
			break
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("Key: ", KeyName, ", Poped Value: ", strValue)
		}
		time.Sleep(time.Second * 2)
	}

	fmt.Println("Pop finished...............")

	//Push
	if err := rdb.LPush(ctx, KeyName, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10).Err(); err != nil {
		panic(err)
	} else {
		fmt.Println("Push again finished...............")
	}

	//BRPop
	for {
		//Wait for seconds
		strValue, err := rdb.BRPop(ctx, time.Duration(1)*time.Second, KeyName).Result()
		if err == redis.Nil {
			fmt.Println("Key is empty")
			break
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("Key: ", KeyName, ", BLPoped Value: ", strValue)
		}
	}

}
