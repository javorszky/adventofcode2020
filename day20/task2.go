package day20

import "strings"

func task2() {
	tss := make(tileSets2, 0)
	// there are 144 tiles, so that's a 12x12 grid.
	for _, tileString := range getInputs() {
		tss.addTileSet(newTileSet2(parseTileTask2(tileString)))
	}
	i := image2{
		W:     12,
		H:     12,
		Tiles: make(map[int]map[int]tilev2, 0),
	}

	_, _ = fitTileIntoImage2(i, 1, tss)
}

func parseTileTask2(s string) tilev2 {
	rows := strings.Split(s, "\n")
	m := reTile.FindStringSubmatch(rows[0])

	t := tilev2{
		ID: m[1] + "000",
	}

	var (
		s2     strings.Builder
		s4     strings.Builder
		inside = make([]string, 0)
	)
	for idx, row := range rows[1:] {
		s4.WriteString(row[0:1])
		s2.WriteString(row[9:10])
		switch idx {
		case 0:
			t.Top = row
		case 9:
			t.Bottom = row
		default:
			inside = append(inside, row[1:9])
		}
	}
	t.Right = s2.String()
	t.Left = s4.String()
	t.Content = inside
	return t
}
