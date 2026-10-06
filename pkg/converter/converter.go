package converter

import (
	"strings"

	"github.com/solid-state-dan/chinese-text-converter/pkg/dictionary"
)

// DefaultMaxWordLength sets a reasonable ceiling for word matching windows.
const DefaultMaxWordLength = 5

// Converter handles text segmentation and transformation using a dictionary.
type Converter struct {
	dict          dictionary.Dictionary
	maxWordLength int
}

// New creates a Converter instance wrapping a loaded dictionary.
func New(dict dictionary.Dictionary) *Converter {
	return &Converter{
		dict:          dict,
		maxWordLength: DefaultMaxWordLength,
	}
}

// ToPinyin converts a Hanzi string into spaced, numbered Pinyin.
func (c *Converter) ToPinyin(text string) string {
	runes := []rune(text)
	textLength := len(runes)

	if textLength == 0 {
		return ""
	}

	var pinyinTokens []string
	cursor := 0

	for cursor < textLength {
		// Calculate the starting window boundary
		windowEnd := cursor + c.maxWordLength
		if windowEnd > textLength {
			windowEnd = textLength
		}

		matched := false

		// Shrink window from max size down to single rune
		for end := windowEnd; end > cursor; end-- {
			subStr := string(runes[cursor:end])

			if pinyin, found := c.dict[subStr]; found {
				pinyinTokens = append(pinyinTokens, pinyin)
				cursor = end // Advance past the matched word
				matched = true
				break
			}
		}

		// Fallback for non-dictionary characters (punctuation, numbers, English)
		if !matched {
			pinyinTokens = append(pinyinTokens, string(runes[cursor]))
			cursor++
		}
	}

	return strings.Join(pinyinTokens, " ")
}
