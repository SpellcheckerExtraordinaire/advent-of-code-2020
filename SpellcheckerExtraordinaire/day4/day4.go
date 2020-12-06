package day4

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type AoC4Solver uint

func generatePassportMap(passport []string) map[string]string {
	regex := regexp.MustCompile("([a-z]{3}):(\\S+)")
	passportMap := make(map[string]string)
	for _, line := range passport {
		groups := regex.FindAllStringSubmatch(line, -1)
		for _, group := range groups {
			passportMap[group[1]] = group[2]
		}
	}
	return passportMap
}

func hasRequiredFields(passportMap map[string]string) bool {
	if len(passportMap) == 8 {
		return true
	} else if len(passportMap) < 7 {
		return false
	}

	_, ok := passportMap["cid"]
	return !ok
}

func isWithinBounds(s string, min int, max int) bool {
	number, _ := strconv.ParseInt(s, 10, 0)
	return int(number) >= min && int(number) <= max
}

func isValidPassport(passportMap map[string]string) bool {
	if !isWithinBounds(passportMap["byr"], 1920, 2002) {
		return false
	}
	if !isWithinBounds(passportMap["iyr"], 2010, 2020) {
		return false
	}
	if !isWithinBounds(passportMap["eyr"], 2020, 2030) {
		return false
	}

	heightRegex := regexp.MustCompile("(\\d+)(cm|in)")
	captured := heightRegex.FindStringSubmatch(passportMap["hgt"])
	if len(captured) != 3 {
		return false
	}
	validHeight := false
	switch {
	case captured[2] == "cm":
		validHeight = isWithinBounds(captured[1], 150, 193)
	case captured[2] == "in":
		validHeight = isWithinBounds(captured[1], 59, 76)
	}
	if !validHeight {
		return false
	}

	if len(passportMap["hcl"]) != 7 {
		return false
	}
	hairColorRegex := regexp.MustCompile("#(\\d|[a-f]){6}")
	captured = hairColorRegex.FindStringSubmatch(passportMap["hcl"])
	if len(captured) != 2 {
		return false
	}

	eyeColorRegex := regexp.MustCompile("(amb|blu|brn|gry|grn|hzl|oth)")
	captured = eyeColorRegex.FindStringSubmatch(passportMap["ecl"])
	if len(captured) != 2 {
		return false
	}

	if len(passportMap["pid"]) != 9 {
		return false
	}
	passportIDRegex := regexp.MustCompile("\\d{9}")
	captured = passportIDRegex.FindStringSubmatch(passportMap["pid"])
	if len(captured) != 1 {
		return false
	}

	return true
}

func (me AoC4Solver) SolvePartOne(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	passportCount := 0
	passport := make([]string, 0, 8)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 1 {
			passport = append(passport, line)
		} else {
			if hasRequiredFields(generatePassportMap(passport)) {
				passportCount += 1
			}
			passport = make([]string, 0, 8)
		}
	}
	fmt.Println("Valid Passports: " + strconv.Itoa(passportCount))
}

func (me AoC4Solver) SolvePartTwo(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	passportCount := 0
	passport := make([]string, 0, 8)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 1 {
			passport = append(passport, line)
		} else {
			passportMap := generatePassportMap(passport)
			if hasRequiredFields(passportMap) && isValidPassport(passportMap) {
				passportCount += 1
			}
			passport = make([]string, 0, 8)
		}
	}
	fmt.Println("Valid Passports: " + strconv.Itoa(passportCount))
}

func (me AoC4Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC4Solver(4)
	aoc.SolvePuzzle(solver, sampleOnly)
}
