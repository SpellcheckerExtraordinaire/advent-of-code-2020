package day5

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type AoC5Solver uint

func seatId(boardingPass string) int {
	var row byte = 0
	for i := 0; i < 7; i++ {
		var bit byte
		switch {
		case boardingPass[i] == byte('B'):
			bit = 1
		case boardingPass[i] == byte('F'):
			bit = 0
		}
		row = row | (bit << (6 - i))
	}

	var column byte = 0
	for i := 0; i < 3; i++ {
		var bit byte
		switch {
		case boardingPass[7+i] == byte('R'):
			bit = 1
		case boardingPass[7+i] == byte('L'):
			bit = 0
		}
		column = column | (bit << (2 - i))
	}
	return int(row)*8 + int(column)
}

func (me AoC5Solver) SolvePartOne(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	highestSeat := 0
	for scanner.Scan() {
		line := scanner.Text()
		newSeat := seatId(line)
		if newSeat > highestSeat {
			highestSeat = newSeat
		}
	}
	fmt.Println("Highest Seat: " + strconv.Itoa(highestSeat))
}

func (me AoC5Solver) SolvePartTwo(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	var seats [128 * 8]bool

	for scanner.Scan() {
		line := scanner.Text()
		newSeat := seatId(line)
		seats[newSeat] = true
	}

	var mySeat int
	for seatId, occupied := range seats {
		if !occupied && (seatId > 0 && seats[seatId-1]) && seats[seatId+1] {
			mySeat = seatId
			break
		}
	}

	fmt.Println("SeatID: " + strconv.Itoa(mySeat))
}

func (me AoC5Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC5Solver(5)
	aoc.SolvePuzzle(solver, sampleOnly)
}
