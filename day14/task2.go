package day14

import (
	"fmt"
	"strings"
)

func task2() {
	_ = getInputs()
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
