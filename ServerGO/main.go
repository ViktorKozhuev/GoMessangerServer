package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Message struct{
	UserName string
	MessageText string
	TimeStamp string
}


func main(){
	r := chi.NewRouter()
	fmt.Println("Service is running")
	messages := make([]Message, 0)

	r.Post("/api/Messanger", func(w http.ResponseWriter, r *http.Request) {
		mes := Message{}
		raw, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(raw, &mes)
		fmt.Println(mes)
		messages = append(messages, mes)

	})

	r.Get("/api/Messanger/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		id_num, _ := strconv.Atoi(id)
		if id_num >= 0 && id_num < len(messages) {
			msg := messages[id_num]
			data, _ := json.Marshal(msg)
			w.Write(data)
		}
	})

	http.ListenAndServe(":5000", r)
}