package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reMem = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

func task1() {
	input := getInputs()
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	memory := make(map[int64]int64, 0)
	for _, line := range input {
		switch line[:3] {
		case "mas":
			mask = line[7:]
		case "mem":
			matches := reMem.FindStringSubmatch(line)
			address, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				panic(fmt.Sprintf("can't turn string %s into int64: %s", matches[1], err))
			}
			value, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				panic(fmt.Sprintf("can't turn string %s into int64: %s", matches[2], err))
			}

			memory[address] = binaryToDecimal(applyMask(mask, decimalToBinary(value)))
		}
	}
	sum := int64(0)
	for _, value := range memory {
		sum = sum + value
	}
	fmt.Printf("\nDay 14 task 1: the sum of all values in memory is %d\n", sum)
}

func decimalToBinary(i int64) string {
	b := strconv.FormatInt(i, 2)
	switch {
	case len(b) == 36:
		return b
	case len(b) > 36:
		return b[len(b)-36:]
	default:
		return strings.Repeat("0", 36-len(b)) + b
	}

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
