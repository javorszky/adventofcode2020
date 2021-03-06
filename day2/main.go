package day2

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const filename = "day2/input.txt"
const lineParseRegex = `(\d+)-(\d+) ([a-z]): (\w+)`

var re = regexp.MustCompile(lineParseRegex)
var reError = errors.New("did not find matches in the line for the pattern")

func Tasks() {
	goodPasswords := 0
	goodPasswordsT2 := 0
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

		if checkPasswordTask2(min, max, letter, pw) {
			goodPasswordsT2++
		}
	}

	fmt.Printf("\nDay 2 task 1: there were %d good passwords\n"+
		"Day 2 task 2: also %d passwords according to the new rule\n", goodPasswords, goodPasswordsT2)
}

func checkPassword(min, max int, letter, pw string) bool {
	c := strings.Count(pw, letter)
	return min <= c && max >= c
}

func checkPasswordTask2(pos1, pos2 int, letter, pw string) bool {
	return (letter == string(pw[pos1-1])) != (letter == string(pw[pos2-1]))
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
