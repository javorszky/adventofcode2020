package day24

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	black = "black"
	white = "white"
)

var wheel = []string{se, sw, ne, nw, e, w}

func task2() {
	_ = getInputs()
}

func getWorldBlackWhite(in map[string]int) map[string]string {
	out := make(map[string]string, 0)
	for k, v := range in {
		if v%2 == 0 {
			out[k] = white
			continue
		}

		out[k] = black
	}
	return out
}

func getAdjacentAddresses(s string) []string {
	dirs := make([]string, 0)
	d := getDirectionsFromCoordinate(s)

	for _, dir := range wheel {
		// make a copy
		e := make(directions, 0)
		for k, v := range d {
			e[k] = v
		}
		e[dir]++
		dirs = append(dirs, getCoordinate(collapseDirections(e)))
	}

	return dirs
}

func getDirectionsFromCoordinate(coordinate string) directions {
	dirs := strings.Split(coordinate, ".")
	if len(dirs) != 6 {
		panic(fmt.Sprintf("receives coordinate does not have 6 elements; '%s'", coordinate))
	}
	ints := make([]int, 0)
	for _, v := range dirs {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("could not turn '%s' into an int", v))
		}
		ints = append(ints, n)
	}

	return directions{
		ne: ints[0],
		e:  ints[1],
		se: ints[2],
		sw: ints[3],
		w:  ints[4],
		nw: ints[5],
	}
}
