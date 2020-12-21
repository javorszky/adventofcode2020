package day20

import (
	"errors"
	"fmt"
)

// image holds tiles in some configuration. W and H are the width and height of the image. The map is a 0 based
// coordinate, so a 6x6 image will have its tiles from 0-5 / 0-5.
type image struct {
	W, H  int
	Tiles map[int]map[int]tile
}

func fitTileIntoImage(i image, tileNumber int, tss tileSets) (image, error) {
	var newImage image
	x, y, err := getXY(i, tileNumber)
	if err != nil {
		return i, fmt.Errorf("fitTileIntoImage: %w", err)
	}

	for tileID, ts := range tss {
		// we're using a tileset, we can't use the same tileset for the next move, so remove it from the possibilities.
		ntss := make(tileSets, 0)
		for tid, tsToCopy := range tss {
			if tid == tileID {
				continue
			}
			ntss[tid] = tsToCopy
		}

		// create a copy of the current tileset to check.
		leftTilesToCheck := make([]tile, 0)
		var topTilesToCheck []tile
		// do we need to check left?
		if x > 0 {
			// get the right side of the tile to the left
			leftTile, ok := i.Tiles[x-1][y]
			if !ok {
				return i, errors.New(fmt.Sprintf("can't find tile to the left of where we wantx to current one: x: %d, y: %d", x, y))
			}

			for _, tileFromSet := range ts {
				// compare the right side of leftTile to the left side of the current tile
				if leftTile.Right == tileFromSet.Left {
					leftTilesToCheck = append(leftTilesToCheck, tileFromSet)
				}
			}
			topTilesToCheck = make([]tile, len(leftTilesToCheck))
			copy(topTilesToCheck, leftTilesToCheck)
		} else {
			topTilesToCheck = make([]tile, len(ts))
			copy(topTilesToCheck, ts)
		}

		tilesToTry := make([]tile, 0)
		// do we need to check top?
		if y > 0 {
			// get the bottom side of the tile to the top
			topTile, ok := i.Tiles[x][y-1]
			if !ok {
				return i, errors.New(fmt.Sprintf("can't find tile to the left of where we wantx to current one: x: %d, y: %d", x, y))
			}

			for _, tileFromSet := range topTilesToCheck {
				// compare the right side of leftTile to the left side of the current tile
				if topTile.Bottom == tileFromSet.Top {
					tilesToTry = append(tilesToTry, tileFromSet)
				}
			}
		} else {
			tilesToTry = make([]tile, len(topTilesToCheck))
			copy(tilesToTry, topTilesToCheck)
		}

		for _, tileToTry := range tilesToTry {
			if i.Tiles[x] == nil {
				i.Tiles[x] = make(map[int]tile, 0)
			}
			i.Tiles[x][y] = tileToTry
			if len(ntss) > 0 {
				newImage, err = fitTileIntoImage(i, tileNumber+1, ntss)
				if err == nil {
					return newImage, nil
				}
			} else {
				return i, nil
			}
		}
	}

	return newImage, fmt.Errorf("reached the end with no more possible fit")
}

func getXY(i image, n int) (int, int, error) {
	if n < 1 || n > (i.W*i.H) {
		return 0, 0, errors.New(fmt.Sprintf("getXY: out of bounds: %d, min 1, max %d", n, i.W*i.H))
	}
	// x is the modulo of the width
	return (n - 1) % i.W, (n - 1) / i.H, nil
}
