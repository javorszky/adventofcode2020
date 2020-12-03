package main

import (
	"io/ioutil"
	"strings"
)

const filename = "input.txt"

func main() {

}

// getInputs reads the input.txt file, and arranges the contents into a map of unique numbers.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}
