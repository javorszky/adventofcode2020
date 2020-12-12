package day12

import (
	"io/ioutil"
	"strings"
)

var rotMap = map[int]int{
	90:  270,
	180: 180,
	270: 90,
}

type mover interface {
	posX() int
	posY() int
}

func manhattanDistance(s mover) int {
	return abs(s.posX()) + abs(s.posY())
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
