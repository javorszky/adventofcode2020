package day17

// x (left/right), y (front/back), z (up/down) coordinates. The starting state is all on z=0, with top left tile on
// x=0 and y=0.
type grid map[int]map[int]map[int]string

// Neighbours will return a slice of strings up to 27 long.
func (g grid) neighbours(x, y, z int) grid {
	n := make(grid, 0)
	// iterate on x
	for i := -1; i <= 1; i++ {
		n[x+i] = make(map[int]map[int]string, 0)
		// iterate on y
		for j := -1; j <= 1; j++ {
			n[x+i][y+j] = make(map[int]string, 0)
			// iterate on z
			for k := -1; k <= 1; k++ {
				// but skip the coordinate we want the neighbours for.
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				n[x+i][y+j][z+k] = g.state(x+i, y+j, z+k)
			}
		}
	}
	return n
}

// state returns the conway cube at coordinates x, y, z. It's either active, or inactive. A point that isn't recorded in
// the world is inactive.
func (g grid) state(x, y, z int) string {
	state := "."
	// y/z plane does not exist at coordinate x, there are no cubes recorded on it, therefore all of them are inactive.
	if g[x] == nil {
		return state
	}

	// z column does not exist at coordinates x/y, therefore all cubes there are inactive.
	if g[x][y] == nil {
		return state
	}

	// z column exists at x/y coordinates, let's grab the state at the specified z point.
	s, ok := g[x][y][z]

	// there was no cube in the column at that coordinate, therefore the cube that was supposed to be there is inactive.
	if !ok {
		return state
	}

	// return the state of the found and recorded cube.
	return s
}

// setStateAt sets the value at xyz coordinates to state.
func (g grid) setStateAt(x, y, z int, state string) grid {
	if g[x] == nil {
		g[x] = make(map[int]map[int]string, 0)
	}
	if g[x][y] == nil {
		g[x][y] = make(map[int]string, 0)
	}

	g[x][y][z] = state

	return g
}

type world struct {
	Grid grid
}

func task1() {
	_ = getInputs()
}
