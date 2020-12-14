package day14

import (
	"fmt"
	"strconv"
	"strings"
)

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
