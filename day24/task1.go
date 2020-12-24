package day24

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	se          = "se"
	sw          = "sw"
	ne          = "ne"
	nw          = "nw"
	e           = "e"
	w           = "w"
	replacement = "XX"
)

var (
	reSE = regexp.MustCompile("se")
	reSW = regexp.MustCompile("sw")
	reNE = regexp.MustCompile("ne")
	reNW = regexp.MustCompile("nw")
)

type directions map[string]int

func task1() {
	fmt.Printf("\nDay 24 task 1: There will be %d black tiles at the end\n", countBlackTiles(getInputs()))
}

func countBlackTiles(tiles []string) int {
	world := make(map[string]int, 0)
	for _, tile := range tiles {
		world[getCoordinate(collapseDirections(parseDirections(tile)))]++
	}

	black := 0
	for _, flipped := range world {
		if flipped%2 == 1 {
			black++
		}
	}
	return black
}

func parseDirections(s string) directions {
	d := make(directions, 0)

	// extract se
	m := reSE.FindAllString(s, -1)
	d[se] = len(m)
	s = reSE.ReplaceAllString(s, replacement)

	// extract sw
	m = reSW.FindAllString(s, -1)
	d[sw] = len(m)
	s = reSW.ReplaceAllString(s, replacement)

	// extract ne
	m = reNE.FindAllString(s, -1)
	d[ne] = len(m)
	s = reNE.ReplaceAllString(s, replacement)

	// extract nw
	m = reNW.FindAllString(s, -1)
	d[nw] = len(m)
	s = reNW.ReplaceAllString(s, replacement)

	// extract e and w
	d[e] = strings.Count(s, e)
	d[w] = strings.Count(s, w)
	return d
}

func collapseDirections(in directions) directions {
	d := make(directions, 0)
	// copy the thing
	for k, v := range in {
		d[k] = v
	}

	coordinate := getCoordinate(d)
	for {
		collapseHelper := [][]string{
			{se, w, sw},
			{sw, nw, w},
			{w, ne, nw},
			{nw, e, ne},
			{ne, se, e},
			{e, sw, se},
		}
		for _, triplets := range collapseHelper {
			absd := smoler(d[triplets[0]], d[triplets[1]])
			d[triplets[0]] = d[triplets[0]] - absd
			d[triplets[1]] = d[triplets[1]] - absd
			d[triplets[2]] = d[triplets[2]] + absd
		}

		// se / nw
		if d[se] > d[nw] {
			d[se] = d[se] - d[nw]
			d[nw] = 0
		} else {
			d[nw] = d[nw] - d[se]
			d[se] = 0
		}

		// sw / ne
		if d[sw] > d[ne] {
			d[sw] = d[sw] - d[ne]
			d[ne] = 0
		} else {
			d[ne] = d[ne] - d[sw]
			d[sw] = 0
		}

		// e / w
		if d[e] > d[w] {
			d[e] = d[e] - d[w]
			d[w] = 0
		} else {
			d[w] = d[w] - d[e]
			d[e] = 0
		}

		if coordinate == getCoordinate(d) {
			break
		}
		coordinate = getCoordinate(d)
	}

	return d
}

func getCoordinate(d directions) string {
	return fmt.Sprintf("%d.%d.%d.%d.%d.%d",
		d[ne],
		d[e],
		d[se],
		d[sw],
		d[w],
		d[nw],
	)
}

// absDiff will return the absolute difference between two numbers.
func absDiff(a, b int) int {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}

// smoler takes two ints and returns the smaller of the two.
func smoler(a, b int) int {
	if a < b {
		return a
	}
	return b
}
