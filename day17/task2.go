package day17

import (
	"fmt"
	"strings"
)

const (
	wStart = 0
)

// x (left/right), y (front/back), z (up/down), w (?/?) coordinates. The starting state is all on z=0, with top left
// tile on x=0 and y=0.
type hyperGrid map[int]map[int]map[int]map[int]string

// Neighbours will return hyperGrid where coordinates of x, y, z, w are at most one off from current coördinate.
func (g hyperGrid) neighbours(x, y, z, w int) hyperGrid {
	n := make(hyperGrid, 0)
	// iterate on x
	for i := -1; i <= 1; i++ {
		n[x+i] = make(map[int]map[int]map[int]string, 0)
		// iterate on y
		for j := -1; j <= 1; j++ {
			n[x+i][y+j] = make(map[int]map[int]string, 0)
			// iterate on z
			for k := -1; k <= 1; k++ {
				n[x+i][y+j][z+k] = make(map[int]string, 0)
				for l := -1; l <= 1; l++ {
					// but skip the coordinate we want the neighbours for.
					if i == 0 && j == 0 && k == 0 {
						continue
					}
					n[x+i][y+j][z+k][w+l] = g.state(x+i, y+j, z+k, w+l)
				}
			}
		}
	}
	return n
}

// state returns the conway cube at coordinates x, y, z. It's either active, or inactive. A point that isn't recorded in
// the world is inactive.
func (g hyperGrid) state(x, y, z, w int) string {
	state := "."
	// y/z plane does not exist at coordinate x, there are no cubes recorded on it, therefore all of them are inactive.
	if g[x] == nil {
		return state
	}

	// z column does not exist at coordinates x/y, therefore all cubes there are inactive.
	if g[x][y] == nil {
		return state
	}

	// w column does not exists at x/y/z coördinates, therefore all cubes are inactive.
	if g[x][y][z] == nil {
		return state
	}

	s, ok := g[x][y][z][w]

	// there was no cube in the column at that coordinate, therefore the cube that was supposed to be there is inactive.
	if !ok {
		return state
	}

	// return the state of the found and recorded cube.
	return s
}

// setStateAt sets the value at xyz coordinates to state.
func (g hyperGrid) setStateAt(x, y, z, w int, state string) hyperGrid {
	if g[x] == nil {
		g[x] = make(map[int]map[int]map[int]string, 0)
	}
	if g[x][y] == nil {
		g[x][y] = make(map[int]map[int]string, 0)
	}
	if g[x][y][z] == nil {
		g[x][y][z] = make(map[int]string, 0)
	}

	g[x][y][z][w] = state

	return g
}

// edges returns a list of coordinates that have at least one neighbour that exists in our current hyperGrid. This is useful
// to walk when we want to consider the next cycle.
//
// In essence the list is the unique union of neighbours of all known cubes.
func (g hyperGrid) edges() map[int]map[int]map[int]map[int]struct{} {
	e := make(map[int]map[int]map[int]map[int]struct{}, 0)
	// extract x coördinate, loop through everything with x anchored.
	for x, yzw := range g {
		// extract y coördinate, loop through everything with x and y anchored.
		for y, zw := range yzw {
			// extract z coördinate, loop through everything with x, y, and z anchored.
			for z, wCubes := range zw {
				// extract w coördinate from the list of cubes in w where x, y, and z are all anchored.
				for w := range wCubes {
					n := g.neighbours(x, y, z, w)
					for nx, nyzw := range n {
						if e[nx] == nil {
							e[nx] = make(map[int]map[int]map[int]struct{}, 0)
						}
						for ny, nzw := range nyzw {
							if e[nx][ny] == nil {
								e[nx][ny] = make(map[int]map[int]struct{}, 0)
							}
							for nz, wCubes := range nzw {
								if e[nx][ny][nz] == nil {
									e[nx][ny][nz] = make(map[int]struct{}, 0)
								}
								for nw := range wCubes {
									e[nx][ny][nz][nw] = struct{}{}
								}
							}
						}
					}
					// neighbours leaves out the actual cube, but we need it here.
					e[x][y][z][w] = struct{}{}
				}
			}
		}
	}
	return e
}

// cycle will apply the rules of conway cubes to the current state of the hyperGrid, and returns a new hyperGrid where all the
// coördinates were applied, and the hyperGrid has grown.
func (g hyperGrid) cycle() hyperGrid {
	newGrid := make(hyperGrid, 0)
	// step 1: get the edges. This is where the hyperGrid is going to grow next.
	for x, yzw := range g.edges() {
		for y, zw := range yzw {
			for w, wStructs := range zw {
				for z := range wStructs {
					// step 2: get the neighbours of all the cubes in the edged (grown) space.
					n := g.neighbours(x, y, z, w)

					// step 3: figure out what the new state should be at a given point in the edged (grown) space.
					newState := hyperNext(g.state(x, y, z, w), n)

					// step 4: set that state into the hyperGrid.
					newGrid.setStateAt(x, y, z, w, newState)
				}
			}
		}
	}

	// step 5: return the new hyperGrid at the end of the cycle.
	return newGrid
}

func (g hyperGrid) actives() int {
	actives := 0
	// count all the actives in the neighbours.
	for _, yzw := range g {
		for _, zw := range yzw {
			for _, wCubes := range zw {
				for _, cube := range wCubes {
					if cube == active {
						actives++
					}
				}
			}
		}
	}
	return actives
}

// hyperNext will take a current state, and its neighbours, and returns what the next state should be. The rules for
// the state change in the 2020 advent of code day 17 task 2 are:
//
// - If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the
//   cube becomes inactive.
// - If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube
//   remains inactive.
func hyperNext(state string, neighbours hyperGrid) string {
	actives := neighbours.actives()

	switch state {
	case active:
		if actives == 2 || actives == 3 {
			return active
		}
	case inactive:
		if actives == 3 {
			return active
		}
	}
	return inactive
}

func task2() {
	// load our initial state into the world.
	starter := getInputs()
	world := make(hyperGrid, 0)
	for x, ys := range starter {
		for y, state := range strings.Split(ys, "") {
			world = world.setStateAt(x, y, zStart, wStart, state)
		}
	}

	for i := 0; i < cycles; i++ {
		world = world.cycle()
	}

	fmt.Printf("Day 17 task 2: after 6 cycles there are %d active cubes in the grid\n", world.actives())
}
