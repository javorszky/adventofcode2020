package day20

import (
	"fmt"
	"regexp"
	"strconv"
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
	for _, tileString := range getInputs(filename) {
		tss.addTileSet(newTileSet(parseTile(tileString)))
	}
	i := image{
		W:     12,
		H:     12,
		Tiles: make(map[int]map[int]tile, 0),
	}

	img, err := fitTileIntoImage(i, 1, tss)

	if err != nil {
		panic("\nDay 20 task 1 failed")
	}

	product := extractIDsAndMultiplyThem(img.Tiles[0][0], img.Tiles[0][11], img.Tiles[11][0], img.Tiles[11][11])

	fmt.Printf("\nDay 20 task 1: the product of the ids of the tiles in the four corners is %d\n", product)
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

func extractIDsAndMultiplyThem(tiles ...tile) int {
	product := 1
	for _, tile := range tiles {
		tid := tile.ID[:4]
		tidint, err := strconv.Atoi(tid)
		fmt.Printf("multiplying %d by new tid %d\n", product, tidint)
		if err != nil {
			panic(fmt.Sprintf("extract ids and multiplythem: could not turn '%s' into int: %s", tid, err))
		}
		product = product * tidint
		fmt.Printf("-- this gave us %d\n", product)
	}

	return product
}
