package day25

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "day25/input.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() (int, int) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	keys := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	card, err := strconv.Atoi(keys[0])
	if err != nil {
		panic(fmt.Sprintf("could not turn string '%s' into int: %s", keys[0], err))
	}

	door, err := strconv.Atoi(keys[1])
	if err != nil {
		panic(fmt.Sprintf("could not turn string '%s' into int: %s", keys[1], err))
	}

	return card, door
}
