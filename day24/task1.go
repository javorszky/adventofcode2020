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
	_ = getInputs()
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

func collapseDirections(d directions) directions {
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
