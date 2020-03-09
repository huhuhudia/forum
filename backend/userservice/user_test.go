package main

import (
	"fmt"
	"testing"
)

func TestPostUser(t *testing.T) {
	u := User{
		Username: "jack",
		Password: "xwt123456789",
		Avatar:   "www.baidu.com",
		Sex:      false,
		Age:      22,
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
		Age:      22,
		Email:    "yellowbluewhite@foxmail.com",
		Phone:    "13343423220",
	}
	res := GetUser(u)
	fmt.Printf("=====%v", res.Avatar)
}
