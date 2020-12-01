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
	f := func(t int, arr map[int]struct{}) (int, int, error) {
		for i := range arr {
			j := t - i
			if _, ok := arr[j]; ok == true {
				return i, j, nil
			}
		}
		return 0, 0, fmt.Errorf("no pair found for %d", t)
	}

	i1, i2, err := f(2020, m)
	if err != nil {
		fmt.Println("no pair for 2020")
		os.Exit(1)
	}

	fmt.Printf("found the pair for %d: %d and %d. Their sum is %d and product is %d\n", target, i1, i2, i1+i2, i1*i2)

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
