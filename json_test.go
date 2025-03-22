package jsontranslator

import (
	"os"
	"sort"
	"testing"
)

func TestSortJson(t *testing.T) {
	jsonLanguage := NewJsonLanguage()
	file, err := os.ReadFile("./testdata/zh-CN.json")
	if err != nil {
		t.Fatal(err)
	}
	err = jsonLanguage.UnmarshalJSON(file)
	if err != nil {
		t.Fatal(err)
	}

	var keys []string
	for key := range jsonLanguage.Keys() {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		t.Log(key)
	}
}
