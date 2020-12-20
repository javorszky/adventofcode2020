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
// Top - top from left to right
// Right - right from top to bottom
// Bottom - bottom from left to right
// Left - left from top to bottom
type tile struct {
	ID                       int
	Top, Right, Bottom, Left string
}

// flipV returns a tile that's been vertically flipped, which means Right and Left were reversed, and Top and Bottom
// have switched places.
func (t tile) flipV() tile {
	t.Right = reverseString(t.Right)
	t.Left = reverseString(t.Left)
	tempS1 := t.Top
	t.Top = t.Bottom
	t.Bottom = tempS1
	return t
}

// flipH returns a tile that's been horizontally flipped, which means Top and Bottom have been reversed, and Right and
// Left have switched places.
func (t tile) flipH() tile {
	t.Top = reverseString(t.Top)
	t.Bottom = reverseString(t.Bottom)
	tempS2 := t.Right
	t.Right = t.Left
	t.Left = tempS2
	return t
}

// rotate will return a tile that's been rotated 90 degrees clockwise, so:
//
// Top becomes Left
// Left becomes Bottom
// Bottom becomes Right
// Right becomes Top
func (t tile) rotate() tile {
	tempS1 := t.Top    // save top
	t.Top = t.Left     // top becomes left
	t.Left = t.Bottom  // left becomes bottom
	t.Bottom = t.Right // bottom becomes right
	t.Right = tempS1
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
			t.Top = row
		case 9:
			t.Bottom = row
		}
	}
	t.Right = s2.String()
	t.Left = s4.String()
	return t
}
