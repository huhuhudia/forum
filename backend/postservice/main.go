package main

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/apiv2/topic", handleTopic)
	http.HandleFunc("/apiv2/comment", handleComment)
	http.HandleFunc("/apiv2/mainpage", handleMainPageRequest)
	http.HandleFunc("/apiv2/commentsOfAuthor", handleGetCommentsOfAuthor)
	http.HandleFunc("/apiv2/message", handleMessage)
	http.HandleFunc("/apiv2/reply", handleReply)
	http.HandleFunc("/apiv2/hot", handleHot)
	http.ListenAndServe(":9999", nil)
}
func handleTopic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("enter handleTopic")
	topic := Topic{}

	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &topic)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	switch topic.Method {
	case http.MethodGet:
		res, err := GetTopic(topic.Tid)
		if err != nil {
			log.Println((err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data, _ := json.Marshal(res)

		w.Write(data)
		return
	case http.MethodPost:
		id, err := PostTopic(topic)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		topic.Tid = id
		data, _ := json.Marshal(topic)
		w.Write(data)
		return
	}

}

func handleComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("enter handleTopic")
	cmt := Comment{}

	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &cmt)
	if err != nil {
		log.Println(err)
		return
	}

	switch cmt.Method {
	case http.MethodGet:
		res, err := GetCommentsOfTopic(cmt.Tid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data, _ := json.Marshal(res)

		w.Write(data)
		return
	case http.MethodPost:
		err := PostComment(cmt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)

		data, _ := json.Marshal(cmt)
		w.Write(data)
		return
	}
}

type MainPageReq struct {
	Target string `json:"target"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

func handleMainPageRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("enter handleMainPageReq")
	params := MainPageReq{}
	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	var res []Topic
	if params.Target == "" {
		res, err = GetAllTopic(params.Offset, params.Limit)
	} else if params.Target != "" {

		res, err = GetAllTopicContains(params.Target, params.Offset, params.Limit)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(res)
	w.Write(data)
	return
}

type CommentOfAuthorRequst struct {
	Author_id int `json:"author_id"`
}

func handleGetCommentsOfAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("enter handleTopic")
	params := CommentOfAuthorRequst{}

	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &params)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := GetCommentsOfAuthor(params.Author_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

type MessageReq struct {
	Uid int `json:"uid"`
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("enter handleTopic")
	params := MessageReq{}

	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v", err)
	}
	res, err := getMessage(params.Uid, 18, getMessageSlotKey)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(generateJson(res))

}

func handleReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := MessageReq{}

	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v", err)
	}
	res, err := getMessage(params.Uid, 18, getReplySlotKey)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(generateJson(res))

}

func handleHot(w http.ResponseWriter, r *http.Request) {
	res := getHotestTopic(0, "hot")
	jsondata, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}

func generateJson(src []string) []byte {
	res := `[` + strings.Join(src, ",") + `]`
	return []byte(res)
}
