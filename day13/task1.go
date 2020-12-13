package day13

import (
	"fmt"
	"strconv"
	"strings"
)

func task1() {
	_, _ = t1formatInput(getInputs())
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
