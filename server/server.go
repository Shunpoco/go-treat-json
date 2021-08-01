package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", postHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Request struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	log.Println(req)

	req.Text = "hoge"

	json.NewEncoder(w).Encode(req)
}
