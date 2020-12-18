package day18

import (
	"fmt"
	"strconv"
	"strings"
)

func task1() {
	sum := 0
	for _, line := range getInputs() {
		sum = sum + evaluateExpression(line)
	}

	fmt.Printf("\nDay 18 task 1: the sum of all evaluations is %d\n", sum)

}

func evaluateExpression(s string) int {
	inGroup := false
	depth := 0
	evaluated := 0
	var sb strings.Builder
	var f func(int) int
	f = func(c int) int {
		return c
	}

	for _, char := range s {
		switch string(char) {
		// if we encounter an opening parens
		case "(":
			// increment depth by one
			depth++
			// if we're not in a group yet, start by adding the group
			if !inGroup {
				inGroup = true
				continue
			}
			sb.WriteString(string(char))
		case ")":
			depth--
			if depth == 0 {
				inGroup = false
				// we've just closed an outermost group
				n := evaluateExpression(sb.String())
				evaluated = f(n)
				sb.Reset()
				continue
			}
			sb.WriteString(string(char))
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// this is a number, so we're either going to add it to the string builder, or evaluate
			if inGroup {
				sb.WriteString(string(char))
				continue
			}
			i, err := strconv.Atoi(string(char))
			if err != nil {
				panic(fmt.Sprintf("can't turn '%s' into int: %s", string(char), err))
			}
			evaluated = f(i)
		case "+":
			if inGroup {
				sb.WriteString(string(char))
				continue
			}
			f = func(i int) int {
				return evaluated + i
			}
		case "*":
			if inGroup {
				sb.WriteString(string(char))
				continue
			}
			f = func(i int) int {
				return evaluated * i
			}
		}
	}
	return evaluated
}
