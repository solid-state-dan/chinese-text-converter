package processor_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/solid-state-dan/chinese-text-converter/pkg/converter"
	"github.com/solid-state-dan/chinese-text-converter/pkg/dictionary"
	"github.com/solid-state-dan/chinese-text-converter/pkg/processor"
)

func TestProcessSRTDualStream(t *testing.T) {
	// Mock dictionary for test isolated execution
	mockDict := dictionary.Dictionary{
		"你好": "ni3 hao3",
		"世界": "shi4 jie4",
		"今天": "jin1 tian1",
		"天气": "tian1 qi4",
		"很":  "hen3",
		"好":  "hao3",
	}

	conv := converter.New(mockDict)
	proc := processor.New(conv)

	inputSRT := `1
00:00:01,000 --> 00:00:03,000
你好世界！

2
00:00:04,000 --> 00:00:07,000
今天天气很好。
`

	expectedSRT := `1
00:00:01,000 --> 00:00:03,000
你好世界！
ni3 hao3 shi4 jie4 ！

2
00:00:04,000 --> 00:00:07,000
今天天气很好。
jin1 tian1 tian1 qi4 hen3 hao3 。
`

	reader := strings.NewReader(inputSRT)
	var outputBuffer bytes.Buffer

	err := proc.ProcessSRTDualStream(reader, &outputBuffer)
	if err != nil {
		t.Fatalf("Unexpected error during stream processing: %v", err)
	}

	got := outputBuffer.String()
	if got != expectedSRT {
		t.Errorf("\n--- GOT ---\n%s\n--- EXPECTED ---\n%s", got, expectedSRT)
	}
}
