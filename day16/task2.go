package day16

import (
	"fmt"
	"strconv"
	"strings"
)

type rule struct {
	name     string
	min, max int
}

// Valid returns whether a given value is valid for the given rule.
func (r rule) Valid(value int) bool {
	return r.min <= value && value <= r.max
}

// newRule returns a rule after normalization of values.
func newRule(name, n1, n2 string) rule {
	n1int, err := strconv.Atoi(n1)
	if err != nil {
		panic(fmt.Sprintf("newRule n1: %s, err: %s", n1, err))
	}

	n2int, err := strconv.Atoi(n2)
	if err != nil {
		panic(fmt.Sprintf("newRule n2: %s, err: %s", n2, err))
	}

	normalizedMin, normalizedMax := n1int, n2int

	if n1int > n2int {
		normalizedMin = n2int
		normalizedMax = n1int
	}

	return rule{
		name: name,
		min:  normalizedMin,
		max:  normalizedMax,
	}
}

func task2() {
	_, _, _ = getInputs()

	//rules, _, nearbyTickets := getInputs()

}

func filterInvalidTickets(values map[int]struct{}, nearbyTickets []string) []string {
	validTix := make([]string, 0)
	for _, ticket := range nearbyTickets {
		valid := false
		tValues := strings.Split(ticket, ",")
		for _, tValue := range tValues {
			tv, err := strconv.Atoi(tValue)
			if err != nil {
				panic(fmt.Sprintf("checknearbytickets: %s, err: %s", tValue, err))
			}
			if _, ok := values[tv]; ok {
				valid = true
				break
			}
		}
		if valid {
			validTix = append(validTix, ticket)
		}
	}

	return validTix
}

// createRules takes a slice of rule strings, and returns a slice of individual rules.
func createRules(rules []string) []rule {
	rs := make([]rule, 0)
	for _, r := range rules {
		matches := reRule.FindStringSubmatch(r)
		rs = append(rs, newRule(matches[1], matches[2], matches[3]))
		rs = append(rs, newRule(matches[1], matches[4], matches[5]))
	}
	return rs
}
