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
	Method   string `json:"method"`
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Sex      bool   `json:"sex"`
	Birthday int    `json:"birthday"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func GetUser(u User) *User {

	row := db.QueryRow("select uid,username, password, avatar, sex, birthday, email, phone  from users where username=?", u.Username)

	log.Println("++++")
	var username, password, avatar, email, phone sql.NullString
	var sex sql.NullBool
	var uid, birthday sql.NullInt64
	err := row.Scan(&uid, &username, &password, &avatar, &sex, &birthday, &email, &phone)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with username:%s \n", u.Username)
	case err != nil:
		log.Fatalf("query error %v\n", err)
	default:
		res := &User{
			Uid:      int(uid.Int64),
			Sex:      sex.Bool,
			Birthday: int(birthday.Int64),
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

func PostUser(u User) (int, error) {

	fmt.Println("!!!!")
	stmt, err := db.Prepare(`INSERT INTO users(username, password, sex, birthday, email)
		values (?,?,?,?,?);`)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()
	res, err := stmt.Exec(u.Username, u.Password, u.Sex, u.Birthday, u.Email)
	id, _ := res.LastInsertId()

	return int(id), err
}
