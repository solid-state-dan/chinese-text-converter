package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/solid-state-dan/chinese-text-converter/pkg/converter"
	"github.com/solid-state-dan/chinese-text-converter/pkg/dictionary"
	"github.com/solid-state-dan/chinese-text-converter/pkg/processor"
)

func main() {
	// 1. Load Dictionary
	fmt.Println("Loading CC-CEDICT dictionary...")
	start := time.Now()

	dict, err := dictionary.Load("data/cedict_ts.u8")
	if err != nil {
		log.Fatalf("Failed to load dictionary: %v", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Successfully loaded %d entries in %v!\n\n", len(dict), elapsed)

	conv := converter.New(dict)
	proc := processor.New(conv)

	// 2. Create a small test .srt file
	sampleSRT := `1
00:00:01,000 --> 00:00:03,000
你好，

2
00:00:04,000 --> 00:00:07,000
世界！
`
	inputPath := "data/sample.srt"
	outputPath := "data/sample_pinyin.srt"

	os.WriteFile(inputPath, []byte(sampleSRT), 0644)

	// 3. Process File
	start = time.Now()
	err = proc.ProcessSRTDual(inputPath, outputPath)
	if err != nil {
		log.Fatalf("Processing failed: %v", err)
	}

	fmt.Printf("Processed %s -> %s in %v!\n\n", inputPath, outputPath, time.Since(start))

	// 4. Print generated file contents to terminal
	outputContent, _ := os.ReadFile(outputPath)
	fmt.Println("Generated SRT Output:")
	fmt.Println(string(outputContent))
}
