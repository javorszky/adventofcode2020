package day16

import (
	"fmt"
	"strconv"
	"strings"
)

func task1() {
	values := make(map[int]struct{}, 0)
	//rules, myTicket, nearbyTickets := getInputs()
	rules, _, nearbyTickets := getInputs()
	for _, rule := range rules {
		matches := reRule.FindStringSubmatch(rule)
		values = mergeRanges(values, matches[2], matches[3])
		values = mergeRanges(values, matches[4], matches[5])
	}

	serint, sermap := checkNearbyTickets(values, nearbyTickets)
	fmt.Printf("\nDay 16 task 1\n"+
		"scanning error rate when duplicate values are counted each time: %d (this was the correct one)\n"+
		"scanning error rate when duplicate values are counted only once: %d\n", serint, sermap)
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

func checkNearbyTickets(values map[int]struct{}, nearbyTickets []string) (int, int) {
	sumMap := make(map[int]struct{}, 0)
	sum := 0
	for _, ticket := range nearbyTickets {
		tValues := strings.Split(ticket, ",")
		for _, tValue := range tValues {
			tv, err := strconv.Atoi(tValue)
			if err != nil {
				panic(fmt.Sprintf("checknearbytickets: %s, err: %s", tValue, err))
			}
			if _, ok := values[tv]; !ok {
				sum = sum + tv
				sumMap[tv] = struct{}{}
			}
		}
	}

	summ := 0
	for idx := range sumMap {
		summ = summ + idx
	}

	return sum, summ
}
