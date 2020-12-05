package aoc1

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type AoC1Solver uint

var numberCache []int = nil

func findPair(numbers []int) (int, int) {
	for i, num1 := range numbers {
		for j := len(numbers) - 1; j > i; j-- {
			num2 := numbers[j]
			if num1+num2 == 2020 {
				return num1, num2
			}
		}
	}
	fmt.Println("Whoopsie Daisie")
	return 0, 0
}

func findTriple(numbers []int) (int, int, int) {
	for i, num1 := range numbers {
		for j := len(numbers) - 1; j > i; j-- {
			num2 := numbers[j]
			sum := num1 + num2
			if sum >= 2020 {
				continue
			}
			for k := i; k < j; k++ {
				num3 := numbers[k]
				if sum+num3 == 2020 {
					return num1, num2, num3
				}
			}
		}
	}
	fmt.Println("No luck here")
	return 0, 0, 0
}

func convertInput(input string) []int {
	if numberCache != nil {
		return numberCache
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	numbers := make([]int, 0, 200)

	for scanner.Scan() {
		nextNumber, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		numbers = append(numbers, int(nextNumber))
	}

	return numbers
}

func (me AoC1Solver) SolvePartOne(input string) {
	numbers := convertInput(input)
	first, second := findPair(numbers)
	fmt.Printf("Result: %v\n", first*second)
}

func (me AoC1Solver) SolvePartTwo(input string) {
	numbers := convertInput(input)
	first, second, third := findTriple(numbers)
	fmt.Printf("Result: %v\n", first*second*third)
}

func (me AoC1Solver) Day() uint {
	return uint(me)
}

func Solve() {
	solver := AoC1Solver(1)
	aoc.SolvePuzzle(solver)
}
