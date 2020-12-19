package day19

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// task1 will do the calculations necessary to compute day 19 task 1:
//
// "Your goal is to determine the number of messages that completely match rule 0."
func task1() {
	rules, messages := getInputs()
	parsedRules := parseRules(rules)

	ruleZero := "^" + despacer.Replace(findRule("0", parsedRules)) + "$"

	reMessages := regexp.MustCompile(ruleZero)

	n := 0
	for _, message := range messages {
		if reMessages.MatchString(message) {
			n++
		}
	}

	fmt.Printf("\nDay 19 task 1: there are %d messages matching the ruleset\n\nbtw the regex needed for it is:\n----\n%s\n----\n\n", n, ruleZero)
}

func findRule(s string, rules map[string]string) string {
	rule, ok := rules[s]
	if !ok {
		panic(fmt.Sprintf("could not find rule %s in rules map", s))
	}
	switch rule {
	case `"a"`, `"b"`:
		return strings.Trim(rule, `"`)
	default:
		ors := strings.Split(rule, " | ")
		orsStrings := make([]string, 0)
		for _, or := range ors {
			numbers := strings.Split(or, " ")
			numbersStrings := make([]string, 0)
			for _, number := range numbers {
				numbersStrings = append(numbersStrings, findRule(number, rules))
			}
			orsStrings = append(orsStrings, strings.Join(numbersStrings, " "))
		}
		return "( " + strings.Join(orsStrings, " | ") + " )"
	}
}

// getInputs reads the input.txt file and returns the rules and messages as slices of strings each.
func getInputs() ([]string, []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := strings.TrimRight(string(data), "\n")
	parts := strings.Split(input, "\n\n")
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}
