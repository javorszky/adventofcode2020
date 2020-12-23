package day23

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "day23/input.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	trimmed := strings.TrimRight(string(data), "\n")
	out := make([]int, 0)
	for _, n := range strings.Split(trimmed, "") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("cant turn string '%s' into int: %s", n, err))
		}
		out = append(out, i)
	}
	return out
}
