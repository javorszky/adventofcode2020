package day7

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const filename = "day7/input.txt"
const noBags = "no other bags"
const shinyGold = "shiny gold"

var reDep = regexp.MustCompile(`^\d+ (.+)? bags?$`)
var reNumberedDep = regexp.MustCompile(`(\d)+ (.+)? bags?$`)

func Tasks() {
	task1()

	task2()
}

func task1() {
	// deps has the format of "colour is contained in these other colours", so
	// "wavy yellow": {
	//   "dull purple":{},
	//   "pale purple":{},
	// },
	// means that dull purple and pale purple both contain wavy yellow.
	//
	// there's a "colour" called "no other bags" which contain all the top level bags.
	deps := make(map[string]map[string]struct{}, 0)

	for _, dep := range getInputs() {
		// break it up at contain
		outer, inners := extractDep(dep)
		for _, i := range inners {
			if deps[i] == nil {
				deps[i] = make(map[string]struct{}, 0)
			}
			deps[i][outer] = struct{}{}
		}
	}

	// f returns the bags in a map that can hold the bag named s. If it's in the no other bags, it returns true in bool.
	// if it can't be found, returns an error.
	var f func(string) map[string]interface{}
	endBags := make(map[string]struct{}, 0)
	f = func(s string) map[string]interface{} {
		containers, ok := deps[s]
		endBags[s] = struct{}{}
		if !ok {
			return nil
		}
		intermediary := make(map[string]interface{}, 0)
		for d := range containers {
			intermediary[d] = f(d)
		}

		return intermediary
	}

	_ = f(shinyGold)

	fmt.Printf("\nDay 7 task 1: the number of different colours an outermost bag can be for a shiny gold is %d\n", len(endBags)-1)
}

func task2() {
	deps := make(map[string]map[string]int, 0)

	for _, dep := range getInputs() {
		// break it up at contain
		outer, innerMap := extractContains(dep)
		deps[outer] = innerMap
	}

	fmt.Printf("Day 7 task 2: 1 shiny gold bag contains %d total other bags\n", t2accumulator(deps, shinyGold))
}

func t2accumulator(m map[string]map[string]int, s string) int {
	var accumulator func(string) int
	accumulator = func(s string) int {
		i := 0
		for colour, num := range m[s] {
			i = i + num + num*accumulator(colour)
		}
		return i
	}

	return accumulator(s)
}

func extractDep(s string) (string, []string) {
	s = strings.TrimRight(s, ".")
	parts := strings.Split(s, " bags contain ")
	deps := make([]string, 0)

	dependents := strings.Split(parts[1], ", ")
	for _, dep := range dependents {
		found := reDep.FindAllStringSubmatch(dep, -1)
		if found == nil {
			if "no other bags" == dep {
				return parts[0], []string{noBags}
			}
			//if len(found) < 1 && len(found[0]) < 2 {
			fmt.Printf("something's gone wrong. Can't find correct matches for this string:\n%s\n", dep)
			os.Exit(1)
		}
		deps = append(deps, found[0][1])
	}
	return parts[0], deps
}

func extractContains(s string) (string, map[string]int) {
	s = strings.TrimRight(s, ".")
	parts := strings.Split(s, " bags contain ")
	deps := make(map[string]int, 0)

	dependents := strings.Split(parts[1], ", ")
	for _, dep := range dependents {
		found := reNumberedDep.FindAllStringSubmatch(dep, -1)
		if found == nil {
			if noBags == dep {
				return noBags, map[string]int{}
			}
			fmt.Printf("something's gone wrong. Can't find correct matches for this string:\n%s\n", dep)
			os.Exit(1)
		}

		i, err := strconv.Atoi(found[0][1])
		if err != nil {
			fmt.Printf("can't turn string into int: %s\n", found[0][1])
			os.Exit(1)
		}
		deps[found[0][2]] = i
	}
	return parts[0], deps
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// strip the trailing newline
	sData := strings.TrimRight(string(data), "\n")

	return strings.Split(sData, "\n")
}
