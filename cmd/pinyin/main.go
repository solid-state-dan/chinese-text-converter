package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/solid-state-dan/chinese-text-converter/pkg/converter"
	"github.com/solid-state-dan/chinese-text-converter/pkg/dictionary"
	"github.com/solid-state-dan/chinese-text-converter/pkg/processor"
)

func main() {
	// 1. Define command-line flags
	inputPath := flag.String("i", "", "Path to input .srt file (required)")
	outputPath := flag.String("o", "", "Path to output .srt file (required)")
	dictPath := flag.String("dict", "data/cedict_ts.u8", "Path to CC-CEDICT dictionary file")

	// Custom usage message when running -help or passing invalid flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -i <input.srt> -o <output.srt> [-dict <cedict.u8>]\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// 2. Validate required flags
	if *inputPath == "" || *outputPath == "" {
		fmt.Fprintln(os.Stderr, "Error: Both -i and -o flags are required.")
		flag.Usage()
		os.Exit(1)
	}

	// 3. Load dictionary
	fmt.Printf("Loading dictionary from %s...\n", *dictPath)
	start := time.Now()

	dict, err := dictionary.Load(*dictPath)
	if err != nil {
		log.Fatalf("Failed to load dictionary: %v", err)
	}
	fmt.Printf("Loaded dictionary in %v.\n", time.Since(start))

	// 4. Initialize converter and processor
	conv := converter.New(dict)
	proc := processor.New(conv)

	// 5. Process SRT file
	fmt.Printf("Processing %s -> %s...\n", *inputPath, *outputPath)
	processStart := time.Now()

	err = proc.ProcessSRTDual(*inputPath, *outputPath)
	if err != nil {
		log.Fatalf("Processing failed: %v", err)
	}

	fmt.Printf("Successfully generated dual-line subtitles in %v!\n", time.Since(processStart))
}
