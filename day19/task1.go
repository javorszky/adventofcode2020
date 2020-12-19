package day19

import "fmt"

// task1 will do the calculations necessary to compute day 19 task 1:
//
// "Your goal is to determine the number of messages that completely match rule 0."
func task1() {
	rules, messages := getInputs()
	fmt.Printf("rules:\n%#v\n\nmessages:\n%#v\n", rules, messages)
}
