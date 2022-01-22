package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"wordcount/logic"
)

var (
	WordRe = regexp.MustCompile(`\w+`)
)

type request struct {
	Text string `json:"text"`
}

type response struct {
	Result []logic.WordCount `json:"result"`
}

func wordCount(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "only method POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var r request
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// match whole words, removes any punctuation/whitespace
	words := WordRe.FindAllString(strings.ToLower(r.Text), -1)

	sorted := logic.CountAndLimitSort(words)

	res, err := json.Marshal(response{Result: sorted})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

	return
}

func main() {
	http.HandleFunc("/count", wordCount)

	http.ListenAndServe(":8090", nil)
}
