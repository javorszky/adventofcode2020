package day16

import (
	"io/ioutil"
	"regexp"
	"strings"
)

const filename = "day16/input.txt"

var reRule = regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)

func Tasks() {
	task1()
	//task2()
}

func task2() {
	_, _, _ = getInputs()
}

// getInputs reads the input.txt file and returns three sets of data: rules, your ticket, nearby tickets.
func getInputs() ([]string, string, []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// separate the data into chunks
	split1 := strings.Split(strings.TrimRight(string(data), "\n"), "\nyour ticket:\n")
	split2 := strings.Split(strings.Trim(split1[1], "\n"), "\nnearby tickets:\n")

	// format the chunks
	rules := strings.Split(strings.Trim(split1[0], "\n"), "\n")
	yourTicket := strings.Trim(split2[0], "\n")
	nearbyTickets := strings.Split(strings.Trim(split2[1], "\n"), "\n")

	// return the chunks
	return rules, yourTicket, nearbyTickets
}
