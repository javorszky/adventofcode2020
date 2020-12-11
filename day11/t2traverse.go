package day11

import "strings"

func t2flip(lineLength int, input string) string {
	var sb strings.Builder

	for idx, c := range input {
		// floors never change
		switch string(c) {
		case ".":
			sb.WriteString(".")
		case "L":
			// field is empty, can we make it occupied?
			if t2directionOccupied(idx, lineLength, input) == 0 {
				sb.WriteString("#")
			} else {
				sb.WriteString("L")
			}
		case "#":
			// field is occupied, is it going to be empty?
			if t2directionOccupied(idx, lineLength, input) >= 5 {
				sb.WriteString("L")
			} else {
				sb.WriteString("#")
			}
		}
	}
	return sb.String()
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
