package aoc

import (
	"fmt"
	"io/ioutil"
)

type AoCSolver interface {
	SolvePartOne(string)
	SolvePartTwo(string)
	Day() uint
}

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Could not read file", err)
		return ""
	}
	return string(data)
}

func SolvePuzzle(solver AoCSolver) {
	fmt.Printf("\n== Solution for Day %v ==\n", solver.Day())

	sample := readFile("sample.txt")
	puzzle := readFile("puzzle.txt")

	fmt.Println("-- Part One --")
	fmt.Println("Attempt to solve sample:")
	solver.SolvePartOne(sample)
	fmt.Println("Attempt to solve full input:")
	solver.SolvePartOne(puzzle)

	fmt.Println("-- Part Two --")
	fmt.Println("Attempt to solve sample:")
	solver.SolvePartTwo(sample)
	fmt.Println("Attempt to solve full input:")
	solver.SolvePartTwo(puzzle)

	fmt.Printf("\n\n")
}
