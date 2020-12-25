package day20

import (
	"errors"
	"fmt"
	"strings"
)

// image2 holds tiles in some configuration. W and H are the width and height of the image2. The map is a 0 based
// coordinate, so a 6x6 image2 will have its tiles from 0-5 / 0-5.
type image2 struct {
	W, H  int
	Tiles map[int]map[int]tilev2
}

func fitTileIntoImage2(i image2, tileNumber int, tss tileSets2) (image2, error) {
	var newImage image2
	x, y, err := getXY2(i, tileNumber)
	if err != nil {
		return i, fmt.Errorf("fitTileIntoImage: %w", err)
	}

	for tileID, ts := range tss {
		// we're using a tileset, we can't use the same tileset for the next move, so remove it from the possibilities.
		ntss := make(tileSets2, 0)
		for tid, tsToCopy := range tss {
			if tid == tileID {
				continue
			}
			ntss[tid] = tsToCopy
		}

		// create a copy of the current tileset to check.
		leftTilesToCheck := make([]tilev2, 0)
		var topTilesToCheck []tilev2
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
			topTilesToCheck = make([]tilev2, len(leftTilesToCheck))
			copy(topTilesToCheck, leftTilesToCheck)
		} else {
			topTilesToCheck = make([]tilev2, len(ts))
			copy(topTilesToCheck, ts)
		}

		tilesToTry := make([]tilev2, 0)
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
			tilesToTry = make([]tilev2, len(topTilesToCheck))
			copy(tilesToTry, topTilesToCheck)
		}

		for _, tileToTry := range tilesToTry {
			if i.Tiles[x] == nil {
				i.Tiles[x] = make(map[int]tilev2, 0)
			}
			i.Tiles[x][y] = tileToTry
			if len(ntss) > 0 {
				newImage, err = fitTileIntoImage2(i, tileNumber+1, ntss)
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

func getXY2(i image2, n int) (int, int, error) {
	if n < 1 || n > (i.W*i.H) {
		return 0, 0, errors.New(fmt.Sprintf("getXY: out of bounds: %d, min 1, max %d", n, i.W*i.H))
	}
	// x is the modulo of the width
	return (n - 1) % i.W, (n - 1) / i.H, nil
}

func stitchImage(img image2) []string {
	stitchedImage := make([]string, 0)

	// image has rows and columns of tiles. A 5x5 image will have 5 rows, each row will hae 5 columns, and each tile
	// will have n lines. The order of indexes is [row][line][col]
	// create a 3d slice as long as however many rows there are in the tiles
	stitchedImageRowBatch := make([][][]string, len(img.Tiles))

	// for ever row of tiles in image.
	for rowIDX, cols := range img.Tiles {
		// create the second level of the slice, this time the index is the line number
		stitchedImageRowBatch[rowIDX] = make([][]string, len(cols[0].Content))
		// for every tile in the row.
		for columnIDX, tile := range cols {
			// for every line in a tile's content
			for lineIDX, line := range tile.Content {
				if stitchedImageRowBatch[rowIDX][lineIDX] == nil {
					stitchedImageRowBatch[rowIDX][lineIDX] = make([]string, len(cols))
				}
				stitchedImageRowBatch[rowIDX][lineIDX][columnIDX] = line
			}
		}
	}

	for _, row := range stitchedImageRowBatch {
		for _, line := range row {
			stitchedImage = append(stitchedImage, strings.Join(line, ""))
		}
	}

	return stitchedImage
}
