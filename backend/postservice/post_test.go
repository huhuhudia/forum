package main

import (
	"encoding/json"
	"fmt"

	"testing"
)

func TestPostTopic(t *testing.T) {
	tp := Topic{
		Author_id:   1,
		Title:       "nice day",
		Content:     "ok",
		Author_name: "xuwentao",
	}
	id, err := PostTopic(tp)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("primary key tid: ", id)
}
func TestPostComment(t *testing.T) {
	tp := Comment{
		Tid:         100,
		Author_id:   1,
		Title:       "nice day",
		Content:     "ok",
		Author_name: "xuwentao",
	}
	err := PostComment(tp)

	if err != nil {
		fmt.Println(err)
	}

}

func TestGetComments(t *testing.T) {
	cmts, err := GetCommentsOfTopic(100)
	if err != nil {
		fmt.Printf("%v", err)
	}
	for k, v := range cmts {
		fmt.Printf("k:%v   |  %v\n", k, v)
	}

}

func TestGetTopic(t *testing.T) {
	topic, err := GetTopic(3)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v", topic)
}

func TestGetCommentsOfAuthor(t *testing.T) {
	cmts, err := GetCommentsOfAuthor(1)
	if err != nil {
		fmt.Printf("%v", err)
	}
	for k, v := range cmts {
		fmt.Printf("k:%v   |  %v\n", k, v)
	}
}

func TestGetAllTopics(t *testing.T) {
	topics, err := GetAllTopic(2, 2)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for k, v := range topics {
		fmt.Printf("k:%v  | %v\n ", k, v)
	}
}

func TestGetAllTopicsCotains(t *testing.T) {
	topics, err := GetAllTopicContains("nice", 0, 100)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for k, v := range topics {
		fmt.Printf("k:%v  | %v\n ", k, v)
	}
}

func TestPushMessage(t *testing.T) {
	pushMessage(1, "ni", getReplySlotKey)
	pushMessage(1, "hao", getReplySlotKey)
	pushMessage(1, "ya", getReplySlotKey)
}

func TestGetMessage(t *testing.T) {
	res, err := getMessage(2, 2, getReplySlotKey)
	if err != nil {
		fmt.Printf("%v", err)
	}
	s, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%s", s)
}

func TestFindString(t *testing.T) {
	res := FindReplyToWho("@yellowbluewhite @xuwentao")
	fmt.Println(res)
}

func TestHot(t *testing.T) {
	res := getHotestTopic(0, "testhot")
	fmt.Printf("%v", res)

}
