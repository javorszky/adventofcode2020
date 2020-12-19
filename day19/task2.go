package day19

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	task2RuleReplacer = strings.NewReplacer("8: 42\n", "8: 42 | 42 8\n", "11: 42 31\n", "11: 42 31 | 42 11 31\n")
)

func task2() {
	rules, messages := getInputs2()
	parsedRules := parseRules(rules)
	longestLine := 0
	for _, message := range messages {
		if len(message) > longestLine {
			longestLine = len(message)
		}
	}

	rule11RecursionIsBadReplacements := make([]string, 0)
	i := 0
	for {
		i++
		s11 := strings.Trim(strings.Repeat("42 ", i)+strings.Repeat("31 ", i), " ")
		if len(s11) > longestLine {
			break
		}

		rule11RecursionIsBadReplacements = append(rule11RecursionIsBadReplacements, s11)
	}

	rule8RecursionIsBadReplacements := make([]string, 0)
	i = 0
	for {
		i++
		s8 := strings.Trim(strings.Repeat("42 ", i), " ")
		if len(s8) > longestLine {
			break
		}

		rule8RecursionIsBadReplacements = append(rule8RecursionIsBadReplacements, s8)
	}

	parsedRules["11"] = strings.Join(rule11RecursionIsBadReplacements, " | ")
	parsedRules["8"] = strings.Join(rule8RecursionIsBadReplacements, " | ")

	ruleZero := "^" + despacer.Replace(findRule("0", parsedRules)) + "$"
	reMessages := regexp.MustCompile(ruleZero)

	n := 0
	for _, message := range messages {
		if reMessages.MatchString(message) {
			n++
		}
	}

	fmt.Printf("\nDay 19 task 2: there are %d messages matching the ruleset\n\nbtw the regex is too large to display on terminals with non unlimited scrollbacks, but it's in the repository in a file\n", n)
}

// getInputs reads the input.txt file and returns the rules and messages as slices of strings each.
func getInputs2() ([]string, []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := strings.TrimRight(string(data), "\n")
	parts := strings.Split(input, "\n\n")
	return strings.Split(task2RuleReplacer.Replace(parts[0]), "\n"), strings.Split(parts[1], "\n")
}
