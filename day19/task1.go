package day19

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	reParseRuleString = regexp.MustCompile(`^(\d+): (.*)$`)
)

// task1 will do the calculations necessary to compute day 19 task 1:
//
// "Your goal is to determine the number of messages that completely match rule 0."
func task1() {
	rules, messages := getInputs()
	fmt.Printf("rules:\n%#v\n\nmessages:\n%#v\n", rules, messages)
}

func parseRules(ruleStrings []string) map[string]string {
	rules := make(map[string]string, 0)
	for _, ruleString := range ruleStrings {
		matches := reParseRuleString.FindStringSubmatch(ruleString)
		rules[matches[1]] = matches[2]
	}

	return rules
}

func findRule(s string, rules map[string]string) string {
	rule, ok := rules[s]
	if !ok {
		panic(fmt.Sprintf("could not find rule %s in rules map", s))
	}
	switch rule {
	case "a", "b":
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
