package jsontranslator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateWithLLM(t *testing.T) {
	str := "Hello"
	from := "en"
	to := "zh-CN"

	translated, err := TranslateWithLLM(str, from, to, LargeLanguageModelConfig{
		Provider: Custom,
		Model:    "qwen2.5-7b-instruct-1m",
		BaseURL:  "http://localhost:1234/v1",
		ApiKey:   "",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, translated)

	assert.Equal(t, "你好", translated)
}

func TestTranslates(t *testing.T) {
	var tos = []string{
		"zh-TW",
		"en-US",
		"ja-JP",
	}
	for _, to := range tos {
		translator := New("./testdata/zh-CN.json",
			fmt.Sprintf("./testdata/%s.json", to),
			"zh-CN",
			to,
			LargeLanguageModelConfig{
				Provider: Custom,
				//Model:     "qwen2.5-7b-instruct-1m",
				Model:   "gemma-3-12b-it",
				BaseURL: "http://localhost:1234/v1",
				ApiKey:  "",
			},
		)
		err := translator.Translates()
		assert.NoError(t, err)
	}

}
