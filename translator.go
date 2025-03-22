package jsontranslator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type Provider string

const (
	OpenAI Provider = "openai"
	Custom Provider = "custom"
)

type LargeLanguageModelConfig struct {
	Provider Provider `json:"provider"`
	BaseURL  string   `json:"base_url"`
	ApiKey   string   `json:"api_key"`
	Model    string   `json:"model"`
}

func TranslateWithLLM(str, from, to string, llmConfig LargeLanguageModelConfig) (string, error) {
	config := openai.DefaultConfig(llmConfig.ApiKey)

	var (
		fromKey = getLanguageKeyFromValue(from, AILanguageMap)
		toKey   = getLanguageKeyFromValue(to, AILanguageMap)
	)
	switch llmConfig.Provider {
	case Custom:
		config.BaseURL = llmConfig.BaseURL
	case OpenAI:

	}

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: llmConfig.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: `You are a translation assistant. Translate any text given to you into the specified language. Do not return anything else.`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(`%s to %s: "%s"`, fromKey, toKey, str),
				},
			},
			MaxTokens: 1024,
		},
	)

	if err != nil {
		return "", fmt.Errorf("chat completion request: %w", err)
	}

	content := resp.Choices[0].Message.Content
	content = strings.TrimLeft(content, `"`)
	content = strings.TrimRight(content, `"`)
	content = strings.TrimLeft(content, `「`)
	content = strings.TrimRight(content, `」`)
	content = strings.TrimRight(content, "\n")
	return content, nil
}

type Translator struct {
	input          string
	output         string
	from           string
	to             string
	providerConfig LargeLanguageModelConfig
}

func (c Translator) Translate(str, from, to string) (string, error) {
	return TranslateWithLLM(str, from, to, c.providerConfig)
}

func (c Translator) RecursiveTranslate(data *JsonLanguage) (*JsonLanguage, error) {
	var translatedMap = NewJsonLanguage()
	// 读取目标语言文件
	file, err := os.ReadFile(c.output)
	if err == nil {
		err = translatedMap.UnmarshalJSON(file)
		if err != nil {
			translatedMap = NewJsonLanguage()
		}
	}
	c.recursiveTranslate(data, translatedMap)
	// 排序
	return translatedMap, nil
}

func (c Translator) recursiveTranslate(data any, translatedMap *JsonLanguage) {
	switch v := data.(type) {
	case *JsonLanguage:
		for key, value := range v.AllFromFront() {
			switch vType := value.(type) {
			case string:
				oldValue, ok := translatedMap.Get(key)
				if ok {
					translatedMap.Set(key, oldValue)
					continue
				}
				content, err := c.Translate(vType, c.from, c.to)
				if err != nil {
					fmt.Printf("Error translating %s from %s to %s: %v\n", value, c.from, c.to, err.Error())
					continue
				}
				fmt.Printf("Translated %s from %s to %s: %s\n", value, c.from, c.to, content)
				translatedMap.Set(key, content)
			case *JsonLanguage:
				inner, ok := translatedMap.Get(key)
				if !ok {
					inner = NewJsonLanguage()
				}
				innerMap, ok := inner.(*JsonLanguage)
				if !ok {
					inner = NewJsonLanguage()
				}
				c.recursiveTranslate(vType, innerMap)
				translatedMap.Set(key, inner)
			}
		}
	}
}

func (c Translator) Translates() error {
	// 读取 JSON 文件
	file, err := os.ReadFile(c.input)
	if err != nil {
		return fmt.Errorf("read json file: %w", err)
	}

	jsonLanguage := NewJsonLanguage()
	// 解析 JSON 数据
	err = jsonLanguage.UnmarshalJSON(file)
	if err != nil {
		return fmt.Errorf("parsing json err: %w", err)
	}

	// 递归翻译 JSON
	translatedData, err := c.RecursiveTranslate(jsonLanguage)
	if err != nil {
		return err
	}

	// 将翻译后的数据转换为JSON字符串
	output, err := json.MarshalIndent(translatedData, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling json err: %v", err)
	}
	// 将翻译后的数据写入输出文件
	err = os.WriteFile(c.output, output, 0644)
	if err != nil {
		return fmt.Errorf("writing json file: %w", err)
	}
	return nil
}

func New(input, output, from, to string, providerConfig LargeLanguageModelConfig) *Translator {
	return &Translator{
		input:          input,
		output:         output,
		from:           from,
		to:             to,
		providerConfig: providerConfig,
	}
}
