package day22

import (
	"strconv"
	"strings"
)

func task2() {
	//plays := make(map[string]struct{}, 0)
	//p1, p2 := getInputs()
}

func playString(p1, p2 []int) string {
	var sb strings.Builder
	stringHelper := make([]string, 0)
	for _, v := range p1 {
		stringHelper = append(stringHelper, strconv.Itoa(v))
	}
	sb.WriteString(strings.Join(stringHelper, ","))
	sb.WriteString("|")
	stringHelper = make([]string, 0)
	for _, v := range p2 {
		stringHelper = append(stringHelper, strconv.Itoa(v))
	}
	sb.WriteString(strings.Join(stringHelper, ","))

	return sb.String()
}
