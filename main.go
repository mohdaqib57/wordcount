package wordcount

import (
	"encoding/json"
	"net/http"
	"regexp"

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
