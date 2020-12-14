package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func task2() {
	input := getInputs()
	mask := "000000000000000000000000000000000000"
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

			for _, add := range extrapolateAddresses("", applyMaskT2(mask, decimalToBinary(address))) {
				extrapolatedAddress := binaryToDecimal(add)
				memory[extrapolatedAddress] = value
			}
		}
	}
	sum := int64(0)
	for _, value := range memory {
		sum = sum + value
	}
	fmt.Printf("Day 14 task 2: the sum of all values in memory is %d\n", sum)
}

func applyMaskT2(mask, binary string) string {
	if len(mask) != len(binary) {
		panic(fmt.Sprintf("mask length (%d) and binary lengths (%d) differ in applying mask", len(mask), len(binary)))
	}
	var sb strings.Builder
	for idx, char := range mask {
		switch string(char) {
		case "1":
			sb.WriteString("1")
		case "X":
			sb.WriteString("X")
		default:
			sb.WriteString(string(binary[idx]))
		}
	}
	return sb.String()
}

// extrapolateAddresses takes a masked string, and produces all possible values from it.
func extrapolateAddresses(prefix, s string) []string {
	values := make([]string, 0)
	var sb strings.Builder

	sb.WriteString(prefix)

	for idx, char := range s {
		c := string(char)
		switch c {
		case "X":
			values = append(values, extrapolateAddresses(sb.String()+"0", s[idx+1:])...)
			values = append(values, extrapolateAddresses(sb.String()+"1", s[idx+1:])...)
			return values
		default:
			sb.WriteString(c)

		}
	}
	values = append(values, sb.String())
	return values
}
