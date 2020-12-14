package day13

import (
	"fmt"
	"strconv"
	"strings"
)

// the numbers are all primes.
func task2() {
	_ = earliestTime(t2formatInput(getInputs()))

	pg2()
}

// togetherBus takes two buses, and returns a new one that envelops both of them.
func togetherBus(bus1Start, bus1ID, bus1offset, bus2Start, bus2ID, bus2offset int) (int, int, int) {
	if bus1ID == 0 {
		return bus2Start, bus2ID, bus2offset
	}
	// let's step bus 1, and see when that step plus bus 2's offset matches bus 2 id
	for i := 0; i < bus1ID*bus2ID; i++ {
		nStart := bus1Start + (i * bus1ID) + bus1offset
		n := nStart - bus2Start - bus2offset
		if n%bus2ID == 0 {
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
			lastBusStart, lastBusID, lastBusOffset = togetherBus(lastBusStart, lastBusID, lastBusOffset, 0, bus, offset)
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

func playground() {
	base := 823 * 29 * 37 * 19
	modsNums := map[int]int{
		//2: 17,
		//4: 23,
		9: 41,
	}
	baseMod := 19
	factor := 1
	for newMod, newNumber := range modsNums {
		//newMod = abs((newMod - baseMod) % newNumber)
		for j := 1; j < newNumber*3; j++ {
			m := ((base * j * (factor + newNumber)) + baseMod) % newNumber
			fmt.Printf("(%d * %d) %% %d = %d (%d)\n", base*factor, j, newNumber, m, newMod)
			//fmt.Printf("base %d, basemod %d, iter %d, base*iter %d, base*iter+basemod %d, mod newnum %d, mod %d, newmod %d\n", base, baseMod, j, base*j, base*j+baseMod, newNumber, m, newMod)
			if m == newMod {
				fmt.Printf("new factor: %d -> %d * %d = %d \n", factor, factor, j, factor*j)
				factor = factor * j
				break
			}
		}
	}

	for i := 1; i < 44; i++ {
		n := (base * factor * i) + baseMod
		m823 := n % 823
		m29 := n % 29
		m37 := n % 37
		m19 := n % 19
		m41 := n % 41
		m23 := n % 23
		m17 := n % 17
		fmt.Printf("n: %d, factor: %d, i: %d, 823 (%d - %d), 29 (%d - %d), 37 (%d - %d), 19 (%d - %d), 41 (%d - %d), 23 (%d - %d), 17 (%d - %d)\n", n, factor, i, m823, 19, m29, 19, m37, 19, m19, 0, m41, 9, m23, 4, m17, 2)
	}

	//fx := base*factor + baseMod
	//
	//fmt.Printf("%d mod %d is %d, should be %d\n", fx, 823, (fx)%823, baseMod)
	//fmt.Printf("%d mod %d is %d, should be %d\n", fx, 29, (fx)%29, baseMod)
	//fmt.Printf("%d mod %d is %d, should be %d\n", fx, 37, (fx)%37, baseMod)
	//fmt.Printf("%d mod %d is %d, should be %d\n", fx, 19, (fx)%19, baseMod)
	//fmt.Printf("%d mod %d is %d, should be %d\n", fx, 41, (fx)%41, 9)
	//fmt.Printf("%d mod %d is %d, should be %d\n", fx, 23, (fx)%23, 4)

}

func pg2() {
	return
	base := 41

	newNum := 823
	newMod := 19

	for i := 0; i < 10; i++ {
		n := 842 + (base * newNum)
		m := n % newNum
		t := ""
		if m == newMod {
			t = "<---"
			fmt.Printf("mod: %3d, wantmod: %3d, iter: %3d, number: %5d %s\n", m, newMod, i, n, t)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
