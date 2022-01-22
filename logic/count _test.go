package logic_test

import (
	"testing"

	"wordcount/logic"
)

func TestCount(t *testing.T) {
	expectedCount := 3
	req := []string{
		"lorem",
		"ipsum",
		"test",
		"lorem",
		"lorem",
		"ipsum",
	}

	res := logic.Count(req)
	if len(res) != expectedCount {
		t.Fatalf("expected len of map to be %d", expectedCount)
	}

	if res["lorem"] != 3 {
		t.Fatalf("expected lorem to have 3 count")
	}
}

func TestSort(t *testing.T) {
	expectedCount := 3
	req := map[string]uint32{
		"test":  1,
		"lorem": 3,
		"ipsum": 2,
	}

	res := logic.Sort(req)
	if len(res) != expectedCount {
		t.Fatalf("expected len of wordCount slice to be %d", expectedCount)
	}

	expectedFirst := logic.WordCount{Word: "lorem", Count: 3}
	if res[0] != expectedFirst {
		t.Fatalf("expected 1st slice to be %+v", expectedFirst)
	}

	expectedSecond := logic.WordCount{Word: "ipsum", Count: 2}
	if res[1] != expectedSecond {
		t.Fatalf("expected 2nd slice to be %+v", expectedSecond)
	}

	expectedLast := logic.WordCount{Word: "test", Count: 1}
	if res[2] != expectedLast {
		t.Fatalf("expected 2nd slice to be %+v", expectedLast)
	}
}
