package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"
const target = 2020

func main() {
	m := getInputs()

	// func to find 2 numbers is arr that sum to t.
	f := func(t int, arr map[int]struct{}) (int, int, error) {
		for i := range arr {
			j := t - i
			if _, ok := arr[j]; ok == true {
				return i, j, nil
			}
		}
		return 0, 0, fmt.Errorf("no pair found for %d", t)
	}

	// task 1: find 2 numbers that sum to 2020.
	i1, i2, err := f(2020, m)
	if err != nil {
		fmt.Println("no pair for 2020")
		os.Exit(1)
	}

	fmt.Printf("found the pair for %d: %d and %d. Their sum is %d and product is %d\n", target, i1, i2, i1+i2, i1*i2)

	// task 2: find 3 numbers that sum to 2020. For that take 1 number, subtract from 2020, then find 2 other numbers
	// that sum to the remainder.
	for k := range m {
		l := target - k
		o1, o2, err := f(l, m)
		if err == nil {
			fmt.Printf("found the triplet: %d, %d, and %d. Their sum is %d, and their product is %d\n", k, o1, o2, k+o1+o2, k*o1*o2)
			os.Exit(0)
		}

	}
	fmt.Println("found no triplets :(")
	os.Exit(1)
}

// getInputs reads the input.txt file, and arranges the contents into a map of unique numbers.
func getInputs() map[int]struct{} {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var intMap = make(map[int]struct{})

	b := strings.Split(string(data), "\n")
	for _, element := range b {
		if element == "" {
			continue
		}
		i, err := strconv.Atoi(element)
		if err != nil {
			fmt.Printf("whoops, couldn't convert %s to int: %s\n", element, err)
			continue
		}
		intMap[i] = struct{}{}
	}

	return intMap
}
