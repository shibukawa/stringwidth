package stringwidth

import (
	"golang.org/x/text/width"
)

//go:generate go run cmd/generator/main.go -pkg stringwidth -o emoji_sequence.go

type Opt struct {
	IsAmbiguousWide bool
}

// Calc calculates string display width on terminal.
//
func Calc(src string, opt ...Opt) int {
	ambiguousWidth := 1
	if len(opt) > 0 && opt[0].IsAmbiguousWide {
		ambiguousWidth = 2
	}
	runes := []rune(src)
	size := 0
	// 0, 2
	for i := 0; i < len(runes); i++ {
		match := false
		if sequences, ok := emojiSequences[runes[i]]; ok {
			for _, sequence := range sequences {
				if len(sequence)+i+1 > len(runes) {
					continue
				}
				notMatch := false
				sequenceLength := len(sequence)
				for j, c := range sequence {
					if runes[i+1+j] != c {
						notMatch = true
						break
					}
				}
				if !notMatch {
					size += 2
					i += sequenceLength
					match = true
					break
				}
			}
			if match {
				continue
			}
		}
		switch width.LookupRune(runes[i]).Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			size += 2
		case width.EastAsianAmbiguous:
			size += ambiguousWidth
		default:
			size += 1
		}
	}
	return size
}
