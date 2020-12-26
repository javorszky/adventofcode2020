package day20

import (
	"io/ioutil"
	"strings"
)

const filename = "day20/input.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	return parseInput(getFileContent(fn))
}

func parseInput(data string) []string {
	input := strings.TrimRight(string(data), "\n")
	tiles := strings.Split(input, "\n\n")

	return tiles
}

func getFileContent(s string) string {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func remove(slice []tile, s int) []tile {
	return append(slice[:s], slice[s+1:]...)
}
