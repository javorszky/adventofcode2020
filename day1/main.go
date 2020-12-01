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
	for i := range m {
		j := target - i
		if _, ok := m[j]; ok == true {
			fmt.Printf("found the pair: %d and %d. Their sum is %d and product is %d\n", i, j, i+j, i*j)
			os.Exit(0)
		}
	}
	fmt.Println("found no pairs :(")
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
