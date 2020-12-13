package day13

import (
	"fmt"
	"strconv"
	"strings"
)

// the numbers are all primes.
func task2() {
	e := t2formatInput(getInputs())
	//
	fmt.Printf("the map of offsets and bus ids: %#v\n", e)
	//
	//final := earliestTime(e)
	//
	//for offset, intervals := range e {
	//	for _, int := range intervals {
	//		fmt.Printf("%d mod %d is %d, should be %d\n", final, int, final%int, offset)
	//	}
	//}
	//earliestTime(e)
	//pf := getPrimeFactors(98237489723)
	//fmt.Printf("prime factors of %d are %#v\n", 98237489723, pf)

	//playground()

	pg2()
}

// earliestTime returns the time at which all buses leave according to their offsets. The structure of the map is
// map[offset]interval.
func earliestTime(in map[int][]int) int {
	// let's find the largest interval to do the least amount of calculations.
	largest := 0
	offsetLargest := 0
	for offset, intervals := range in {
		product := 1
		for _, int := range intervals {
			product = product * int
		}
		if product > largest {
			largest = product
			offsetLargest = offset
		}
	}
	fmt.Printf("num %d at %d is good\n", largest, offsetLargest)
	//n := largest + offsetLargest
	//for _, c := range in[offsetLargest] {
	//	fmt.Printf("%d mod %d is %d\n", n, c, n%c)
	//}
	//delete(in, offsetLargest)
	//busFactors := make(map[int]int, 0)
	//for offset, intervals := range in {
	//	for _, int := range intervals {
	//		n2 := int
	//		needMod := offset
	//		for i := 0; i < n2; i++ {
	//			m := (largest * i) % n2
	//			if needMod == m {
	//				busFactors[int] = i
	//				break
	//			}
	//		}
	//	}
	//}
	//fmt.Printf("bus factor is %v\n", busFactors)
	//final := largest
	//for n2, factor := range busFactors {
	//	final = final * (n2 + factor)
	//}

	return 19
}

func t2formatInput(in []string) map[int][]int {
	b := strings.Split(in[1], ",")
	buses := make(map[int][]int, 0)
	for idx, bus := range b {
		if bus == "x" {
			continue
		}
		busNum, err := strconv.Atoi(bus)
		if err != nil {
			fmt.Printf("buses: could not turn bus number %s into int: %s\n", bus, err)
			continue
		}
		if idx >= busNum {
			idx = idx % busNum
		}

		buses[idx] = append(buses[idx], busNum)
	}

	// let's see if the 0 offset value has an offset, at which point it can be moved there
	for i, v := range buses[0] {
		if _, ok := buses[v]; ok {
			buses[v] = append(buses[v], v)
			buses[0] = remove(buses[0], i)
		}
	}
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
	base := 823 * 29 * 37 * 19
	baseMod := 19
	factor := 1
	//m := map[int]int{
	//	9: 41,
	//}
	//for newMod, newNum := range m {
	//	for i := 0; i < newNum; i++ {
	//		n := base*factor*i + baseMod
	//		m := n % newNum
	//		t := ""
	//		if m == newMod {
	//			factor = factor * (i + newNum)
	//			t = "<---"
	//		}
	//		fmt.Printf("%3d, %3d, %3d %s\n", m, newMod, i, t)
	//	}
	//}
	//
	//fmt.Printf("new one, factor is %d\n", factor)

	newNum := 41
	newMod := 9
	for i := 0; i < newNum; i++ {
		n := base*factor*i + baseMod
		m := n % newNum
		t := ""
		fmt.Printf("%3d, %3d, %3d %s\n", m, newMod, i, t)
	}

	for i := 0; i < newNum; i++ {
		n := base*i + baseMod
		m := n % newNum
		t := ""
		if m == newMod {
			factor = i + newNum
			t = "<---"
		}
		fmt.Printf("%3d, %3d, %3d %s\n", m, newMod, i, t)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
