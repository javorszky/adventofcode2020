package day20

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reTile = regexp.MustCompile(`^Tile (\d+):$`)

type tile struct {
	ID                         int
	Side1, Side2, Side3, Side4 string
}

func task1() {
	_ = getInputs()
}

func parseTile(s string) tile {
	rows := strings.Split(s, "\n")
	m := reTile.FindStringSubmatch(rows[0])
	mint, err := strconv.Atoi(m[1])
	if err != nil {
		panic(fmt.Sprintf("can't turn '%s' into an int: %s", m, err))
	}
	t := tile{
		ID: mint,
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
			t.Side1 = row
		case 9:
			t.Side3 = row
		}
	}
	t.Side2 = s2.String()
	t.Side4 = s4.String()
	return t
}
