package day23

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"
)

type cup struct {
	Value     int
	Addresses *map[int]*ring.Ring
}

func task1LinkedList() {
	label := task1GetLabel(getInputs(), 100)
	fmt.Printf("\nDay 23 task 1 linked list: label is %s\n", label)
}

func task1GetLabel(inputs []int, steps int) string {
	r := ring.New(len(inputs))
	addresses := make(map[int]*ring.Ring, 0)
	for _, v := range inputs {
		r.Value = cup{
			Value:     v,
			Addresses: &addresses,
		}
		addresses[v] = r
		r = r.Next()
	}

	for i := 0; i < steps; i++ {
		r = step(r)
	}

	var sb strings.Builder
	rout := addresses[1]
	printRing("rout", rout)
	for i := 0; i < rout.Len(); i++ {
		sb.WriteString(strconv.Itoa(rout.Value.(cup).Value))
		rout = rout.Next()
	}
	return strings.TrimLeft(sb.String(), "1")
}

func step(r *ring.Ring) *ring.Ring {
	l := r.Len()
	subRing := r.Unlink(3)

	a := *r.Value.(cup).Addresses
	currentValue := r.Value.(cup).Value

	nextHelper := make(map[int]struct{}, 0)
	for i := 0; i < subRing.Len(); i++ {
		nextHelper[subRing.Value.(cup).Value] = struct{}{}
		subRing = subRing.Next()
	}

	v := currentValue
	for {
		v--
		if v < 1 {
			v = l
		}
		if _, ok := nextHelper[v]; !ok {
			break
		}

	}

	a[v].Link(subRing)
	return a[currentValue].Next()
}

func printRing(name string, r *ring.Ring) {
	fmt.Printf("printing %s\n", name)
	for i := 0; i < r.Len(); i++ {
		fmt.Printf("%d ", r.Value.(cup).Value)
		r = r.Next()
	}
	fmt.Printf("\ndone printing %s\n\n", name)
}
