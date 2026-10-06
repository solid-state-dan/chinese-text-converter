package main

import (
	"fmt"
	"log"
	"time"

	"github.com/solid-state-dan/chinese-text-converter/pkg/converter"
	"github.com/solid-state-dan/chinese-text-converter/pkg/dictionary"
)

func main() {
	fmt.Println("Loading CC-CEDICT dictionary...")
	start := time.Now()

	dict, err := dictionary.Load("data/cedict_ts.u8")
	if err != nil {
		log.Fatalf("Failed to load dictionary: %v", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Successfully loaded %d entries in %v!\n\n", len(dict), elapsed)

	conv := converter.New(dict)

	testSentences := []string{
		"你好世界",
		"我爱吃北京烤鸭",
		"今天天气很好，我去银行",
		"Hello World! 这是一个测试 123.",
	}

	start = time.Now()
	for _, sentence := range testSentences {
		result := conv.ToPinyin(sentence)
		fmt.Printf("Input:  %s\n", sentence)
		fmt.Printf("Pinyin: %s\n\n", result)
	}
	elapsed = time.Since(start)
	fmt.Printf("Successfully converted entries in %v!\n\n", elapsed)
}
