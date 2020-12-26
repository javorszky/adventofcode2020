package day20

type tileSet2 []tilev2

type tileSets2 map[string]tileSet2

func (tss tileSets2) addTileSet(ts tileSet2) tileSets2 {
	tss[ts[0].ID[:4]] = ts
	return tss
}

// newTileSet2 takes a tile with "000" suffix, and generates all the variations of flips and rotations.
func newTileSet2(t tilev2) tileSet2 {
	tiles := make([]tilev2, 0, 8)

	// 1: save the origin 000
	tiles = append(tiles, t)

	// 2: save the rotated 001
	tr := t.rotate()
	tiles = append(tiles, tr)

	// 3: save the vflipped 100

	tv := t.flipV()
	tiles = append(tiles, tv)

	// 4: save the vflipped, rotated 101
	tvr := tv.rotate()
	tiles = append(tiles, tvr)

	// 5: save the hflipped, 010
	th := t.flipH()
	tiles = append(tiles, th)

	// 6: save the hflipped rotated 011
	thr := th.rotate()
	tiles = append(tiles, thr)

	// 7: save the vflipped hflipped 110
	tvh := th.flipV()
	tiles = append(tiles, tvh)

	// 8: save the vflipped hflipped rotated 111
	tvhr := tvh.rotate()
	tiles = append(tiles, tvhr)

	return tiles
}
