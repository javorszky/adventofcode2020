package day23

import (
	"fmt"
	"strconv"
	"strings"
)

type cupGame struct {
	Elements     []int
	Current, max int
}

func newCupGame(in []int) cupGame {
	slice := createCircleSlice(in)
	return cupGame{
		Elements: slice,
		Current:  in[0],
		max:      len(in),
	}
}

func (c *cupGame) play() {
	// get the three values that are clockwise to the Current one
	v1 := c.Elements[c.Current]
	v2 := c.Elements[v1]
	v3 := c.Elements[v2]
	afterCut := c.Elements[v3] // current will point here (before changing the current one)

	next := c.Current - 1
	for {
		if 0 == next {
			next = c.max
		}
		if next != v1 && next != v2 && next != v3 {
			break
		}
		next--
	}

	afterNext := c.Elements[next] // v3 will point here

	//fmt.Printf("elements: %#v\ncurrent: %d\nv1, v2, v3: %d, %d, %d\nnext: %d\naftercut: %d\nafternext: %d\n",
	//	c.Elements, c.Current, v1, v2, v3, next, afterCut, afterNext)

	// next points to v1 yep
	// v1 points to v2 (has not changed)
	// v2 points to v3 (has not changed)
	// v3 points to afterNext also yep
	// Current points to aftercut also yep
	// Current becomes the one to the right., which is... aftercut?
	c.Elements[next] = v1
	c.Elements[v3] = afterNext
	c.Elements[c.Current] = afterCut
	c.Current = afterCut
}

func printElements(c *cupGame) {
	h := make([]int, 0)
	next := c.Current
	for {
		h = append(h, next)
		next = c.Elements[next]
		if next == c.Current {
			break
		}
	}

	fmt.Printf("%#v\n", h)
}

func task1slice() {
	getLabel(getInputs(), 100)
}

func getLabel(in []int, step int) string {
	cg := newCupGame(in)
	for i := 0; i < step; i++ {
		cg.play()
	}

	var sb strings.Builder
	// start with 1
	next := cg.Elements[1]
	for {
		sb.WriteString(strconv.Itoa(next))
		next = cg.Elements[next]
		if next == 1 {
			break
		}
	}

	return sb.String()
}

func createCircleSlice(in []int) []int {
	crko := make([]int, len(in)+1)
	prev := 0
	first := 0
	for _, i := range in {
		if prev == 0 {
			first = i
			prev = i
			continue
		}
		crko[prev] = i
		prev = i
	}
	crko[prev] = first

	return crko
}
