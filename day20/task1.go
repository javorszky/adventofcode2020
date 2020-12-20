package day20

import (
	"fmt"
	"regexp"
	"strings"
)

var reTile = regexp.MustCompile(`^Tile (\d+):$`)

func reverseString(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]))
	}
	return sb.String()
}

func task1() {
	tiles := make([]tile, 0)
	// there are 144 tiles, so that's a 12x12 grid.
	for _, tileString := range getInputs() {
		tiles = append(tiles, parseTile(tileString))
	}

	fmt.Printf("\nDay 20 task 1: there are %d tiles\n", len(tiles))
}

func parseTile(s string) tile {
	rows := strings.Split(s, "\n")
	m := reTile.FindStringSubmatch(rows[0])

	t := tile{
		ID: m[1] + "000",
	}

	var (
		s2 strings.Builder
		s4 strings.Builder
	)
	for idx, row := range rows[1:] {
		s4.WriteString(row[0:1])
		s2.WriteString(row[9:10])
		switch idx {
		case 0:
			t.Top = row
		case 9:
			t.Bottom = row
		}
	}
	t.Right = s2.String()
	t.Left = s4.String()
	return t
}
