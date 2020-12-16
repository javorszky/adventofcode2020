package day16

import (
	"fmt"
	"strconv"
)

func task1() {
	values := make(map[int]struct{}, 0)
	//rules, myTicket, nearbyTickets := getInputs()
	rules, _, _ := getInputs()
	for _, rule := range rules {
		matches := reRule.FindStringSubmatch(rule)
		fmt.Printf("%#v\n\n", matches)

		values = mergeRanges(values, matches[0], matches[1])
		values = mergeRanges(values, matches[2], matches[3])
	}
}

func mergeRanges(values map[int]struct{}, min, max string) map[int]struct{} {
	m, err := strconv.Atoi(min)
	if err != nil {
		panic(fmt.Sprintf("mergeranges min: %s, err: %s", min, err))
	}

	mx, err := strconv.Atoi(max)
	if err != nil {
		panic(fmt.Sprintf("mergeranges max: %s, err: %s", max, err))
	}

	normalizedMin, normalizedMax := m, mx

	if m > mx {
		normalizedMin = mx
		normalizedMax = m
	}

	for i := normalizedMin; i <= normalizedMax; i++ {
		values[i] = struct{}{}
	}
	return values
}
