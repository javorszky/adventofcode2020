package day20

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func task2() {
	tss := make(tileSets2, 0)
	// there are 144 tiles, so that's a 12x12 grid.
	tiles := getInputs()
	for _, tileString := range tiles {
		tss.addTileSet(newTileSet2(parseTileTask2(tileString)))
	}

	s := int(math.Sqrt(float64(len(tiles))))

	i := image2{
		W:     s,
		H:     s,
		Tiles: make(map[int]map[int]tilev2, 0),
	}

	img, err := fitTileIntoImage2(i, 1, tss)

	if err != nil {
		panic("\nDay 20 task 2 failed")
	}

	// orig and rotated
	oimg := stitchImage(img)
	writeUnflattenedToFile(oimg)
	roimg := rotateContent(oimg)
	writeUnflattenedToFile(roimg)

	// horizontal flip and rotated
	hoimg := hFlipContent(oimg)
	writeUnflattenedToFile(hoimg)
	rhoimg := rotateContent(hoimg)
	writeUnflattenedToFile(rhoimg)

	// vertical flip and rotated
	voimg := vFlipContent(oimg)
	writeUnflattenedToFile(voimg)
	rvoimg := rotateContent(voimg)
	writeUnflattenedToFile(rvoimg)

	// v and h flipped and rotated
	vhoimg := vFlipContent(hFlipContent(oimg))
	writeUnflattenedToFile(vhoimg)
	rvhoimg := rotateContent(vhoimg)
	writeUnflattenedToFile(rvhoimg)

	images := [][]string{
		oimg, roimg, hoimg, rhoimg, voimg, rvoimg, vhoimg, rvhoimg,
	}

	lineLength := len(oimg[0])
	originalMonster := strings.Split(seaMonsterPattern, "\n")

	first, offsets := seaMonsterOffsets(originalMonster, lineLength)
	seamonster := seaMonster{
		W:       len(originalMonster[0]),
		H:       len(originalMonster),
		First:   first,
		Offsets: offsets,
	}
	//stitchedImage := []string{
	//	".#.#..#.##...#.##..#####",
	//	"###....#.#....#..#......",
	//	"##.##.###.#.#..######...",
	//	"###.#####...#.#####.#..#",
	//	"##.#....#.##.####...#.##",
	//	"...########.#....#####.#",
	//	"....#..#...##..#.#.###..",
	//	".####...#..#.....#......",
	//	"#..#.##..#..###.#.##....",
	//	"#.####..#.####.#.#.###..",
	//	"###.#.#...#.######.#..##",
	//	"#.####....##..########.#",
	//	"##..##.#...#...#.#.#.#..",
	//	"...#..#..#.#.##..###.###",
	//	".#.#....#.##.#...###.##.",
	//	"###.#...#..#.##.######..",
	//	".#.#.###.##.##.#..#.##..",
	//	".####.###.#...###.#..#.#",
	//	"..#.#..#..#.#.#.####.###",
	//	"#..####...#.#.#.###.###.",
	//	"#####..#####...###....##",
	//	"#.##..#..#...#..####...#",
	//	".#.###..##..##..####.##.",
	//	"...###...##...#...#..###",
	//}
	var flattenedImage string
	for _, imgToCheck := range images {
		flattenedImage, lineLength := flattenImage(imgToCheck)
		// copyimage is going to be the one that gets modified
		copyImage := flattenedImage
		checkCoords := make([]int, 0)
		// let's loop through the entire original image
		for idx, char := range flattenedImage {
			lastOffset := seamonster.Offsets[len(seamonster.Offsets)-1]
			lenFlattenedImage := len(flattenedImage)
			// skip characters that are either too close to the left, or too close to the right edge, or are not hash.
			// We're looking to find all the hashes that can be starter points for sea monsters.
			if idx%lineLength > lineLength-seamonster.W+seamonster.First || // it's not too close to the right side.
				idx%lineLength < seamonster.First || // not too close to the left side
				idx+lastOffset > lenFlattenedImage || // make sure the last part of the sea monster is still in pic.
				string(char) != "#" { // not a hash character
				continue
			}

			checkCoords = append(checkCoords, idx)

			//for each hash that is in the zone of suitable candidates, let's assume we found the monster.
			foundMonster := true

			// and make sure we save the coordinates for this monster, so we can change them later.
			monsterCoords := make([]int, 0, 15)

			// let's step through the offsets for the seamonster we're trying to find, and check whether all the other
			// characters are also #. If not, break, and move on to the next character.
			for _, offset := range seamonster.Offsets {
				//fmt.Printf("checking for character %d at offset %d with total offset %d\n", sidx, offset, offset+idx)
				// the next char at specified offset from first is not a hash, break off.
				if flattenedImage[idx+offset-1:idx+offset] != "#" {
					foundMonster = false
					break
				}
				monsterCoords = append(monsterCoords, idx+offset)
			}
			if foundMonster {
				copyImage = replaceMonster(copyImage, monsterCoords)
			}
		}
		copyImage = replaceMonster(copyImage, checkCoords)

		// save the changed and found seamonsters into the original image and let's move on to the next seamonster.
		flattenedImage = copyImage
		writeUnflattenedToFile(unflattenImage(flattenedImage, lineLength))
	}

	fmt.Printf("Day 20 task 2: the roughness of the seas is %d\n", strings.Count(flattenedImage, "#"))
}

func replaceMonster(image string, coords []int) string {
	for _, coord := range coords {
		image = image[:coord] + "O" + image[coord+1:]
	}
	return image
}

func flattenImage(img []string) (string, int) {
	lineLength := len(img[0]) // should be 120?
	flattenedImage := strings.Join(img, "")

	return flattenedImage, lineLength
}

func unflattenImage(img string, linelength int) []string {
	s := make([]string, 0)
	l := len(img)
	previous := 0
	for i := linelength; i <= l; i = i + linelength {
		s = append(s, img[previous:i])
		previous = i
	}
	return s
}

func writeUnflattenedToFile(img []string) {
	f, err := os.OpenFile("day20/images.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("opening file failed: %s", err))
	}

	defer f.Close()
	//
	//if _, err := f.WriteString(strings.Join(sm, "\n")); err != nil {
	//	panic(fmt.Sprintf("writing string to file failed: %s", err))
	//}

	if _, err := f.WriteString("\n\n" + strings.Join(img, "\n")); err != nil {
		panic(fmt.Sprintf("writing string to file failed: %s", err))
	}
}
