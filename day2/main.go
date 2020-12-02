package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const filename = "input.txt"
const lineParseRegex = `(\d+)-(\d+) ([a-z]): (\w+)`

var re = regexp.MustCompile(lineParseRegex)
var reError = errors.New("did not find matches in the line for the pattern")

func main() {
	goodPasswords := 0
	for _, l := range getInputs() {
		if l == "" {
			continue
		}
		min, max, letter, pw, err := parseLine(l)
		if err != nil {
			fmt.Printf("parsing line: %s: %s\n", l, err)
			continue
		}
		if checkPassword(min, max, letter, pw) {
			goodPasswords++
		}
	}

	fmt.Printf("there were %d good passwords", goodPasswords)
}

func checkPassword(min, max int, letter, pw string) bool {
	c := strings.Count(pw, letter)
	fmt.Printf("checking %s for %s between %d and %d. count is %d\n", pw, letter, min, max, c)

	return min <= c && max >= c
}

// parseLine takes a line of policy and password, and returns the min, max, letter, and password from the line. If there
// was an error, it returns a non-nil error.
func parseLine(line string) (int, int, string, string, error) {
	stuff := re.FindAllStringSubmatch(line, -1)

	if len(stuff[0]) != 5 {
		return 0, 0, "", "", reError
	}

	min, err := strconv.Atoi(stuff[0][1])
	if err != nil {
		return 0, 0, "", "", fmt.Errorf("convert min string to int: %w", err)
	}
	max, err := strconv.Atoi(stuff[0][2])
	if err != nil {
		return 0, 0, "", "", fmt.Errorf("convert max string to int: %w", err)
	}

	return min, max, stuff[0][3], stuff[0][4], nil
}

// getInputs reads the input.txt file, and arranges the contents into a map of unique numbers.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}
