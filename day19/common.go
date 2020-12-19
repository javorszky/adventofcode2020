package day19

import (
	"regexp"
	"strings"
)

const filename = "day19/input.txt"

var (
	reParseRuleString = regexp.MustCompile(`^(\d+): (.*)$`)
	despacer          = strings.NewReplacer(" ", "")
)

func parseRules(ruleStrings []string) map[string]string {
	rules := make(map[string]string, 0)
	for _, ruleString := range ruleStrings {
		matches := reParseRuleString.FindStringSubmatch(ruleString)
		rules[matches[1]] = matches[2]
	}

	return rules
}
