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
	//position := getWorldBlackWhite(getWorld(getInputs()))

}

func blackTilesAfterNFlips(in map[string]string, n int) int {
	for i := 0; i < n; i++ {
		edged := getExpandedWorldCoordinates(in)
		in = flip(in, edged)
	}

	blackTiles := 0
	for _, colour := range in {
		if colour == black {
			blackTiles++
		}
	}
	return blackTiles
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

func getExpandedWorldCoordinates(in map[string]string) map[string]struct{} {
	out := make(map[string]struct{}, 0)
	for coordinate := range in {
		out[coordinate] = struct{}{}
		for _, neighbour := range getAdjacentAddresses(coordinate) {
			out[neighbour] = struct{}{}
		}
	}
	return out
}

func flip(in map[string]string, world map[string]struct{}) map[string]string {
	var getBlackTiles func([]string) int
	getBlackTiles = func(addresses []string) int {
		blacks := 0
		for _, address := range addresses {
			colour, ok := in[address]
			if !ok {
				continue
			}
			if colour == black {
				blacks++
			}
		}
		return blacks
	}
	out := make(map[string]string, 0)
	for address := range world {
		colour, ok := in[address]
		if !ok {
			colour = white
		}
		blackTiles := getBlackTiles(getAdjacentAddresses(address))
		switch blackTiles {
		case 2:
			// if we have a white tile with 2 black tiles adjacent, it turns black. If we have a black tile with two
			// adjacent black tiles, black remains.
			out[address] = black
		case 1:
			// Any tile with 1 black tile next to it retains its colour. If the tile did not exist (we're in expanded
			// universe territory), then the colour has already been set to white by this point, and it's safe to assign
			// it to the outgoing map.
			out[address] = colour
		default:
			// Cases 0, 3, 4, 5, 6
			// Any black tile that has no black tiles adjacent will flip to white. White tiles only change if there are
			// exactly two black tiles next to them, which has been handled.
			//
			// Any black tile with more than 2 black tiles adjacent will flip to white. White tiles only flip if there
			// are exactly 2 black tiles next to them, which has already been handled.
			out[address] = white
		}
	}
	return out
}
