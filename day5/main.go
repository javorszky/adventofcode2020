package day5

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "day5/input.txt"

func Tasks() {
	highSeat := 0
	for _, pass := range getInputs() {
		if pass == "" {
			continue
		}
		seat := findSeat(pass)
		if seat > highSeat {
			highSeat = seat
		}
	}

	fmt.Printf("highest seat number is %d\n", highSeat)
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func findSeat(code string) int {
	r := makeSlice(128)
	for _, fb := range strings.Split(code[:7], "") {
		r = binarySearch(r, fb)
	}
	row := r[0]
	c := makeSlice(8)
	for _, lr := range strings.Split(code[7:], "") {
		c = binarySearch(c, lr)
	}

	col := c[0]
	return row*8 + col
}

func binarySearch(in []int, end string) []int {
	if end == "B" || end == "R" {
		return in[len(in)/2:]
	}
	return in[:len(in)/2]
}

func makeSlice(n int) []int {
	i := make([]int, 0, n)
	for j := 0; j < n; j++ {
		i = append(i, j)
	}
	return i
}
