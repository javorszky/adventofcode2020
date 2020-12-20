package day20

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reTile = regexp.MustCompile(`^Tile (\d+):$`)

// tile represents a tile with its ID and four sides. These are:
//
// Side1 - top from left to right
// Side2 - right from top to bottom
// Side3 - bottom from left to right
// Side4 - left from top to bottom
type tile struct {
	ID                         int
	Side1, Side2, Side3, Side4 string
}

// flipV returns a tile that's been vertically flipped, which means Side2 and Side4 were reversed, and Side1 and Side3
// have switched places.
func (t tile) flipV() tile {
	t.Side2 = reverseString(t.Side2)
	t.Side4 = reverseString(t.Side4)
	tempS1 := t.Side1
	t.Side1 = t.Side3
	t.Side3 = tempS1
	return t
}

// flipH returns a tile that's been horizontally flipped, which means Side1 and Side3 have been reversed, and Side2 and
// Side4 have switched places.
func (t tile) flipH() tile {
	t.Side1 = reverseString(t.Side1)
	t.Side3 = reverseString(t.Side3)
	tempS2 := t.Side2
	t.Side2 = t.Side4
	t.Side4 = tempS2
	return t
}

// rotate will return a tile that's been rotated 90 degrees clockwise, so:
//
// Side1 becomes Side4
// Side4 becomes Side3
// Side3 becomes Side2
// Side2 becomes Side1
func (t tile) rotate() tile {
	tempS1 := t.Side1 // save top
	t.Side1 = t.Side4 // top becomes left
	t.Side4 = t.Side3 // left becomes bottom
	t.Side3 = t.Side2 // bottom becomes right
	t.Side2 = tempS1
	return t
}

func reverseString(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]))
	}
	return sb.String()
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
