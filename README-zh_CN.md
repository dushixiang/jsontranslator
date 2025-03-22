# JSON Translator

[English](./README.md) | 简体中文

[![Go Report Card](https://goreportcard.com/badge/github.com/dushixiang/jsontranslator)](https://goreportcard.com/report/github.com/dushixiang/jsontranslator)

专业的 JSON 文件翻译工具，基于 Go 语言开发，支持大语言模型（LLM）的多语言翻译，保持 JSON 结构完整性。

## 功能特性

- 🚀 多语言互译（支持 80+ 语言）
- 🧩 完整保留 JSON 结构和键序
- 🔧 可配置的 LLM 提供商（OpenAI/Custom）
- 🔄 嵌套对象递归翻译
- 📁 批量文件处理支持
- ✅ 增量翻译模式

## 快速开始

### 安装

```bash
# 克隆仓库
git clone https://github.com/dushixiang/jsontranslator.git
cd jsontranslator

# 编译项目
go build -o jsont ./cmd/jsont
```

### 基本用法

1. 创建配置文件 `config.json`：

    ```json
    {
      "llm": {
        "provider": "custom",
        "base_url": "http://localhost:1234/v1",
        "api_key": "",
        "model": "gemma-3-12b-it"
      },
      "from": "zh-CN",
      "to": [
        "zh-TW",
        "en-US",
        "ja-JP"
      ]
    }
    ```

2. 运行翻译任务：

    ```shell
    ./jsont -c config.json
    ```

### 配置说明

| 参数       | 类型       | 必须 | 说明                      | 
|----------|----------|----|-------------------------| 
| provider | string   | 是  | LLM 提供商 (openai/custom) | 
| base_url | string   | 是  | LLM 服务端点                | 
| api_key  | string   | 否  | 提供商 API 密钥              | 
| model    | string   | 是  | 使用的模型名称                 | 
| from     | string   | 是  | 源语言代码 (如 zh-CN)         | 
| to       | []string | 是  | 目标语言代码列表                |

### 支持语言

完整列表请参考 [languages.go](languages.go)，主要支持：

* 简体中文 (zh-CN)
* 繁体中文 (zh-TW)
* 英语 (en-US)
* 日语 (ja-JP)
* 韩语 (ko-KR)
* 法语 (fr-FR)
* 西班牙语 (es-ES)