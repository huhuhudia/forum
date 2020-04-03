package main

import (
	"fmt"
	"testing"
)

func TestPostUser(t *testing.T) {
	u := User{
		Username: "test-username",
		Password: "test-password",
		Avatar:   "www.baidu.com",
		Sex:      false,
		Birthday: 22123,
		Email:    "yellowbluewhite@foxmail.com",
		Phone:    "13343423220",
	}
	PostUser(u)
}

func TestGetUser(t *testing.T) {
	u := User{
		Username: "jack",
		Password: "",
		Avatar:   "",
		Sex:      false,
		Birthday: 22,
		Email:    "yellowbluewhite@foxmail.com",
		Phone:    "13343423220",
	}
	res := GetUser(u)
	fmt.Printf("=====%v", res)
}

func TestSetFans(t *testing.T) {
	addFans(33, 2)
	addFans(33, 3)
}

func TestSetFollow(t *testing.T) {
	addFollow(33, 4)
	addFollow(1, 5)
}

func TestGetFollow(t *testing.T) {
	res, _ := getFollow(1)
	fmt.Println(res)
}

func TestGetFans(t *testing.T) {
	res, _ := getFans(1)
	fmt.Println(res)
}
