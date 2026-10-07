package processor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/solid-state-dan/chinese-text-converter/pkg/converter"
)

// State tracking for SRT parsing.
type state int

const (
	stateExpectSeq state = iota
	stateExpectTimecode
	stateExpectText
)

// Processor handles streaming file conversion.
type Processor struct {
	conv *converter.Converter
}

// New creates a new file Processor instance.
func New(conv *converter.Converter) *Processor {
	return &Processor{conv: conv}
}

// ProcessSRTDual reads an input file and writes a dual Hanzi + Pinyin SRT file.
func (p *Processor) ProcessSRTDual(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	return p.ProcessSRTDualStream(inputFile, outputFile)
}

// ProcessSRTDualStream accepts generic Reader/Writer streams for testing or file I/O.
func (p *Processor) ProcessSRTDualStream(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriter(w)
	defer writer.Flush() // Ensure all buffered data is written to disk on return.

	currentState := stateExpectSeq

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		switch currentState {
		case stateExpectSeq:
			// 1. Write sequence number as-is.
			writer.WriteString(line)
			writer.WriteString("\n")
			if trimmed != "" {
				currentState = stateExpectTimecode
			}

		case stateExpectTimecode:
			// 2. Write timecode line as-is.
			writer.WriteString(line)
			writer.WriteString("\n")
			currentState = stateExpectText

		case stateExpectText:
			if trimmed == "" {
				// Blank line signals the end of the subtitle block.
				writer.WriteString("\n")
				currentState = stateExpectSeq
			} else {
				// 3. Write original Hanzi line.
				writer.WriteString(line)
				writer.WriteString("\n")
				// 4. Generate and write Pinyin line directly underneath.
				pinyin := p.conv.ToPinyin(line)
				writer.WriteString(pinyin)
				writer.WriteString("\n")
			}
		}
	}

	return scanner.Err()
}
