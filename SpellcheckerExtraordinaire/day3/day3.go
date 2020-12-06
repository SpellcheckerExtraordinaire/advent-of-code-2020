package day3

import (
	"aoc-2020-go/aoc"
)

type AoC3Solver uint

func (me AoC3Solver) SolvePartOne(input string) {

}

func (me AoC3Solver) SolvePartTwo(input string) {

}

func (me AoC3Solver) Day() uint {
	return uint(me)
}

func Solve() {
	solver := AoC3Solver(3)
	aoc.SolvePuzzle(solver)
}
