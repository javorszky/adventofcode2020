package day8

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const filename = "day8/input.txt"

func Tasks() {
	task1()

	task2()
}

type infiniteLoop struct {
	Message string
}

func (i infiniteLoop) Error() string {
	return fmt.Sprintf("infinite loop error: %s", i.Message)
}

func task1() {
	// instructions line by line as a slice of strings.
	instructions := getInputs()

	// accumulator starts at 0.
	acc := 0

	// keep track of previous code executions so we know when we're about to hit an infinite loop.
	jumps := make(map[int]struct{}, 0)

	// kick off running the program on line 0.
	l := 0

	var infiniErr infiniteLoop
	var err error

	// create reusable function to execute the code. Accumulator and instruction set are outside of this func.

	for {
		acc, l, jumps, err = run(acc, l, jumps, instructions)
		if errors.As(err, &infiniErr) {
			fmt.Printf("day 8 task 1: infinite code current state of accumulator is %d", acc)
			os.Exit(0)
		}
		if err != nil {
			fmt.Printf("big oof, bad code: %s", err)
			os.Exit(1)
		}
	}
}

func t1Number(s string) (int, error) {
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, fmt.Errorf("failed to parse into number: %s: %w", s, err)
	}
	if string(s[0]) == "-" {
		return 0 - n, nil
	}
	return n, nil
}

func task2() {

}

func run(acc, line int, jumps map[int]struct{}, program []string) (int, int, map[int]struct{}, error) {
	var next int

	// have we run this previously?
	if _, ok := jumps[line]; ok == true {
		return acc, next, jumps, infiniteLoop{Message: fmt.Sprintf("we encountered line %d again", line)}
	}
	jumps[line] = struct{}{}

	inst := program[line]
	n, err := t1Number(inst[4:])
	if err != nil {
		return acc, next, jumps, fmt.Errorf("error parsing instruction number code: %w", err)
	}
	switch inst[:4] {
	case "acc ":
		acc = acc + n
		fallthrough
	case "nop ":
		next = line + 1
	case "jmp ":
		next = line + n
	default:
		return acc, next, jumps, fmt.Errorf("unknown instruction received: %s\nterminating...\n", inst)
	}

	return acc, next, jumps, nil
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// strip the trailing newline
	sData := strings.TrimRight(string(data), "\n")

	return strings.Split(sData, "\n")
}
