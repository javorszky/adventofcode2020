package day4

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	filename = "day4/input.txt"
)

var (
	reHCL = regexp.MustCompile("^#[0-9a-f]{6}$")
	rePID = regexp.MustCompile(`^\d{9}$`)
)

type identity struct {
	ECL, PID, EYR, HCL, BYR, IYR, CID, HGT string
}

func (i identity) IsValidv2() bool {
	return i.eclValid() &&
		i.pidValid() &&
		i.eyrValid() &&
		i.hclValid() &&
		i.byrValid() &&
		i.iyrValid() &&
		i.hgtValid()
}

func (i identity) IsValidv1() bool {
	return i.ECL != "" &&
		i.PID != "" &&
		i.EYR != "" &&
		i.HCL != "" &&
		i.BYR != "" &&
		i.IYR != "" &&
		i.HGT != ""
}

// eclValid return whether eye colour is valid: exactly one of: amb blu brn gry grn hzl oth.
func (i identity) eclValid() bool {
	switch i.ECL {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

// pidValid returns whether pid is a valid 9 digit string including leading zeroes.
func (i identity) pidValid() bool {
	return rePID.MatchString(i.PID)
}

// eyrValid will return whether expiration year is 4 digits, at least 2020, at most 2030.
func (i identity) eyrValid() bool {
	b, err := strconv.Atoi(i.EYR)
	if err != nil {
		return false
	}

	return 2020 <= b && b <= 2030
}

// hclValid returns whether hair colour is hash followed by a hex colour code: 6 chars of 0-9a-f.
func (i identity) hclValid() bool {
	return reHCL.MatchString(i.HCL)
}

// byrValid will return if birth year is 4 digits, at least 1920, at most 2002.
func (i identity) byrValid() bool {
	b, err := strconv.Atoi(i.BYR)
	if err != nil {
		return false
	}
	return 1920 <= b && b <= 2002
}

// iyrValid will return whether issue year is 4 digits, at least 2010, at most 2020.
func (i identity) iyrValid() bool {
	b, err := strconv.Atoi(i.IYR)
	if err != nil {
		return false
	}
	return 2010 <= b && b <= 2020
}

// hgtValid will return whether height ends in either in or cm, and if cm 150 - 193, in 59 - 76.
func (i identity) hgtValid() bool {
	if len(i.HGT) < 3 {
		return false
	}
	switch i.HGT[len(i.HGT)-2:] {
	case "cm":
		n, err := strconv.Atoi(i.HGT[:len(i.HGT)-2])
		if err != nil {
			return false
		}
		return 150 <= n && n <= 193
	case "in":
		n, err := strconv.Atoi(i.HGT[:len(i.HGT)-2])
		if err != nil {
			return false
		}
		return 59 <= n && n <= 76
	}
	return false
}

//ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
//iyr:1937 iyr:2017 cid:147 hgt:183cm

func Tasks() {

	task1()
	task2()

}

func task1() {
	valid := 0
	for _, s := range getInputs() {
		i := parseData(s)
		if i.IsValidv1() {
			valid++
		}
	}
	fmt.Printf("\nDay 4 task 1: there are %d valid data\n", valid)
}

func task2() {
	valid := 0
	for _, s := range getInputs() {
		i := parseData(s)
		if i.IsValidv2() {
			valid++
		}
	}
	fmt.Printf("Day 4 task 2: there are %d valid data\n", valid)
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
//iyr:1937 iyr:2017 cid:147 hgt:183cm
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
