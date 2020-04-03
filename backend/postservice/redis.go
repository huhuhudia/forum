package main

import (
	"fmt"
	"log"
	"strconv"
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

func getMessageSlotKey(uid int) string {
	return fmt.Sprintf("user-%d-slot", uid)
}

func getReplySlotKey(uid int) string {
	return fmt.Sprintf("user-%d-reply", uid)
}

func pushMessage(uid int, message string, generateKey func(int) string) error {
	k := generateKey(uid)
	return rdb.LPush(k, message).Err()
}

func getMessage(uid, limit int, generateKey func(int) string) ([]string, error) {
	k := generateKey(uid)

	res, err := rdb.LRange(k, 0, int64(limit-1)).Result()

	if err != nil {
		return res, err
	}
	return res, nil
}

func generateFansKey(id int) string {
	return fmt.Sprintf("user-%d-fans", id)
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

func incrScore(tid int, key string) {
	rdb.Do("zincrby", key, 1, tid)
}

func getHotestTopic(limit int, key string) []Topic {
	opt := &redis.ZRangeBy{}
	opt.Max = "+inf"
	opt.Min = "-inf"
	cmd := rdb.ZRevRangeByScoreWithScores(key, opt)
	res, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return []Topic{}
	}
	topics := []Topic{}
	for i := 0; i < len(res); i++ {
		member := res[i].Member.(string)
		if err != nil {
			log.Println(err)
			continue
		}
		uid, err := strconv.Atoi(member)
		if err != nil {
			log.Println(err)
			continue
		}

		score := int(res[i].Score)

		tmp, err := GetTopic(uid)
		if err != nil {
			log.Println(err)
			continue
		}

		tmp.Score = score
		topics = append(topics, *tmp)
	}
	return topics
}
