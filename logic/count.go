package logic

import "sort"

type WordCount struct {
	Word  string `json:"word"`
	Count uint32 `json:"count"`
}

func Count(words []string) map[string]uint32 {
	wordFreq := make(map[string]uint32)
	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}

// https://stackoverflow.com/questions/18695346/how-can-i-sort-a-mapstringint-by-its-values/44380276#44380276
func Sort(req map[string]uint32) []WordCount {
	ss := make([]WordCount, 0)
	for k, v := range req {
		ss = append(ss, WordCount{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Count > ss[j].Count
	})

	return ss
}
