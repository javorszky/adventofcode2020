package day13

import (
	"fmt"
	"strconv"
	"strings"
)

// the numbers are all primes.
func task2() {
	start := earliestTime(t2formatInput(getInputs()))
	fmt.Printf("Day 13 task 2: the timestamp is %d\n", start)
}

// togetherBus takes two buses, and returns a new one that envelops both of them.
func togetherBus(bus1Start, bus1ID, bus1offset, bus2Start, bus2ID, bus2offset int) (int, int, int) {
	if bus1ID == 0 {
		return bus2Start, bus2ID, bus2offset
	}

	// let's step bus 1, and see when that step plus bus 2's offset matches bus 2 id
	for i := 1; i < bus1ID*bus2ID; i++ {
		// this is the step that bus1 is going to do each iteration. Start + offset + i * bus id.
		nStart := bus1Start + (i * bus1ID) - bus1offset

		// this takes the number that bus1 arrived on, subtracts bus2 offset, and bus2 start, to check for mod bus 2 id.
		n := nStart - bus2Start + bus2offset
		// if we have a mod 0, that means that the buses will overlap at this time, starts and offsets accounted for.
		m := n % bus2ID
		if m == 0 {
			return nStart, bus1ID * bus2ID, 0
		}
	}
	return 0, 0, 0
}

// earliestTime returns the time at which all buses leave according to their offsets. The structure of the map is
// map[offset]interval.
func earliestTime(in map[int][]int) int {
	lastBusStart, lastBusID, lastBusOffset := 0, 0, 0
	for offset, buses := range in {
		for _, bus := range buses {
			if lastBusID == 0 {
				lastBusStart, lastBusID, lastBusOffset = 0, bus, offset
				continue
			}
			inStart, inID, inOffset := togetherBus(lastBusStart, lastBusID, lastBusOffset, 0, bus, offset)
			lastBusStart, lastBusID, lastBusOffset = inStart, inID, inOffset
		}
	}
	return lastBusStart
}

func t2formatInput(in []string) map[int][]int {
	b := strings.Split(in[1], ",")
	buses := make(map[int][]int, 0)
	// let's turn each bus id into an offset with id.
	for idx, bus := range b {
		// we don't care about exes.
		if bus == "x" {
			continue
		}
		busNum, err := strconv.Atoi(bus)
		if err != nil {
			fmt.Printf("buses: could not turn bus number %s into int: %s\n", bus, err)
			continue
		}
		// if the index is bigger than the bus id, we can reduce the index to be only the mod. If idx is 50 and bus id
		// is 19, then 50 % 19 and 12 % 19 give us the same result, so let's use 12 instead.
		if idx >= busNum {
			idx = idx % busNum
		}

		// after reducing the indexes, a bunch of bus ids will be on the same difference.
		buses[idx] = append(buses[idx], busNum)
	}

	// let's see if the 0 offset value has an offset, at which point it can be moved there
	for i, v := range buses[0] {
		if _, ok := buses[v]; ok {
			buses[v] = append(buses[v], v)
			buses[0] = remove(buses[0], i)
		}
	}
	// if we don't have a bus for an offset, get rid of the offset.
	if len(buses[0]) == 0 {
		delete(buses, 0)
	}

	return buses
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
