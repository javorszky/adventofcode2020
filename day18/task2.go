package day18

import (
	"fmt"
	"strconv"
	"strings"
)

func task2() {
	sum := 0
	for _, line := range getInputs() {
		sum = sum + evaluateExpression2(line)
	}

	fmt.Printf("Day 18 task 2: the sum of all advanced evaluations is %d\n", sum)
}

// evaluateExpression2 evaluates a given e itermxpression where the addition takes precedence. As well as parentheses.
func evaluateExpression2(s string) int {
	inGroup := false
	depth := 0
	evaluated := 0
	var sb strings.Builder
	var f func(int) int
	f = func(c int) int {
		return c
	}

	multiplicationGroups := make([]int, 0)

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
				n := evaluateExpression2(sb.String())
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
			multiplicationGroups = append(multiplicationGroups, evaluated)
			f = func(c int) int {
				return c
			}
		}
	}
	multiplicationGroups = append(multiplicationGroups, evaluated)

	product := 1
	for _, n := range multiplicationGroups {
		product = product * n
	}

	return product
}
