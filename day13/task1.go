package day13

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func task1() {
	t, buses := t1formatInput(getInputs())
	diffs := make([]int, 0)
	diffMap := make(map[int]int, 0)
	for _, bus := range buses {
		d := t1next(t, bus)
		diffs = append(diffs, d)
		diffMap[d] = bus
	}
	sort.Ints(diffs)
	smolest := diffs[0]
	fmt.Printf("\nDay 13 task 1: the product of the time to wait and bus number is %d\n", smolest*diffMap[smolest])
}

func t1formatInput(in []string) (int, []int) {
	b := strings.Split(in[1], ",")
	buses := make([]int, 0)
	for _, bus := range b {
		if bus == "x" {
			continue
		}
		busNum, err := strconv.Atoi(bus)
		if err != nil {
			fmt.Printf("buses: could not turn bus number %s into int: %s\n", bus, err)
			continue
		}
		buses = append(buses, busNum)
	}

	ts, err := strconv.Atoi(in[0])
	if err != nil {
		fmt.Printf("time: could not turn bus number %s into int: %s\n", in[0], err)
	}

	return ts, buses
}

// t1next takes a time stamp and an interval, and returns how far in the future the next time is from t where interval
// is divisible by that time.
func t1next(t int, interval int) int {
	var c int
	if t%interval == 0 {
		return 0
	}
	c = t / interval
	return ((c + 1) * interval) - t
}
