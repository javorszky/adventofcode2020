package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "input.txt"
const startRow = 0
const stepRow = 1
const startCol = 0
const stepCol = 3

func main() {
	slope := getInputs()
	task1(slope)

	task2(slope)
}

func task1(slope []string) {
	trees := 0
	for i := 0; i < len(slope); i++ {
		r, c := step(i)
		if isTree(r, c, slope) {
			trees++
		}
	}
	fmt.Printf("We encountered %d trees.\n", trees)
}

func task2(slope []string) {

	steps := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	pTrees := 1
	for _, step := range steps {
		localTree := 0
		for i := 0; i < len(slope); i++ {
			r, c := varStep(i, step[0], step[1])
			if isTree(r, c, slope) {
				localTree++
			}
		}

		pTrees = pTrees * localTree
	}

	fmt.Printf("after having checked all the slopes, the product is %d\n", pTrees)
}

// getInputs reads the input.txt file, and arranges the contents into a map of unique numbers.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func step(n int) (int, int) {
	return startRow + n*stepRow, startCol + n*stepCol
}

func varStep(n, varStepRow, varStepCol int) (int, int) {
	return startRow + n*varStepRow, startCol + n*varStepCol
}

// isTree returns whether the ground at the given coordinates is a tree or not. Column is arbitrary number, is not
// constrained by width of the map.
func isTree(row, column int, slope []string) bool {
	if len(slope) <= row {
		return false
	}
	trueColumn := column % len(slope[0])

	r := slope[row]

	if len(r) == 0 {
		return false
	}

	return "#" == string(r[trueColumn])
}
