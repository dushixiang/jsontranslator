package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/dushixiang/jsontranslator"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "config.json", "config file")
	flag.Parse()

	log.SetFlags(log.Lshortfile)
}

func main() {
	fmt.Println("json translator version:", jsontranslator.VERSION)
	config, err := jsontranslator.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Dir(config.Input)
	for _, to := range config.To {
		fmt.Println("------------------------------------------")
		fmt.Printf("Translating: %s -> %s\n", config.From, to)
		translator := jsontranslator.New(config.Input,
			filepath.Join(dir, fmt.Sprintf("%s.json", to)),
			config.From,
			to,
			config.LLM,
		)
		err = translator.Translates()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Translated: %s -> %s\n", config.From, to)
	}
}
