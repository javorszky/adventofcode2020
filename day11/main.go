package day11

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "day11/input.txt"

type floorIsLave struct {
	Message string
}

// Error returns the string representation of the error.
func (f floorIsLave) Error() string {
	return f.Message
}

type outOfBounds struct {
	Message string
}

// Error returns the string representation of the error.
func (o outOfBounds) Error() string {
	return o.Message
}

func Tasks() {
	task1()
	task2()
}

func task1() {
	_ = getInputs()
}

// t1normalize returns one string and a line length to be used for all the other bits. This assumes that the input slice
// is uniform in its line length.
func t1normalize(input []string) (int, string) {
	lineLength := len(input[0])
	var sb strings.Builder
	for _, line := range input {
		sb.WriteString(line)
	}
	return lineLength, sb.String()
}

func topLeft(c, lineLength int, input string) string {
	tl := c - lineLength - 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == 0 {
		return ""
	}

	return string(input[tl])
}

func top(c, lineLength int, input string) string {
	tl := c - lineLength
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl {
		return ""
	}

	return string(input[tl])
}

func topRight(c, lineLength int, input string) string {
	tl := c - lineLength + 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == lineLength-1 {
		return ""
	}

	return string(input[tl])
}

func right(c, lineLength int, input string) string {
	tl := c + 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == lineLength-1 {
		return ""
	}

	return string(input[tl])
}

func bottomRight(c, lineLength int, input string) string {
	tl := c + lineLength + 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == lineLength-1 {
		return ""
	}

	return string(input[tl])
}

func bottom(c, lineLength int, input string) string {
	tl := c + lineLength
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl {
		return ""
	}

	return string(input[tl])
}

func bottomLeft(c, lineLength int, input string) string {
	tl := c + lineLength - 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == 0 {
		return ""
	}

	return string(input[tl])
}

func left(c, lineLength int, input string) string {
	tl := c - 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == 0 {
		return ""
	}

	return string(input[tl])
}

// t1adjacent returns the number of occupied seats adjacent to the one specified by the coordinates. The coordinates h,v
// stand for horizontal and vertical. Horizontal is a character inside of a line of string, vertical is an entire line
// of string inside the slice of strings.
//
// Returns error if the coordinate is a floor, or if the coordinate is out of bounds.
func t1adjacent(h, v int, state []string) (int, error) {
	if v < 0 || v > len(state)-1 {
		return 0, outOfBounds{Message: fmt.Sprintf("vertical is out of bounds: %d", v)}
	}
	if h < 0 || h > len(state[0])-1 {
		return 0, outOfBounds{Message: fmt.Sprintf("horizontal is out of bounds: %d", h)}
	}

	if string(state[v][h]) == "." {
		return 0, floorIsLave{Message: fmt.Sprintf("spot at coordinate h: %d, v: %d is a floor", h, v)}
	}
	return 0, nil
}

func task2() {
	_ = getInputs()
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
