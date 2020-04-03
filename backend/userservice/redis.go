package main

import (
	"fmt"
	"log"
	"time"

	redis "github.com/go-redis/redis/v7"
)

var rdb *redis.Client

func init() {

	rdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	res := rdb.Ping()

	log.Println(res)
}
func generateFollowKey(id int) string {
	return fmt.Sprintf("user-%d-follow", id)
}
func generateFansKey(id int) string {
	return fmt.Sprintf("user-%d-fans", id)
}
func addFollow(selfId, followId int) {
	rdb.SAdd(generateFollowKey(selfId), followId)
}

func addFans(selfId, fansId int) {
	rdb.SAdd(generateFansKey(selfId), fansId)
}

func getFollow(selfId int) ([]string, error) {
	s := generateFollowKey(selfId)
	res, err := rdb.SMembers(s).Result()
	if err != nil {
		return nil, err
	}
	log.Printf("%v", res)

	return res, nil
}

func getFans(selfId int) ([]string, error) {
	s := generateFansKey(selfId)
	res, err := rdb.SMembers(s).Result()
	if err != nil {
		return nil, err
	}
	log.Printf("%v", res)

	return res, nil
}
