package main

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/user", dispatch_user)
	http.HandleFunc("/api/fans", handleFans)
	http.HandleFunc("/api/follow", handleFollow)
	http.ListenAndServe(":8888", nil)

}

type RelationReq struct {
	Method  string `json:"method"`
	SelfID  int    `json:"self_id"`
	OtherID int    `json:"other_id"`
}

func handleFans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("enter dispath")
	var params RelationReq
	// var input []byte

	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &params)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch params.Method {
	case http.MethodGet:
		res, err := getFans(params.SelfID)
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		jsonstr, err := json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonstr)
		return
	case http.MethodPost:
		addFans(params.SelfID, params.OtherID)
		w.WriteHeader(http.StatusOK)
		return
	}

}
func handleFollow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("enter dispath")
	var params RelationReq
	// var input []byte

	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &params)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch params.Method {
	case http.MethodGet:
		res, err := getFollow(params.SelfID)
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		jsonstr, err := json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonstr)
		return
	case http.MethodPost:
		addFollow(params.SelfID, params.OtherID)
		w.WriteHeader(http.StatusOK)
		return
	}

}

func dispatch_user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("enter dispath")
	var u User
	// var input []byte

	defer r.Body.Close()
	s, n := ioutil.ReadAll(r.Body)
	log.Println("s : ", s, "len :", n)
	err := json.Unmarshal([]byte(s), &u)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch u.Method {
	case http.MethodGet:
		res := GetUser(u)
		if res == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		resp_json, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp_json)
	case http.MethodPost:
		id, err := PostUser(u)
		u.Uid = id
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		resp_json, err := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Write(resp_json)

	default:
		w.Write([]byte("not support"))
	}
}
