package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

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

type Topic struct {
	Method      string `json:"method"`
	Tid         int    `json:"tid"`
	Author_id   int    `json:"author_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Up          int    `json:"up"`
	Down        int    `json:"down"`
	Date        int    `json:"date"`
	Author_name string `json:"author_name"`
	Score       int    `json:"score"`
}

type Comment struct {
	Method      string `json:"method"`
	Cid         int    `json:"cid"`
	Tid         int    `json:"tid"`
	Author_id   int    `json:"author_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Up          int    `json:"up"`
	Down        int    `json:"down"`
	Author_name string `json:"author_name"`
	Date        int    `json:"date"`
}

func PostTopic(topic Topic) (int, error) {

	fmt.Println("!!!!")
	stmt, err := db.Prepare(`INSERT INTO topics(author_id, title, content, date,author_name)
		values (?,?,?,?,?);`)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()
	res, err := stmt.Exec(topic.Author_id, topic.Title, topic.Content, time.Now().Unix(), topic.Author_name)
	id, _ := res.LastInsertId()
	return int(id), err
}

func PostComment(cmt Comment) error {

	fmt.Println("!!!!")
	stmt, err := db.Prepare(`INSERT INTO comments(tid ,author_id,content, date,author_name)
		values (?,?,?,?,?);`)

	if err != nil {
		fmt.Println(err.Error())
	}
	cmt.Date = int(time.Now().Unix())
	defer stmt.Close()
	_, err = stmt.Exec(cmt.Tid, cmt.Author_id, cmt.Content, cmt.Date, cmt.Author_name)
	fans, _ := getFans(cmt.Author_id)
	replyList := FindReplyToWho(cmt.Content)
	message := jsonfyComment(cmt)
	for i := 0; i < len(fans); i++ {

		fanId, _ := strconv.Atoi(fans[i])
		pushMessage(fanId, message, getMessageSlotKey)

	}
	for i := 0; i < len(replyList); i++ {

		pushMessage(replyList[i], message, getReplySlotKey)

	}

	return err
}

func FindReplyToWho(content string) []int {
	reg := regexp.MustCompile(`@[^\s]+`)
	res := reg.FindAllString(content, -1)
	if len(res) == 0 {
		return []int{}
	}
	var tmp []string
	log.Println(res)
	for i := 0; i < len(res); i++ {
		tmp = append(tmp, "?")
	}
	placeHold := strings.Join(tmp, ",")

	queryString := fmt.Sprintf("select uid from users where username in (%s)", placeHold)
	var params []interface{}
	for i := 0; i < len(tmp); i++ {
		params = append(params, res[i][1:])
	}
	fmt.Println(params)
	rows, err := db.Query(queryString, params...)

	log.Println(queryString)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return []int{}
	}
	uids := []int{}

	for rows.Next() {
		var uid int
		rows.Scan(&uid)
		uids = append(uids, uid)
	}
	return uids
}

func jsonfyComment(cmt Comment) string {
	res, _ := json.Marshal(cmt)
	return string(res)
}

func GetCommentsOfTopic(tid int) ([]Comment, error) {
	res := []Comment{}
	rows, err := db.Query("select cid ,tid, author_id, content, up, down, date, author_name from comments where tid=?", tid)
	defer rows.Close()
	log.Println("++++")

	for rows.Next() {
		var author_name, content sql.NullString
		var cid, tid, author_id, up, down, date sql.NullInt64

		err = rows.Scan(&cid, &tid, &author_id, &content, &up, &down, &date, &author_name)
		if err != nil {
			return res, err
		}
		res = append(res, Comment{
			Cid:         int(cid.Int64),
			Tid:         int(tid.Int64),
			Author_id:   int(author_id.Int64),
			Content:     content.String,
			Up:          int(up.Int64),
			Down:        int(up.Int64),
			Date:        int(date.Int64),
			Author_name: author_name.String,
		})
	}
	return res, err
}

func GetTopic(Tid int) (*Topic, error) {

	row := db.QueryRow("select tid,author_id, title, content, up, down, date, author_name from topics where tid=? order by date desc", Tid)

	var author_name, content, title sql.NullString
	var tid, author_id, up, down, date sql.NullInt64

	err := row.Scan(&tid, &author_id, &title, &content, &up, &down, &date, &author_name)
	if err != nil {
		return nil, err
	}
	incrScore(Tid, "hot")
	return &Topic{
		Tid:         int(tid.Int64),
		Author_id:   int(author_id.Int64),
		Title:       title.String,
		Content:     content.String,
		Up:          int(up.Int64),
		Down:        int(down.Int64),
		Date:        int(date.Int64),
		Author_name: author_name.String,
	}, err

}

func GetCommentsOfAuthor(author_id int) ([]Comment, error) {
	res := []Comment{}
	rows, err := db.Query("select cid ,tid, author_id, content, up, down, date, author_name from comments where author_id=? order by date desc", author_id)
	defer rows.Close()
	log.Println("++++")

	for rows.Next() {
		var author_name, content sql.NullString
		var cid, tid, author_id, up, down, date sql.NullInt64

		err = rows.Scan(&cid, &tid, &author_id, &content, &up, &down, &date, &author_name)
		if err != nil {
			return res, err
		}
		res = append(res, Comment{
			Cid:         int(cid.Int64),
			Tid:         int(tid.Int64),
			Author_id:   int(author_id.Int64),
			Content:     content.String,
			Up:          int(up.Int64),
			Down:        int(up.Int64),
			Date:        int(date.Int64),
			Author_name: author_name.String,
		})
	}
	return res, err
}

func GetAllTopic(offset, limit_num int) ([]Topic, error) {
	res := []Topic{}
	rows, err := db.Query("select tid, author_id,title, content, up, down, date, author_name from  topics where 1 order by date desc limit ?,?", offset, limit_num)
	log.Printf("[+] offset:%v  limit:%v", offset, limit_num)
	defer rows.Close()
	log.Println("++++")
	for rows.Next() {
		var author_name, content, title sql.NullString
		var tid, author_id, up, down, date sql.NullInt64
		err := rows.Scan(&tid, &author_id, &title, &content, &up, &down, &date, &author_name)
		if err != nil {
			return res, err
		}
		res = append(res, Topic{
			Tid:         int(tid.Int64),
			Author_id:   int(author_id.Int64),
			Title:       title.String,
			Content:     content.String,
			Up:          int(up.Int64),
			Down:        int(down.Int64),
			Date:        int(date.Int64),
			Author_name: author_name.String,
		})
	}
	return res, err

}

func GetAllTopicContains(target string, offset, limit_num int) ([]Topic, error) {
	res := []Topic{}
	target = "%" + target + "%"
	log.Println(target)
	rows, err := db.Query(`select tid, author_id,title, content, up, down, date, author_name from  topics where  content like ? or title like ? order by date desc limit ?,?`, target, target, offset, limit_num)
	log.Printf("[+] offset:%v  limit:%v", offset, limit_num)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	log.Println("++++")
	for rows.Next() {
		var author_name, content, title sql.NullString
		var tid, author_id, up, down, date sql.NullInt64
		err := rows.Scan(&tid, &author_id, &title, &content, &up, &down, &date, &author_name)
		if err != nil {
			return res, err
		}
		res = append(res, Topic{
			Tid:         int(tid.Int64),
			Author_id:   int(author_id.Int64),
			Title:       title.String,
			Content:     content.String,
			Up:          int(up.Int64),
			Down:        int(down.Int64),
			Date:        int(date.Int64),
			Author_name: author_name.String,
		})
	}
	return res, err
}
