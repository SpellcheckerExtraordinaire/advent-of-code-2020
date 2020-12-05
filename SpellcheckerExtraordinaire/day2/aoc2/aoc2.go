package aoc2

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"regexp"

	"strconv"
	"strings"
)

type AoCXSolver uint

type password struct {
	min    int
	max    int
	letter string
	text   string
}

func lineToPassword(line string) password {
	regex := regexp.MustCompile("(\\d*)-(\\d*) ([a-z]): ([a-z]*)")
	groups := regex.FindStringSubmatch(line)

	min, _ := strconv.ParseInt(groups[1], 10, 0)
	max, _ := strconv.ParseInt(groups[2], 10, 0)
	letter := groups[3]
	text := groups[4]

	return password{int(min), int(max), letter, text}
}

func checkValidity(pwd password) bool {
	count := strings.Count(pwd.text, pwd.letter)
	return count <= pwd.max && count >= pwd.min
}

func (me AoCXSolver) SolvePartOne(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	correctPasswords := 0
	for scanner.Scan() {
		if checkValidity(lineToPassword(scanner.Text())) {
			correctPasswords++
		}
	}

	fmt.Println("Valid Passwords: ", correctPasswords)

}

func (me AoCXSolver) SolvePartTwo(input string) {

}

func (me AoCXSolver) Day() uint {
	return uint(me)
}

func Solve() {
	solver := AoCXSolver(0)
	aoc.SolvePuzzle(solver)
}
