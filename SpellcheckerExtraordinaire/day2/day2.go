package day2

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"regexp"

	"strconv"
	"strings"
)

type AoC2Solver uint

type password struct {
	first  int
	second int
	letter string
	text   string
}

func lineToPassword(line string) password {
	regex := regexp.MustCompile("(\\d*)-(\\d*) ([a-z]): ([a-z]*)")
	groups := regex.FindStringSubmatch(line)

	first, _ := strconv.ParseInt(groups[1], 10, 0)
	second, _ := strconv.ParseInt(groups[2], 10, 0)
	letter := groups[3]
	text := groups[4]

	return password{int(first), int(second), letter, text}
}

func checkValidityOldJob(pwd password) bool {
	count := strings.Count(pwd.text, pwd.letter)
	return count >= pwd.first && count <= pwd.second
}

func checkValidityNewJob(pwd password) bool {
	first := pwd.text[pwd.first-1] == []byte(pwd.letter)[0]
	second := pwd.text[pwd.second-1] == []byte(pwd.letter)[0]
	return first != second
}

func countPolicyConformation(input string, policy func(password) bool) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	correctPasswords := 0
	for scanner.Scan() {
		if policy(lineToPassword(scanner.Text())) {
			correctPasswords++
		}
	}

	return correctPasswords
}

func (me AoC2Solver) SolvePartOne(input string) {
	fmt.Println("Valid Passwords: ", countPolicyConformation(input, checkValidityOldJob))
}

func (me AoC2Solver) SolvePartTwo(input string) {
	fmt.Println("Valid Passwords: ", countPolicyConformation(input, checkValidityNewJob))
}

func (me AoC2Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC2Solver(2)
	aoc.SolvePuzzle(solver, sampleOnly)
}
