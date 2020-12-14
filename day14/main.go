package day14

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const filename = "day14/input.txt"

var reMem = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

func Tasks() {
	task1()
	task2()
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
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
