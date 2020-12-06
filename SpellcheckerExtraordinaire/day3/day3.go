package day3

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type AoC3Solver uint

func generateTreeMap(input string) [][]int {
	trees := make([][]int, 0, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, make([]int, len(line), len(line)))

		for x, tree := range line {
			weight := 0
			if tree == '#' {
				weight = 1
			}
			trees[y][x] = weight
		}

		y++
	}
	return trees
}

func hitTrees(trees [][]int, slopeX int, slopeY int) int {
	treesHit := 0
	x := 0
	for y := slopeY; y < len(trees); y += slopeY {
		x = (x + slopeX) % len(trees[0])
		treesHit += trees[y][x]
	}
	return treesHit
}

func (me AoC3Solver) SolvePartOne(input string) {
	trees := generateTreeMap(input)
	fmt.Println("Result: " + strconv.Itoa(hitTrees(trees, 3, 1)))
}

func (me AoC3Solver) SolvePartTwo(input string) {
	trees := generateTreeMap(input)
	product := 1
	product *= hitTrees(trees, 1, 1)
	product *= hitTrees(trees, 3, 1)
	product *= hitTrees(trees, 5, 1)
	product *= hitTrees(trees, 7, 1)
	product *= hitTrees(trees, 1, 2)

	fmt.Println("Result: " + strconv.Itoa(product))
}

func (me AoC3Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC3Solver(3)
	aoc.SolvePuzzle(solver, sampleOnly)
}
