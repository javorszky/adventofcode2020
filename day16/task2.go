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

	ruleStrings, _, nearbyTickets := getInputs()
	values := make(map[int]struct{}, 0)
	allRules := make(map[string]struct{}, 0)
	for _, rule := range ruleStrings {
		matches := reRule.FindStringSubmatch(rule)
		allRules[matches[1]] = struct{}{}
		values = mergeRanges(values, matches[2], matches[3])
		values = mergeRanges(values, matches[4], matches[5])
	}

	// i have a slice of rule structs.
	rules := createRules(ruleStrings)

	// i have a list of valid tickets.
	validTix := filterInvalidTickets(values, nearbyTickets)

	// I have a list of possibilities from all tickets for all columns.
	possibilities := make([]map[int][]string, 0)
	for _, validTicket := range validTix {
		possibilities = append(possibilities, parseTicketPossibilities(validTicket, rules))
	}

	sherlockPossibilities(possibilities, allRules)
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

func parseTicketPossibilities(ticket string, rules []rule) map[int][]string {
	possibilities := make(map[int][]string, 0)

	tValues := strings.Split(ticket, ",")
	for idx, tValue := range tValues {
		row := idx + 1
		tv, err := strconv.Atoi(tValue)
		if err != nil {
			panic(fmt.Sprintf("checknearbytickets: %s, err: %s", tValue, err))
		}
		for _, rule := range rules {
			if rule.Valid(tv) {
				possibilities[row] = append(possibilities[row], rule.name)
			}
		}
	}
	return possibilities
}

func sherlockPossibilities(in []map[int][]string, allRules map[string]struct{}) map[int]string {
	// used to hold rows that have been most definitely been figured out.
	//used := make(map[string]struct{}, 0)

	// used to store all the possible column names for a given column from all the tickets.
	intermediary := make(map[int][]map[string]struct{}, 0)

	// used to store a second step of intermediary where we discarded all possibilities that do not agree between the
	// individual columns. This is similar to the intermediary, except collapses the slice into a map of values that are
	// all present in the maps in the slice.
	uniqueIntermediary := make(map[int]map[string]struct{}, 0)

	// first let's loop through all the possibilities, and build the intermediary.
	for _, pos := range in {
		for col, titles := range pos {
			t := make(map[string]struct{}, 0)
			for _, title := range titles {
				t[title] = struct{}{}
			}
			intermediary[col] = append(intermediary[col], t)
		}
	}

	for column, possibilities := range intermediary {
		colPos := make(map[string]struct{}, 0)
		for k, v := range allRules {
			colPos[k] = v
		}

		for _, possibility := range possibilities {
			for col := range colPos {
				if _, ok := possibility[col]; !ok {
					delete(colPos, col)
				}
			}
		}

		uniqueIntermediary[column] = colPos
	}

	return collapsePossibilities(uniqueIntermediary)
}

func collapsePossibilities(unique map[int]map[string]struct{}) map[int]string {
	next := ""
	out := make(map[int]string, 0)
	goon := true
	for {
		if !goon {
			break
		}
		goon = false
		for column, titles := range unique {
			// remove the last found column from the possibilities first.
			delete(titles, next)
			unique[column] = titles
		}
		for column, titles := range unique {
			if len(titles) == 1 {
				for t := range titles {
					goon = true
					next = t
					out[column] = t
				}
				break
			}
		}
	}
	return out
}
