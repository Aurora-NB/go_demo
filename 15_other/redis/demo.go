package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

// 47.108.93.159
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "47.108.93.159:6379",
		Password: "jinweimin",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("连接成功")
}

func main() {
	err := rdb.Set("ab", "1", 0).Err()
	if err != nil {
		fmt.Println("set failed, err: ", err)
		return
	}
	result, err := rdb.Get("ab").Result()
	if err == redis.Nil {
		fmt.Println("key not exist")
	} else if err != nil {
		fmt.Println("get failed, err: ", err)
		return
	} else {
		fmt.Println(result)
	}

	size := rdb.DBSize()
	fmt.Println(size)

	keys := rdb.Keys("*")
	fmt.Println(keys)

}
