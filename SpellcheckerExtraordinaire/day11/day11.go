package day11

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strings"
    "strconv"
)

type AoC11Solver uint

const (
    Floor int = iota+1
    EmptySeat
    OccupiedSeat
)

func generateSeatMap(input string) [][]int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

    seatMap := make([][]int, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
        row := make([]int, 0, 0)
        for _, letter := range line {
            switch letter {
            case '.':
                row = append(row, Floor)
            case 'L':
                row = append(row, EmptySeat)
            case '#':
                row = append(row, OccupiedSeat)
            }
        }
        seatMap = append(seatMap, row)
	}
    return seatMap
}

func getNeighbors(seatMap [][]int, x int, y int) []int {
    neighbors := make([]int, 0, 8)
    for i := y-1; i <= y+1; i++ {
        if i < 0 || i >= len(seatMap) {
            continue
        }
        for j := x-1; j <= x+1; j++ {
            if (j < 0 || j >= len(seatMap[i])) || (j == x && i == y) {
                continue
            }
            neighbors = append(neighbors, seatMap[i][j])
        }
    }
    return neighbors
}

func isValid(seatMap [][]int, x, y int) bool {
    return y >= 0 && y < len(seatMap) && x >= 0 && x < len(seatMap[0])
}

func raycast(seatMap [][]int, x, y, dx, dy int) int {
    x += dx
    y += dy
    for isValid(seatMap, x, y) {
        if seatMap[y][x] == EmptySeat {
            return EmptySeat
        }
        if seatMap[y][x] == OccupiedSeat {
            return OccupiedSeat
        }
        x += dx
        y += dy
    }
    return Floor
}

func getLineOfSight(seatMap [][]int, x int, y int) []int {
    neighbors := []int{
        raycast(seatMap, x, y, 0, 1),
        raycast(seatMap, x, y, 0, -1),
        raycast(seatMap, x, y, 1, 0),
        raycast(seatMap, x, y, -1, 0),
        raycast(seatMap, x, y, 1, 1),
        raycast(seatMap, x, y, 1, -1),
        raycast(seatMap, x, y, -1, 1),
        raycast(seatMap, x, y, -1, -1),
    }
    return neighbors
}

func countOccupied(neighbors []int) int {
    occupiedCount := 0
    for _, seat := range neighbors {
        if seat == OccupiedSeat {
            occupiedCount++
        }
    }
    return occupiedCount
}

func gridCopy(seatMap [][]int) [][]int {
    newMap := make([][]int, 0, len(seatMap))
    for _, row := range seatMap {
        newRow := make([]int, 0, len(row))
        for _, val := range row {
            newRow = append(newRow, val)
        }
        newMap = append(newMap, newRow)
    }
    return newMap
}

func cycle(seatMap [][]int) ([][]int, bool) {
    didChange := false
    newSeatMap := gridCopy(seatMap)
    for y := range seatMap {
        for x := range seatMap[y] {
            seat := seatMap[y][x]
            if seat == Floor {
                continue
            }
            neighbors := getNeighbors(seatMap, x, y)
            occupiedCount := countOccupied(neighbors)
            if seat == EmptySeat && occupiedCount == 0 {
                newSeatMap[y][x] = OccupiedSeat
                didChange = true
            }
            if seat == OccupiedSeat && occupiedCount >= 4 {
                newSeatMap[y][x] = EmptySeat
                didChange = true
            }
        }
    }
    return newSeatMap, didChange
}

func lineOfSightCycle(seatMap [][]int) ([][]int, bool) {
    didChange := false
    newSeatMap := gridCopy(seatMap)
    for y := range seatMap {
        for x := range seatMap[y] {
            seat := seatMap[y][x]
            if seat == Floor {
                continue
            }
            neighbors := getLineOfSight(seatMap, x, y)
            occupiedCount := countOccupied(neighbors)
            if seat == EmptySeat && occupiedCount == 0 {
                newSeatMap[y][x] = OccupiedSeat
                didChange = true
            }
            if seat == OccupiedSeat && occupiedCount >= 5 {
                newSeatMap[y][x] = EmptySeat
                didChange = true
            }
        }
    }
    return newSeatMap, didChange
}

func printMap(seatMap [][]int) {
    for _, row := range seatMap {
        for _, letter := range row {
            switch letter {
            case OccupiedSeat:
                fmt.Print("#")
            case Floor:
                fmt.Print(".")
            case EmptySeat:
                fmt.Print("L")
            }
        }
        fmt.Println()
    }
}

func totalOccupied(seatMap [][]int) int {
    count := 0
    for y := range seatMap {
        for x := range seatMap[y] {
            if seatMap[y][x] == OccupiedSeat {
                count++
            }
        }
    }
    return count
}

func (me AoC11Solver) SolvePartOne(input string) {
    seatMap := generateSeatMap(input)
    didChange := true
    for didChange {
        seatMap, didChange = cycle(seatMap)
    }
    count := totalOccupied(seatMap)
	fmt.Println("Occupied seats: " + strconv.Itoa(count))
}

func (me AoC11Solver) SolvePartTwo(input string) {
    seatMap := generateSeatMap(input)
    didChange := true
    for didChange {
        seatMap, didChange = lineOfSightCycle(seatMap)
    }
    count := totalOccupied(seatMap)
	fmt.Println("Occupied seats: " + strconv.Itoa(count))

}

func (me AoC11Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC11Solver(11)
	aoc.SolvePuzzle(solver, sampleOnly)
}
