package day20

import (
	"fmt"
	"strings"
)

// tile represents a tile with its ID and four sides. These are:
//
// Top - top from left to right
// Right - right from top to bottom
// Bottom - bottom from left to right
// Left - left from top to bottom
type tilev2 struct {
	ID, Top, Right, Bottom, Left string
	Content, WithBorders         []string
}

// flipV returns a tilev2 that's been vertically flipped, which means Right and Left were reversed, and Top and Bottom
// have switched places.
func (t tilev2) flipV() tilev2 {
	t.Right = reverseString(t.Right)
	t.Left = reverseString(t.Left)
	tempS1 := t.Top
	t.Top = t.Bottom
	t.Bottom = tempS1
	t.Content = vFlipContent(t.Content)
	t.WithBorders = vFlipContent(t.WithBorders)

	tempID := t.ID[:4]
	switch t.ID[4:5] {
	case "0":
		tempID = tempID + "1" + t.ID[5:]
	case "1":
		tempID = tempID + "0" + t.ID[5:]
	default:
		panic(fmt.Sprintf("unexpected character in ID '%s' at position 4: '%s'", t.ID, string(t.ID[4])))
	}

	t.ID = tempID

	return t
}

// flipH returns a tilev2 that's been horizontally flipped, which means Top and Bottom have been reversed, and Right and
// Left have switched places.
func (t tilev2) flipH() tilev2 {
	t.Top = reverseString(t.Top)
	t.Bottom = reverseString(t.Bottom)
	tempS2 := t.Right
	t.Right = t.Left
	t.Left = tempS2
	t.Content = hFlipContent(t.Content)
	t.WithBorders = hFlipContent(t.WithBorders)

	tempID := t.ID[:5]
	switch t.ID[5:6] {
	case "0":
		tempID = tempID + "1" + t.ID[6:]
	case "1":
		tempID = tempID + "0" + t.ID[6:]
	default:
		panic(fmt.Sprintf("unexpected character in ID '%s' at position 5: '%s'", t.ID, string(t.ID[5])))
	}

	t.ID = tempID

	return t
}

// rotate will return a tilev2 that's been rotated 90 degrees clockwise, so:
//
// Top becomes reverse Left
// Left becomes Bottom
// Bottom becomes reverse Right
// Right becomes Top
func (t tilev2) rotate() tilev2 {
	tempS1 := t.Top                   // save top
	t.Top = reverseString(t.Left)     // top becomes left
	t.Left = t.Bottom                 // left becomes bottom
	t.Bottom = reverseString(t.Right) // bottom becomes right
	t.Right = tempS1
	t.Content = rotateContent(t.Content)
	t.WithBorders = rotateContent(t.WithBorders)

	tempID := ""
	switch t.ID[6:7] {
	case "0": // it was not rotated yet, so rotate it once, leave fliph and flipv as is.
		tempID = t.ID[:6] + "1"
	case "1": // it was rotated once. Rotating it again is the same as flipping both vertically and horizontally the 0
		// rotated version.
		tempID = t.ID[:4]

		// flipV
		switch t.ID[4:5] {
		case "1":
			tempID = tempID + "0"
		case "0":
			tempID = tempID + "1"
		default:
			panic(fmt.Sprintf("unexpected character in ID '%s' at position 5: '%s'", t.ID, t.ID[4:5]))
		}

		// flipH
		switch t.ID[5:6] {
		case "1":
			tempID = tempID + "0"
		case "0":
			tempID = tempID + "1"
		default:
			panic(fmt.Sprintf("unexpected character in ID '%s' at position 5: '%s'", t.ID, t.ID[5:6]))
		}
		// and reset rotation to 0.
		tempID = tempID + "0"
	default:
		panic(fmt.Sprintf("unexpected character in ID '%s' at position 7: '%s'", t.ID, t.ID[6:7]))
	}

	t.ID = tempID

	return t
}

func vFlipContent(c []string) []string {
	flippedC := make([]string, len(c))
	for idx, contentLine := range c {
		flippedC[len(c)-idx-1] = contentLine
	}
	return flippedC
}

func hFlipContent(c []string) []string {
	flippedC := make([]string, len(c))
	for idx, contentLine := range c {
		flippedC[idx] = reverseString(contentLine)
	}
	return flippedC
}

func rotateContent(c []string) []string {
	rotatedC := make([]string, len(c[0]))
	rotationHelper := make(map[int]strings.Builder, 0)

	for idx := range c[0] {
		var builder strings.Builder
		rotationHelper[idx] = builder
	}

	for _, line := range c {
		for idx, char := range line {
			builder := rotationHelper[idx]
			builder.WriteString(string(char))
			rotationHelper[idx] = builder
		}
	}

	for col, sb := range rotationHelper {
		rotatedC[col] = reverseString(sb.String())
	}

	return rotatedC
}

func parseTileTask2(s string) tilev2 {
	rows := strings.Split(s, "\n")
	m := reTile.FindStringSubmatch(rows[0])

	t := tilev2{
		ID: m[1] + "000",
	}

	var (
		s2          strings.Builder
		s4          strings.Builder
		inside      = make([]string, 0)
		withBorders = make([]string, 0)
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

		withBorders = append(withBorders, row)
	}
	t.Right = s2.String()
	t.Left = s4.String()
	t.Content = inside
	t.WithBorders = withBorders
	return t
}
