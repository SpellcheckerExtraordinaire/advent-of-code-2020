package day8

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type AoC8Solver uint

var rom [][]string = make([][]string, 0, 0)
var acc int = 0

var alreadyExecuted map[int]bool = make(map[int]bool)

func intFromString(s string) int {
	factor := 0
	if s[0] == byte('+') {
		factor = 1
	} else {
		factor = -1
	}

	slice := s[1:]
	abs, _ := strconv.ParseInt(slice, 10, 0)
	return factor * int(abs)
}

func cycle(ip int) (int, bool) {
	_, exists := alreadyExecuted[ip]
	if exists {
		return ip, false
	}

	alreadyExecuted[ip] = true

	switch rom[ip][0] {
	case "nop":
		{
			return (ip + 1), true
		}
	case "acc":
		{
			acc += intFromString(rom[ip][1])
			return (ip + 1), true
		}
	case "jmp":
		{
			ip += intFromString(rom[ip][1])
			return ip, true
		}
	}
	fmt.Println("You done goofed")
	return 0, false
}

func flush() {
	acc = 0
	alreadyExecuted = make(map[int]bool)
}

func (me AoC8Solver) SolvePartOne(input string) {
	flush()
	rom = make([][]string, 0, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		rom = append(rom, split)
	}

	instructionPointer := 0
	shouldContinue := true
	for {
		instructionPointer, shouldContinue = cycle(instructionPointer)
		if !shouldContinue {
			break
		}
	}

	fmt.Println("Acc: " + strconv.Itoa(acc))
}

func switcheroo(index int, nopJmps []int) {
	if index < 0 {
		return
	}
	adress := nopJmps[index]
	if rom[adress][0] == "jmp" {
		rom[adress][0] = "nop"
	} else {
		rom[adress][0] = "jmp"
	}
}

func (me AoC8Solver) SolvePartTwo(input string) {
	flush()
	rom = make([][]string, 0, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	nopJmps := make([]int, 0, 0)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		if split[0] == "nop" || split[0] == "jmp" {
			nopJmps = append(nopJmps, len(rom))
		}
		rom = append(rom, split)
	}

	instructionPointer := 0
	correctionIndex := 0
	shouldContinue := true
	for {
		instructionPointer, shouldContinue = cycle(instructionPointer)
		if instructionPointer == len(rom) {
			break
		}

		if !shouldContinue {
			switcheroo(correctionIndex-1, nopJmps)
			switcheroo(correctionIndex, nopJmps)
			flush()
			instructionPointer = 0
			correctionIndex++
		}
	}

	fmt.Println("Acc: " + strconv.Itoa(acc))

}

func (me AoC8Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC8Solver(8)
	aoc.SolvePuzzle(solver, sampleOnly)
}
