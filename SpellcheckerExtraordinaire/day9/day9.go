package day9

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type AoC9Solver uint

type Buffer struct {
	content []int
	cursor  int
}

func (b *Buffer) write(num int) {
	b.content[b.cursor] = num
	b.cursor = (b.cursor + 1) % len(b.content)
}

func (b *Buffer) isValidNext(next int) bool {
	for i, num1 := range b.content {
		for j := i + 1; j < len(b.content); j++ {
			num2 := b.content[j]

			if (num1 + num2) == next {
				return true
			}
		}
	}
	return false
}

func (b *Buffer) findInvalid(preambleLength int, xmas []int) int {
	i := preambleLength
	for {
		if !b.isValidNext(xmas[i]) {
			break
		}
		b.write(xmas[i])
		i++
	}
	return xmas[i]
}

func populate(input string) (*Buffer, int, []int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	xmas := make([]int, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.ParseInt(line, 10, 0)
		xmas = append(xmas, int(number))
	}

	// HACK
	var preambleLength int
	if len(xmas) > 50 {
		preambleLength = 25
	} else {
		preambleLength = 5
	}
	buffer := Buffer{make([]int, preambleLength, preambleLength), 0}
	for i := 0; i < preambleLength; i++ {
		buffer.write(xmas[i])
	}

	return &buffer, preambleLength, xmas
}

func findRange(invalid int, xmas []int) (int, int) {
	for start := 0; start < len(xmas); start++ {
		sum := xmas[start]
		end := start + 1
		for sum < invalid {
			sum += xmas[end]
			end++
		}
		if sum == invalid {
			return start, end - 1
		}
	}
	fmt.Println(":(")
	return 0, 0
}

func findCandidatesInRange(start int, end int, xmas []int) (int, int) {
	min := 9999999
	max := 0
	for i := start; i <= end; i++ {
		num := xmas[i]
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func (me AoC9Solver) SolvePartOne(input string) {
	buffer, preambleLength, xmas := populate(input)
	invalid := buffer.findInvalid(preambleLength, xmas)
	fmt.Println("First incorrect: " + strconv.Itoa(invalid))
}

func (me AoC9Solver) SolvePartTwo(input string) {
	buffer, preambleLength, xmas := populate(input)
	invalid := buffer.findInvalid(preambleLength, xmas)

	start, end := findRange(invalid, xmas)
	min, max := findCandidatesInRange(start, end, xmas)

	fmt.Println("Sum of Range is " + strconv.Itoa(min+max))

}

func (me AoC9Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC9Solver(9)
	aoc.SolvePuzzle(solver, sampleOnly)
}
