package day12

import (
	"fmt"
	"strconv"
)

type ship2 struct {
	WX, WY, SX, SY int
}

func (s ship2) posX() int {
	return s.SX
}

func (s ship2) posY() int {
	return s.SY
}

func task2() {
	s := ship2{
		WX: 10,
		WY: -1,
		SX: 0,
		SY: 0,
	}

	ins := getInputs()

	newS := navigate2(s, ins)

	fmt.Printf("Day 12 task 2: the total movement of the ship was %d\n", manhattanDistance(newS))
}

func navigate2(s ship2, instruction []string) ship2 {
	for _, cmd := range instruction {
		s = moveShip2(s, cmd)
	}
	return s
}

// newDirection2 receives a ship, a direction (R or L), and an amount (90, 180, 270), and returns a new ship with
// rotated waypoint coordinates.
func newDirection2(s ship2, dir string, amnt int) ship2 {
	if dir == "L" {
		return newDirection2(s, "R", rotMap[amnt])
	}
	wx, wy := s.WX, s.WY
	switch amnt {
	case 90:
		s.WX = -wy
		s.WY = wx
	case 180:
		s.WX = -wx
		s.WY = -wy
	case 270:
		s.WX = wy
		s.WY = -wx
	}

	return s
}

func moveShip2(s ship2, instruction string) ship2 {
	c := instruction[0:1]
	amnt, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(fmt.Sprintf("moveship inst: %s: %s", instruction, err))
	}

	switch c {
	case forward:
		return ship2{
			WX: s.WX,
			WY: s.WY,
			SX: s.SX + s.WX*amnt,
			SY: s.SY + s.WY*amnt,
		}
	case right, left:
		return newDirection2(s, c, amnt)
	case east:
		return ship2{
			SX: s.SX,
			SY: s.SY,
			WX: s.WX + amnt,
			WY: s.WY,
		}
	case south:
		return ship2{
			SX: s.SX,
			SY: s.SY,
			WX: s.WX,
			WY: s.WY + amnt,
		}
	case west:
		return ship2{
			SX: s.SX,
			SY: s.SY,
			WX: s.WX - amnt,
			WY: s.WY,
		}
	case north:
		return ship2{
			SX: s.SX,
			SY: s.SY,
			WX: s.WX,
			WY: s.WY - amnt,
		}
	default:
		fmt.Printf("uh... could not move ship anywhere... instruction: %s\n", instruction)
		return s
	}
}
