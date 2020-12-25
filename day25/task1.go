package day25

import (
	"errors"
	"fmt"
)

const transformMod int = 20201227

func task1() {
	card, door := getInputs()
	fmt.Printf("\nDay 25 task 1: card public key is %d, door pkey is %d\n", card, door)

	loopCard, err := findLoopSize(1, 7, card)
	if err != nil {
		panic(fmt.Sprintf("could not find card loop: %s", err))
	}
	loopDoor, err := findLoopSize(1, 7, door)
	if err != nil {
		panic(fmt.Sprintf("could not find door loop: %s", err))
	}
	fmt.Printf(" - Loop size for card with pkey %d is %d\n", card, loopCard)
	fmt.Printf(" - Loop size for door with pkey %d is %d\n", door, loopDoor)

	encryptionKey := transformNTimes(1, card, loopDoor)
	encryptionKey2 := transformNTimes(1, door, loopCard)
	if encryptionKey != encryptionKey2 {
		panic(fmt.Sprintf("the two encryption keys do not match.\nCard pkey with loop door: %d\nDoor pkey with loop card: %d\n", encryptionKey, encryptionKey2))
	}
	fmt.Printf("The encryption key they're trying to establish is %d\n", encryptionKey)
}

func transformNTimes(what, subject, loopSize int) int {
	for i := 0; i < loopSize; i++ {
		what = transformOnce(what, subject)
	}

	return what
}

func transformOnce(what, subject int) int {
	what = what * subject
	what = what % transformMod
	return what
}

func findLoopSize(what, subject, target int) (int, error) {
	for i := 1; i < 10000000; i++ {
		what = transformOnce(what, subject)
		if what == target {
			return i, nil
		}
	}

	return 0, errors.New("could not find loop in 10,000 tries")
}
