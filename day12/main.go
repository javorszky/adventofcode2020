package day12

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	filename = "day12/input.txt"
	east     = "E"
	south    = "S"
	west     = "W"
	north    = "N"
	forward  = "F"
	right    = "R"
	left     = "L"
	rotation = "ESWN"
)

var rotMap = map[int]int{
	90:  270,
	180: 180,
	270: 90,
}

func Tasks() {
	task1()
	task2()
}

type ship struct {
	X, Y      int
	Direction string
}

func task1() {
	s := ship{
		X:         0,
		Y:         0,
		Direction: east,
	}

	ins := getInputs()

	newS := navigate(s, ins)

	fmt.Printf("\nDay 12 task 1: the total movement of the ship was %d\n", manhattanDistance(newS))
}

func navigate(s ship, ins []string) ship {
	for _, cmd := range ins {
		s = moveShip(s, cmd)
	}
	return s
}

func manhattanDistance(s ship) int {
	return abs(s.X) + abs(s.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveShip(s ship, instruction string) ship {
	c := instruction[0:1]
	amnt, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(fmt.Sprintf("moveship inst: %s: %s", instruction, err))
	}

	switch c {
	case forward:
		return moveShip(s, fmt.Sprintf("%s%d", s.Direction, amnt))
	case right, left:
		return ship{
			Direction: newDirection(s.Direction, c, amnt),
			X:         s.X,
			Y:         s.Y,
		}
	case east:
		return ship{
			Direction: s.Direction,
			X:         s.X + amnt,
			Y:         s.Y,
		}
	case south:
		return ship{
			Direction: s.Direction,
			X:         s.X,
			Y:         s.Y + amnt,
		}
	case west:
		return ship{
			Direction: s.Direction,
			X:         s.X - amnt,
			Y:         s.Y,
		}
	case north:
		return ship{
			Direction: s.Direction,
			X:         s.X,
			Y:         s.Y - amnt,
		}
	default:
		fmt.Printf("uh... could not move ship anywhere... instruction: %s\n", instruction)
		return s
	}
}

// newDirection receives a base (E, S, W, N), and an amount (90, 180, 270), and sends back a new direction.
func newDirection(base, dir string, amnt int) string {
	if dir == "L" {
		return newDirection(base, "R", rotMap[amnt])
	}
	idx := strings.Index(rotation, base)
	if idx == -1 {
		panic(fmt.Sprintf("invalid base: %s", base))
	}
	newIDX := (len(rotation) + idx + (amnt / 90)) % 4
	return rotation[newIDX : newIDX+1]
}

func task2() {
	_ = getInputs()
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}