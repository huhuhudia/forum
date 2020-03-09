package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/user", dispatch_user)
	http.HandleFunc("/", handle_index)
	http.ListenAndServe(":8888", nil)

}
func dispatch_user(w http.ResponseWriter, r *http.Request) {
	log.Println("enter dispath")
	var u User
	var input []byte
	err := json.Unmarshal(input, u)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = r.Body.Read(input)
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:
		GetUser(u)
	case http.MethodPost:
		PostUser(u)
	default:
		w.Write([]byte("not support"))
	}
}

func handle_index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "123")
}
