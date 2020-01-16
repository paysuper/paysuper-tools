package string

import (
	"strings"
)

const (
	panMaskFirstSymbolsCount = 6
	panMaskLastSymbolsCount  = 4
	panMaskedSymbol          = "*"
)

func MaskBankCardNumber(pan string) string {
	rSymCount := len(pan) - (panMaskFirstSymbolsCount + panMaskLastSymbolsCount)

	return pan[:panMaskFirstSymbolsCount] +
		strings.Repeat(panMaskedSymbol, rSymCount) +
		pan[rSymCount+panMaskFirstSymbolsCount:]
}
