package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
	db  *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:xwt123456789@tcp(172.17.0.2)/forum?charset=utf8")

	checkErr(err)
	err = db.Ping()
	checkErr(err)
	fmt.Println("ping ok")
}
func checkErr(e error) {
	if e != nil {
		panic(e.Error())
	}
}

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Sex      bool   `json:"sex"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func GetUser(u User) *User {

	row := db.QueryRow("select uid,username, password, avatar, sex, age, email, phone  from users where username=?", u.Username)

	log.Println("++++")
	var username, password, avatar, email, phone sql.NullString
	var sex sql.NullBool
	var uid, age sql.NullInt64
	err := row.Scan(&uid, &username, &password, &avatar, &sex, &age, &email, &phone)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with username:%s \n", u.Username)
	case err != nil:
		log.Fatalf("query error %v\n", err)
	default:
		res := &User{
			Uid:      int(uid.Int64),
			Sex:      sex.Bool,
			Age:      int(age.Int64),
			Username: username.String,
			Password: password.String,
			Avatar:   avatar.String,
			Phone:    phone.String,
			Email:    email.String,
		}
		return res
	}
	return nil
}

func PostUser(u User) {

	fmt.Println("!!!!")
	stmt, err := db.Prepare(`INSERT INTO users(username, password, avatar, sex, age, email, phone)
		values (?,?,?,?,?,?,?);`)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()
	_, err = stmt.Exec(u.Username, u.Password, u.Avatar, u.Sex, u.Age, u.Email, u.Phone)
	checkErr(err)
	fmt.Println("post user success")
}
