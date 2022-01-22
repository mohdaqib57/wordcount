package logic

import "sort"

// CountAndLimitSort is to count, sort and limit to only show the top ten most used words
func CountAndLimitSort(req []string) []WordCount {
	s := Sort(Count(req))

	// only get top ten
	if len(s) > 10 {
		res := make([]WordCount, 0)

		for i, v := range s {
			res = append(res, WordCount{Word: v.Word, Count: v.Count})

			if i == 9 {
				break
			}
		}
		return res
	}

	return s
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
