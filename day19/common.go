package day19

import (
	"io/ioutil"
	"regexp"
	"strings"
)

const filename = "day19/input.txt"

var (
	reParseRuleString = regexp.MustCompile(`^(\d+): (.*)$`)
	despacer          = strings.NewReplacer(" ", "")
)

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

func parseRules(ruleStrings []string) map[string]string {
	rules := make(map[string]string, 0)
	for _, ruleString := range ruleStrings {
		matches := reParseRuleString.FindStringSubmatch(ruleString)
		rules[matches[1]] = matches[2]
	}

	return rules
}
