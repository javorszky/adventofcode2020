package day20

import (
	"fmt"
	"math"
	"strings"
)

func task2() {
	tss := make(tileSets2, 0)
	// there are 144 tiles, so that's a 12x12 grid.
	tiles := getInputs(filename)
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
	//oimg := stitchImage(img2)
	oimg := stitchImage(img)

	roughness := seaRoughness(oimg)

	fmt.Printf("Day 20 task 2: the roughness of the seas is %d\n", roughness)
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

func seaRoughness(stitchedImage []string) int {
	flattenedImage, lineLength := flattenImage(stitchedImage)
	seamonsters := newSeaMonsters(seaMonsterPattern, lineLength)
	for _, seamonster := range seamonsters {
		replaceMonsters := make([][]int, 0)
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

			//checkCoords = append(checkCoords, idx)

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
				replaceMonsters = append(replaceMonsters, monsterCoords)
			}
		}

		fmt.Printf("number of monsters found: %d\n", len(replaceMonsters))

		for _, monsterCoords := range replaceMonsters {
			flattenedImage = replaceMonster(flattenedImage, monsterCoords)
		}
	}

	return strings.Count(flattenedImage, "#")
}

func seaRoughnessRotatedMap(stitchedImage []string) int {
	images := [][]string{
		stitchedImage,
		rotateContent(stitchedImage),
		rotateContent(rotateContent(stitchedImage)),
		rotateContent(rotateContent(rotateContent(stitchedImage))),
		hFlipContent(stitchedImage),
		rotateContent(hFlipContent(stitchedImage)),
		rotateContent(rotateContent(hFlipContent(stitchedImage))),
		rotateContent(rotateContent(rotateContent(hFlipContent(stitchedImage)))),
	}

	monstar := strings.Split(seaMonsterPattern, "\n")
	first, offsets := seaMonsterOffsets(monstar, len(stitchedImage[0]))
	seaMonster := seaMonster{
		W:       len(monstar[0]),
		H:       len(monstar),
		First:   first,
		Offsets: offsets,
	}

	for _, image := range images {
		flattenedImage, lineLength := flattenImage(image)

		replaceMonsters := make([][]int, 0)
		for idx, char := range flattenedImage {
			lastOffset := seaMonster.Offsets[len(seaMonster.Offsets)-1]
			lenFlattenedImage := len(flattenedImage)
			// skip characters that are either too close to the left, or too close to the right edge, or are not hash.
			// We're looking to find all the hashes that can be starter points for sea monsters.
			if idx%lineLength > lineLength-seaMonster.W+seaMonster.First || // it's not too close to the right side.
				idx%lineLength < seaMonster.First || // not too close to the left side
				idx+lastOffset > lenFlattenedImage || // make sure the last part of the sea monster is still in pic.
				string(char) != "#" { // not a hash character
				continue
			}

			//checkCoords = append(checkCoords, idx)

			//for each hash that is in the zone of suitable candidates, let's assume we found the monster.
			foundMonster := true

			// and make sure we save the coordinates for this monster, so we can change them later.
			monsterCoords := make([]int, 0, 15)

			// let's step through the offsets for the seamonster we're trying to find, and check whether all the other
			// characters are also #. If not, break, and move on to the next character.
			for _, offset := range seaMonster.Offsets {
				//fmt.Printf("checking for character %d at offset %d with total offset %d\n", sidx, offset, offset+idx)
				// the next char at specified offset from first is not a hash, break off.
				if flattenedImage[idx+offset-1:idx+offset] != "#" {
					foundMonster = false
					break
				}
				monsterCoords = append(monsterCoords, idx+offset)
			}
			if foundMonster {
				replaceMonsters = append(replaceMonsters, monsterCoords)
			}
		}

		fmt.Printf("number of monsters found: %d\n", len(replaceMonsters))

		for _, monsterCoords := range replaceMonsters {
			flattenedImage = replaceMonster(flattenedImage, monsterCoords)
		}
	}

	return 0
}

func findMonsterCoords(flattenedImage string, lineLength int, seaMonster seaMonster) [][]int {
	fmt.Printf("%#v\n", seaMonster.Offsets)
	replaceMonsters := make([][]int, 0)
	for idx, char := range flattenedImage {
		lastOffset := seaMonster.Offsets[len(seaMonster.Offsets)-1]
		lenFlattenedImage := len(flattenedImage)
		// skip characters that are either too close to the left, or too close to the right edge, or are not hash.
		// We're looking to find all the hashes that can be starter points for sea monsters.
		if idx%lineLength > lineLength-seaMonster.W+seaMonster.First || // it's not too close to the right side.
			idx%lineLength < seaMonster.First || // not too close to the left side
			idx+lastOffset > lenFlattenedImage || // make sure the last part of the sea monster is still in pic.
			string(char) != "#" { // not a hash character
			continue
		}
		fmt.Printf("found hash char at idx %d: '%s'\n", idx, flattenedImage[idx:idx+1])

		//checkCoords = append(checkCoords, idx)

		//for each hash that is in the zone of suitable candidates, let's assume we found the monster.
		foundMonster := true

		// and make sure we save the coordinates for this monster, so we can change them later.
		monsterCoords := make([]int, 0, 15)

		// let's step through the offsets for the seamonster we're trying to find, and check whether all the other
		// characters are also #. If not, break, and move on to the next character.
		for _, offset := range seaMonster.Offsets {
			//fmt.Printf("checking for character %d at offset %d with total offset %d\n", sidx, offset, offset+idx)
			// the next char at specified offset from first is not a hash, break off.
			shiftedMonsterOffset := idx + offset
			fmt.Printf("checking char at offset %d, which is '%s'\n", shiftedMonsterOffset, flattenedImage[shiftedMonsterOffset:shiftedMonsterOffset+1])
			if flattenedImage[shiftedMonsterOffset:shiftedMonsterOffset+1] != "#" {
				foundMonster = false
				fmt.Printf("whoops, did not find a monster part, aborting...\n")
				break
			}
			monsterCoords = append(monsterCoords, idx+offset)
		}
		if foundMonster {
			replaceMonsters = append(replaceMonsters, monsterCoords)
		}
	}

	return replaceMonsters
}
