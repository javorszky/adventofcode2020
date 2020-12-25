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
	row, col, err := getXY2(i, tileNumber)
	if err != nil {
		return i, fmt.Errorf("fitTileIntoImage: %w", err)
	}

	indent := strings.Repeat("--", tileNumber)
	fmt.Printf("%s----\n%srow: %d\n%scol: %d\n%s----\n", indent, indent, row, indent, col, indent)

	for tileID, ts := range tss {
		fmt.Printf("%strying tileset for tile %s\n", indent, tileID)
		// we're using a tileset, we can't use the same tileset for the next move, so remove it from the possibilities.
		ntss := make(tileSets2, 0)
		for tid, tsToCopy := range tss {
			if tid == tileID {
				fmt.Printf("%sremoving tileset for %s to try\n", indent, tid)
				continue
			}
			ntss[tid] = tsToCopy
		}

		// create a copy of the current tileset to check.
		leftTilesToCheck := make([]tilev2, 0)
		var topTilesToCheck []tilev2
		// do we need to check left?
		if col > 0 {
			// get the right side of the tile to the left
			leftTile, ok := i.Tiles[row][col-1]
			if !ok {
				return i, errors.New(fmt.Sprintf("can't find tile to the left of where we wantx to current one: row: %d, col: %d", row, col))
			}

			for _, tileFromSet := range ts {
				// compare the right side of leftTile to the left side of the current tile
				if leftTile.Right == tileFromSet.Left {
					fmt.Printf("%scomparing tile\n%s%s Right %s and\n%s%s Left  %s\n", indent, indent, leftTile.ID, leftTile.Right, indent, tileFromSet.ID, tileFromSet.Left)
					leftTilesToCheck = append(leftTilesToCheck, tileFromSet)
				}
			}
			// we have whittled down tiles from the tileset to those that can match whatever is left of the current
			// one, so create a new slice with the same length as the whittled down one, and copy it over.
			topTilesToCheck = make([]tilev2, len(leftTilesToCheck))
			copy(topTilesToCheck, leftTilesToCheck)
		} else {
			fmt.Printf("%sat left edge, not checking left/right match\n", indent)

			// we have not whittled down any tiles from the tileset because there are no tiles to the left of this one,
			// so create a slice to hold tiles to check for the top the same length as the entire tileset, and copy the
			// tileset into it.
			topTilesToCheck = make([]tilev2, len(ts))
			copy(topTilesToCheck, ts)
		}

		// create a slice that will hold tiles that match both the left and the top tiles, if applicable.
		tilesToTry := make([]tilev2, 0)
		// do we need to check top?
		if row > 0 {
			// get the bottom side of the tile to the top
			topTile, ok := i.Tiles[row-1][col]
			if !ok {
				return i, errors.New(fmt.Sprintf("can't find tile to the left of where we wantx to current one: row: %d, col: %d", row, col))
			}

			for _, tileFromSet := range topTilesToCheck {
				// compare the bottom side of top tile to the top side of the current tile
				if topTile.Bottom == tileFromSet.Top {
					fmt.Printf("%scomparing tile\n%s%s Top    %s and\n%s%s Bottom %s\n", indent, indent, topTile.ID, topTile.Bottom, indent, tileFromSet.ID, tileFromSet.Top)
					tilesToTry = append(tilesToTry, tileFromSet)
				}
			}
		} else {
			fmt.Printf("%sat top edge, not checking top/bottom match\n", indent)
			tilesToTry = make([]tilev2, len(topTilesToCheck))
			copy(tilesToTry, topTilesToCheck)
		}

		for _, tileToTry := range tilesToTry {
			if i.Tiles[row] == nil {
				i.Tiles[row] = make(map[int]tilev2, 0)
			}

			fmt.Printf("%strying %s at %d:%d\n", indent, tileToTry.ID, row, col)
			i.Tiles[row][col] = tileToTry
			if len(ntss) > 0 {
				newImage, err = fitTileIntoImage2(i, tileNumber+1, ntss)
				if err == nil {
					fmt.Printf("%sFound it!\n", indent)
					return newImage, nil
				}
			} else {
				fmt.Printf("%snope...\n", indent)
				return i, nil
			}
		}
	}
	fmt.Printf("%sreturning because no more tiles left to try in current tilesets\n", indent)

	return newImage, fmt.Errorf("reached the end with no more possible fit")
}

// getXY2 returns the x, y, err coordinates.
func getXY2(i image2, n int) (int, int, error) {
	if n < 1 || n > (i.W*i.H) {
		return 0, 0, errors.New(fmt.Sprintf("getXY: out of bounds: %d, min 1, max %d", n, i.W*i.H))
	}
	// x is the modulo of the width
	return (n - 1) / i.H, (n - 1) % i.W, nil
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
