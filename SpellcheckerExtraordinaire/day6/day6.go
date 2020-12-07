package day6

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type AoC6Solver uint

func countIndividualAnswers(answers []string) int {
	answerMap := make(map[rune]bool)
	for _, str := range answers {
		for _, letter := range str {
			answerMap[letter] = true
		}
	}
	return len(answerMap)
}

func (me AoC6Solver) SolvePartOne(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	count := 0
	answers := make([]string, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			answers = append(answers, line)
		} else {
			count += countIndividualAnswers(answers)
			answers = make([]string, 0, 0)
		}
	}
	// handle last line, bit hacky
	if len(answers) != 0 {
		count += countIndividualAnswers(answers)
	}
	fmt.Println("Count: " + strconv.Itoa(count))
}

func (me AoC6Solver) SolvePartTwo(input string) {

}

func (me AoC6Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC6Solver(6)
	aoc.SolvePuzzle(solver, sampleOnly)
}
