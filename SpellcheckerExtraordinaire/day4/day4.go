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

func isValidPassport(passport []string) bool {
	regex := regexp.MustCompile("([a-z]{3}):(\\S+)\\s")
	passportMap := make(map[string]string)
	for _, line := range passport {
		groups := regex.FindAllStringSubmatch(line, -1)
		fmt.Println(groups)
	}
	return false
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
		} else if isValidPassport(passport) {
			passportCount += 1
		}
	}
	fmt.Println("Result: " + strconv.Itoa(passportCount))
}

func (me AoC4Solver) SolvePartTwo(input string) {

}

func (me AoC4Solver) Day() uint {
	return uint(me)
}

func Solve() {
	solver := AoC4Solver(4)
	aoc.SolvePuzzle(solver)
}
