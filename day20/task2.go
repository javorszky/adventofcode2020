package day20

import (
	"strings"
)

func task2() {
	tss := make(tileSets2, 0)
	// there are 144 tiles, so that's a 12x12 grid.
	for _, tileString := range getInputs() {
		tss.addTileSet(newTileSet2(parseTileTask2(tileString)))
	}
	i := image2{
		W:     12,
		H:     12,
		Tiles: make(map[int]map[int]tilev2, 0),
	}

	img, err := fitTileIntoImage2(i, 1, tss)

	if err != nil {
		panic("\nDay 20 task 2 failed")
	}

	stitchedImage := stitchImage(img)
	lineLength := len(stitchedImage[0]) // should be 120?
	seamonsters := newSeaMonsters(seaMonsterPattern, lineLength)

	var flattenedImage string

	// make a big ol
	flattenedImage = strings.Join(stitchedImage, "")

	for _, seamonster := range seamonsters {
		// copyimage is going to be the one that gets modified
		//copyImage := flattenedImage
		for idx, char := range flattenedImage {
			if idx%lineLength >= lineLength-seamonster.W+seamonster.First || string(char) != "#" {
				continue
			}
		}
	}
}
