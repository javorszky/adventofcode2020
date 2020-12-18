package day18

import (
	"io/ioutil"
	"strings"
)

const filename = "day18/input.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	spaceReplacer := strings.NewReplacer(" ", "")
	return strings.Split(spaceReplacer.Replace(strings.TrimRight(string(data), "\n")), "\n")
}
