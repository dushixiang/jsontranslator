# JSON Translator

English | [简体中文](./README-zh_CN.md)

[![Go Report Card](https://goreportcard.com/badge/github.com/dushixiang/jsontranslator)](https://goreportcard.com/report/github.com/dushixiang/jsontranslator)

A professional JSON file translation tool developed in Go, supporting multilingual translation with large language models (LLM) while maintaining JSON structure integrity.

## Features

- 🚀 Multilingual translation (supports 80+ languages)
- 🧩 Preserves JSON structure and key order
- 🔧 Configurable LLM provider (OpenAI/Custom)
- 🔄 Recursive translation for nested objects
- 📁 Batch file processing support
- ✅ Incremental translation mode

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/dushixiang/jsontranslator.git
cd jsontranslator

# Build the project
go build -o jsont ./cmd/jsont
```

### Basic Usage

1. Create a configuration file `config.json`:

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

2. Run the translation task:

    ```shell
    ./jsont -c config.json
    ```

### Configuration

| Parameter  | Type      | Required | Description                     |
|------------|----------|----------|---------------------------------|
| provider   | string   | Yes      | LLM provider (openai/custom)   |
| base_url   | string   | Yes      | LLM service endpoint           |
| api_key    | string   | No       | API key for the provider       |
| model      | string   | Yes      | Model name                     |
| from       | string   | Yes      | Source language code (e.g., zh-CN) |
| to         | []string | Yes      | List of target language codes  |

### Supported Languages

For a complete list, refer to [languages.go](languages.go). Some of the main supported languages include:

* Simplified Chinese (zh-CN)
* Traditional Chinese (zh-TW)
* English (en-US)
* Japanese (ja-JP)
* Korean (ko-KR)
* French (fr-FR)
* Spanish (es-ES)