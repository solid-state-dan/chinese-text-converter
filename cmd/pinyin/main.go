package main

import (
	"fmt"
	"log"
	"time"

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

	// Quick sanity test on lookups
	testKeys := []string{"你好", "银行", "行", "地"}
	for _, key := range testKeys {
		if py, found := dict[key]; found {
			fmt.Printf("  %-6s -> %s\n", key, py)
		} else {
			fmt.Printf("  %-6s -> [NOT FOUND]\n", key)
		}
	}
}
