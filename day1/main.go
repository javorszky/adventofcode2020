package day1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"
const target = 2020

func Tasks() {
	m := getInputs()

	p1, p2, err := findPair(m)
	if err != nil {
		fmt.Printf("could not find pair... %s", err)
		os.Exit(1)
	}
	fmt.Printf("found the pair for %d: %d and %d. Their sum is %d and product is %d\n", target, p1, p2, p1+p2, p1*p2)

	t1, t2, t3, err := findTriplet(m)
	if err != nil {
		fmt.Printf("could not find triplet: %s", err)
		os.Exit(1)
	}
	fmt.Printf("found the triplet: %d, %d, and %d. Their sum is %d, and their product is %d\n", t1, t2, t3, t1+t2+t3, t1*t2*t3)
	os.Exit(0)
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

// findPair will return the pair of numbers that sum to 2020 from the input map.
func findPair(m map[int]struct{}) (int, int, error) {
	return findPairFor(target, m)
}

// findTriple will return three numbers that sum to 2020 in input map.
func findTriplet(m map[int]struct{}) (int, int, int, error) {
	for k := range m {
		l := target - k
		o1, o2, err := findPairFor(l, m)
		if err == nil {
			return k, o1, o2, nil
		}
	}
	return 0, 0, 0, fmt.Errorf("no triplet found for %d", target)
}

// findPairFor will return two integers that sum to t.
func findPairFor(t int, arr map[int]struct{}) (int, int, error) {
	for i := range arr {
		j := t - i
		if _, ok := arr[j]; ok == true {
			return i, j, nil
		}
	}
	return 0, 0, fmt.Errorf("no pair found for %d", t)
}
