package day20

import (
	"strings"
)

// seaMonsterPattern is the pattern that's on day20 task 2. The spaces are important.
const seaMonsterPattern = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

type seaMonster struct {
	First, W, H int
	Offsets     []int
}

func newSeaMonsters(s string, lineLength int) []seaMonster {
	seaMonsters := make([]seaMonster, 0, 8)
	monstarz := make([][]string, 0, 8)
	// original and rotated
	originalMonster := strings.Split(s, "\n")

	smR := rotateContent(originalMonster)

	// vflipped and rotated
	smvFlip := vFlipContent(originalMonster)
	smrvFlip := rotateContent(smvFlip)

	// hflipped and rotated
	smhFlip := hFlipContent(originalMonster)
	smrhFlip := rotateContent(smhFlip)

	// vflipped and hflipped, and rotated
	smhvFlip := hFlipContent(smvFlip)
	smrhvFlip := rotateContent(smhvFlip)

	monstarz = append(monstarz, originalMonster, smR, smvFlip, smrvFlip, smhFlip, smrhFlip, smhvFlip, smrhvFlip)

	for _, monstar := range monstarz {
		seaMonsters = append(seaMonsters, newSeaMonster(monstar, lineLength))
	}

	return seaMonsters
}

func newSeaMonster(pattern []string, lineLength int) seaMonster {
	first, offsets := seaMonsterOffsets(pattern, lineLength)
	return seaMonster{
		W:       len(pattern[0]),
		H:       len(pattern),
		First:   first,
		Offsets: offsets,
	}
}

func seaMonsterOffsets(sm []string, lineLength int) (int, []int) {
	offsets := make([]int, 0, 15)

	for row, line := range sm {
		n := row * lineLength
		for _, char := range line {
			if string(char) == "#" {
				offsets = append(offsets, n)
			}
			n++
		}
	}
	// grab the first offset
	shift := offsets[0]
	shiftedOffsets := make([]int, 0, 15)
	for _, offset := range offsets {
		shiftedOffsets = append(shiftedOffsets, offset-shift)
	}
	return shift, shiftedOffsets
}
