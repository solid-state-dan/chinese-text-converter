package dictionary

import (
	"bufio"
	"os"
	"strings"
)

// Dictionary holds the Simplified Hanzi to Pinyin lookup map.
type Dictionary map[string]string

// Load parses a CC-CEDICT formatted file at the given filePath
// and returns an initialized Dictionary map.
func Load(filePath string) (Dictionary, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dict := make(Dictionary)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// 1. Skip header lines and comments.
		if strings.HasPrefix(line, "#") {
			continue
		}

		// 2. Locate structural markers using IndexByte.
		firstSpace := strings.IndexByte(line, ' ')
		openBracket := strings.IndexByte(line, '[')
		closeBracket := strings.IndexByte(line, ']')

		// Validate line boundaries.
		if firstSpace == -1 || openBracket == -1 || closeBracket == -1 ||
			openBracket <= firstSpace || closeBracket <= openBracket {
			continue
		}

		// 3. Slice key components.
		simplified := line[firstSpace+1 : openBracket-1]
		pinyin := line[openBracket+1 : closeBracket]

		// 4. First Match Wins strategy (preserve first entry encountered).
		if _, exists := dict[simplified]; !exists {
			dict[simplified] = pinyin
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dict, nil
}
