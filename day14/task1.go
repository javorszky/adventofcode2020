package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func task1() {
	_ = getInputs()
}

func decimalToBinary(i int64) string {
	b := strconv.FormatInt(i, 2)

	return b[len(b)-36:]
}

func binaryToDecimal(s string) int64 {
	d, err := strconv.ParseInt(s[len(s)-36:], 2, 64)
	if err != nil {
		panic(fmt.Sprintf("converting %s to binary failed: %s", s, err))
	}
	return d
}

func applyMask(mask, binary string) string {
	if len(mask) != len(binary) {
		panic(fmt.Sprintf("mask length (%d) and binary lengths (%d) differ in applying mask", len(mask), len(binary)))
	}
	var sb strings.Builder
	for idx, char := range mask {
		switch string(char) {
		case "1":
			sb.WriteString("1")
		case "0":
			sb.WriteString("0")
		default:
			sb.WriteString(string(binary[idx]))
		}
	}
	return sb.String()
}
