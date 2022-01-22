package wordcount

import (
	"encoding/json"
	"net/http"
)

type WordCount struct {
	Word  string `json:"word"`
	Count uint32 `json:"count"`
}

type request struct {
	Text string `json:"text"`
}

type response struct {
	Result []WordCount `json:"result"`
}

func wordCount(w http.ResponseWriter, req *http.Request) {
	var r request

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	return
}

func main() {
	http.HandleFunc("/count", wordCount)

	http.ListenAndServe(":8090", nil)
}
