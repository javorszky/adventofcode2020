package day11

import "strings"

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

func topLeft(c, lineLength int, input string) (int, string) {
	tl := c - lineLength - 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == 0 {
		return -1, ""
	}

	return tl, string(input[tl])
}

func top(c, lineLength int, input string) (int, string) {
	tl := c - lineLength
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl {
		return -1, ""
	}

	return tl, string(input[tl])
}

func topRight(c, lineLength int, input string) (int, string) {
	tl := c - lineLength + 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == lineLength-1 {
		return -1, ""
	}

	return tl, string(input[tl])
}

func right(c, lineLength int, input string) (int, string) {
	tl := c + 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == lineLength-1 {
		return -1, ""
	}

	return tl, string(input[tl])
}

func bottomRight(c, lineLength int, input string) (int, string) {
	tl := c + lineLength + 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == lineLength-1 {
		return -1, ""
	}

	return tl, string(input[tl])
}

func bottom(c, lineLength int, input string) (int, string) {
	tl := c + lineLength
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl {
		return -1, ""
	}

	return tl, string(input[tl])
}

func bottomLeft(c, lineLength int, input string) (int, string) {
	tl := c + lineLength - 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == 0 {
		return -1, ""
	}

	return tl, string(input[tl])
}

func left(c, lineLength int, input string) (int, string) {
	tl := c - 1
	// if we're out of bounds or in the left column.
	if 0 > tl || len(input)-1 < tl || c%lineLength == 0 {
		return -1, ""
	}

	return tl, string(input[tl])
}

// t1adjacent returns the number of occupied seats adjacent to the one specified by the coordinates. The coordinates h,v
// stand for horizontal and vertical. Horizontal is a character inside of a line of string, vertical is an entire line
// of string inside the slice of strings.
//
// Returns error if the coordinate is a floor, or if the coordinate is out of bounds.
func t1adjacentOccupied(c, lineLength int, input string) int {
	fs := []func(int, int, string) (int, string){
		top,
		topRight,
		right,
		bottomRight,
		bottom,
		bottomLeft,
		left,
		topLeft,
	}
	var sb strings.Builder
	for _, f := range fs {
		_, seat := f(c, lineLength, input)
		sb.WriteString(seat)
	}

	return strings.Count(sb.String(), "#")
}

func t2directionOccupied(c, lineLength int, input string) int {
	fs := []func(int, int, string) (int, string){
		top,
		topRight,
		right,
		bottomRight,
		bottom,
		bottomLeft,
		left,
		topLeft,
	}
	var sb strings.Builder
	for _, f := range fs {
		sb.WriteString(walk(c, lineLength, input, f))
	}

	return strings.Count(sb.String(), "#")
}

func walk(c, lineLength int, input string, f func(int, int, string) (int, string)) string {
	cn := c
	for {
		cnext, seat := f(cn, lineLength, input)
		switch seat {
		case "L", "#", "":
			return seat
		}
		cn = cnext
	}
}

func t1flip(lineLength int, input string) string {
	var sb strings.Builder

	for idx, c := range input {
		// floors never change
		switch string(c) {
		case ".":
			sb.WriteString(".")
		case "L":
			// field is empty, can we make it occupied?
			if t1adjacentOccupied(idx, lineLength, input) == 0 {
				sb.WriteString("#")
			} else {
				sb.WriteString("L")
			}
		case "#":
			// field is occupied, is it going to be empty?
			if t1adjacentOccupied(idx, lineLength, input) >= 4 {
				sb.WriteString("L")
			} else {
				sb.WriteString("#")
			}
		}
	}
	return sb.String()
}

func t2flip(lineLength int, input string) string {
	var sb strings.Builder

	for idx, c := range input {
		// floors never change
		switch string(c) {
		case ".":
			sb.WriteString(".")
		case "L":
			// field is empty, can we make it occupied?
			if t1adjacentOccupied(idx, lineLength, input) == 0 {
				sb.WriteString("#")
			} else {
				sb.WriteString("L")
			}
		case "#":
			// field is occupied, is it going to be empty?
			if t1adjacentOccupied(idx, lineLength, input) >= 5 {
				sb.WriteString("L")
			} else {
				sb.WriteString("#")
			}
		}
	}
	return sb.String()
}
