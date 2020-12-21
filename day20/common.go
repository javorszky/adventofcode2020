package day20

import (
	"io/ioutil"
	"strings"
)

const filename = "day20/input_example.txt"

//const filename = "day20/input.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := strings.TrimRight(string(data), "\n")
	tiles := strings.Split(input, "\n\n")

	return tiles
}

func remove(slice []tile, s int) []tile {
	return append(slice[:s], slice[s+1:]...)
}
