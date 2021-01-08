package day23

import (
	"container/ring"
	"fmt"
)

func task2LinkedList() {
	//start := getInputs()
	//
	//task2GetProducts(getInputs(), 10)
}

func task2GetProducts(inputs []int, totalCups, steps int) int {
	r := generateRing(inputs, totalCups)
	for i := 0; i < steps; i++ {
		r = step(r)
	}

	a := *r.Value.(cup).Addresses
	p1 := a[1].Next()
	p1v := p1.Value.(cup).Value
	p2 := p1.Next()
	p2v := p2.Value.(cup).Value

	return p1v * p2v
}

func generateRing(start []int, total int) *ring.Ring {
	if total < len(start) {
		panic(fmt.Sprintf("total number of cups (%d) is less than the starting one (%d)", total, len(start)))
	}
	r := ring.New(total)
	addresses := make(map[int]*ring.Ring, 0)
	for _, v := range start {
		r.Value = cup{
			Value:     v,
			Addresses: &addresses,
		}
		addresses[v] = r
		r = r.Next()
	}

	for i := len(start) + 1; i <= total; i++ {
		r.Value = cup{
			Value:     i,
			Addresses: &addresses,
		}
		addresses[i] = r
		r = r.Next()
	}

	return r
}
