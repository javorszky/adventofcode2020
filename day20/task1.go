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
	tss := make(tileSets, 0)
	// there are 144 tiles, so that's a 12x12 grid.
	for _, tileString := range getInputs() {
		tss.addTileSet(newTileSet(parseTile(tileString)))
	}
	i := image{
		W:     3,
		H:     3,
		Tiles: make(map[int]map[int]tile, 0),
	}

	img, err := fitTileIntoImage(i, 1, tss)

	if err != nil {
		panic("day 20 task 1 failed")
	}
	fmt.Printf("dy 20 t 1 we have an image\n\n%#v", img)
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
