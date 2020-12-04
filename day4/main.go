package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	filename = "input.txt"
)

type identity struct {
	ECL, PID, EYR, HCL, BYR, IYR, CID, HGT string
}

func (i identity) IsValid() bool {
	return i.ECL != "" &&
		i.PID != "" &&
		i.EYR != "" &&
		i.HCL != "" &&
		i.BYR != "" &&
		i.IYR != "" &&
		i.HGT != ""
}

//ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
//byr:1937 iyr:2017 cid:147 hgt:183cm

func main() {
	valid := 0
	for _, s := range getInputs() {
		i := parseData(s)
		if i.IsValid() {
			valid++
		}
	}
	fmt.Printf("there are %d valid data", valid)
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n\n")
}

//ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
//byr:1937 iyr:2017 cid:147 hgt:183cm
func parseData(data string) identity {
	slice := strings.Fields(data)

	var i identity
	for _, field := range slice {
		switch field[:2] {
		case "ec":
			i.ECL = field[4:]
			break
		case "pi":
			i.PID = field[4:]
			break
		case "ey":
			i.EYR = field[4:]
			break
		case "hc":
			i.HCL = field[4:]
			break
		case "by":
			i.BYR = field[4:]
			break
		case "iy":
			i.IYR = field[4:]
			break
		case "ci":
			i.CID = field[4:]
			break
		case "hg":
			i.HGT = field[4:]
			break
		}
	}
	return i
}
