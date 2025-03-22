package jsontranslator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elliotchance/orderedmap/v3"
)

// JsonLanguage 表示一个有序的 JSON 对象
type JsonLanguage struct {
	*orderedmap.OrderedMap[string, any]
}

// NewJsonLanguage 创建一个新的 JsonLanguage 实例
func NewJsonLanguage() *JsonLanguage {
	return &JsonLanguage{
		orderedmap.NewOrderedMap[string, any](),
	}
}

// MarshalJSON 实现 json.Marshaler 接口，将 JsonLanguage 实例转换为 JSON 字节切片
func (m *JsonLanguage) MarshalJSON() ([]byte, error) {
	if m.Len() == 0 {
		return []byte("{}"), nil
	}
	var buf bytes.Buffer
	buf.WriteByte('{')
	encoder := json.NewEncoder(&buf)

	for index, el := 0, m.Front(); el != nil; el = el.Next() {
		if index > 0 {
			buf.WriteByte(',')
		}
		// 编码键
		if err := encoder.Encode(el.Key); err != nil {
			return nil, fmt.Errorf("failed to encode key %s: %w", el.Key, err)
		}
		buf.WriteByte(':')
		// 编码值
		if err := encoder.Encode(el.Value); err != nil {
			return nil, fmt.Errorf("failed to encode value for key %s: %w", el.Key, err)
		}
		index++
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON 实现 json.Unmarshaler 接口，将 JSON 字节切片解析为 JsonLanguage 实例
func (m *JsonLanguage) UnmarshalJSON(data []byte) error {
	if m.OrderedMap == nil {
		m.OrderedMap = orderedmap.NewOrderedMap[string, any]()
	}
	decoder := json.NewDecoder(bytes.NewReader(data))

	// 跳过 '{'
	token, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read first token: %w", err)
	}
	if token != json.Delim('{') {
		return fmt.Errorf("expected '{' but got %v", token)
	}

	return unmarshalAnyOrderedMap(decoder, m)
}

// unmarshalAnyOrderedMap 递归解析 JSON 对象
func unmarshalAnyOrderedMap(decoder *json.Decoder, m *JsonLanguage) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read token: %w", err)
		}
		// 到达对象末尾
		if token == json.Delim('}') {
			return nil
		}
		key, ok := token.(string)
		if !ok {
			return fmt.Errorf("unexpected key type: %T", token)
		}
		token, err = decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read value token for key %s: %w", key, err)
		}
		switch token {
		case json.Delim('{'):
			mm := NewJsonLanguage()
			if err := unmarshalAnyOrderedMap(decoder, mm); err != nil {
				return fmt.Errorf("failed to unmarshal nested object for key %s: %w", key, err)
			}
			m.Set(key, mm)
		case json.Delim('['):
			values, err := unmarshalAnyOrderedMapArray(decoder)
			if err != nil {
				return fmt.Errorf("failed to unmarshal array for key %s: %w", key, err)
			}
			m.Set(key, values)
		default:
			m.Set(key, token)
		}
	}
}

// unmarshalAnyOrderedMapArray 递归解析 JSON 数组
func unmarshalAnyOrderedMapArray(decoder *json.Decoder) ([]any, error) {
	var values []any
	for {
		token, err := decoder.Token()
		if err != nil {
			return values, fmt.Errorf("failed to read array token: %w", err)
		}
		// 到达数组末尾
		if token == json.Delim(']') {
			return values, nil
		}
		switch token {
		case json.Delim('{'):
			mm := NewJsonLanguage()
			if err := unmarshalAnyOrderedMap(decoder, mm); err != nil {
				return nil, fmt.Errorf("failed to unmarshal nested object in array: %w", err)
			}
			values = append(values, mm)
		case json.Delim('['):
			vv, err := unmarshalAnyOrderedMapArray(decoder)
			if err != nil {
				return values, fmt.Errorf("failed to unmarshal nested array in array: %w", err)
			}
			values = append(values, vv)
		default:
			values = append(values, token)
		}
	}
}
