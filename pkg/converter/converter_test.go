package converter_test

import (
	"testing"

	"github.com/solid-state-dan/chinese-text-converter/pkg/converter"
	"github.com/solid-state-dan/chinese-text-converter/pkg/dictionary"
)

func TestToPinyin(t *testing.T) {
	// Create a small, controlled mock dictionary for testing
	mockDict := dictionary.Dictionary{
		"你好": "ni3 hao3",
		"世界": "shi4 jie4",
		"船长": "chuan2 zhang3",
	}

	conv := converter.New(mockDict)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic two-word compound",
			input:    "你好",
			expected: "ni3 hao3",
		},
		{
			name:     "Maximal matching handles polyphonic word correctly",
			input:    "船长",
			expected: "chuan2 zhang3", // If split single-char, "长" would incorrectly yield "chang2"
		},
		{
			name:     "Mixed text with punctuation and ASCII",
			input:    "你好, World! 123",
			expected: "ni3 hao3 ,   W o r l d !   1 2 3", // Non-dict runes pass through
		},
		{
			name:     "Empty string input",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.ToPinyin(tt.input)
			if got != tt.expected {
				t.Errorf("\nInput:    %q\nGot:      %q\nExpected: %q", tt.input, got, tt.expected)
			}
		})
	}
}
