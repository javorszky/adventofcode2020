package day12

type ship2 struct {
	WX, WY, SX, SY int
}

func task2() {
	_ = getInputs()
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
