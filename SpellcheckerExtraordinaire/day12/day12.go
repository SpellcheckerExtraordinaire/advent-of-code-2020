package day12

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strings"
    "strconv"
    "math"
)

type AoC12Solver uint

const (
    NORTH = iota
    EAST
    SOUTH
    WEST
)

type Direction int

type Vector struct {
    x int
    y int
}

type Matrix struct {
    ij00 int
    ij01 int
    ij10 int
    ij11 int
}

type Ship struct {
    pos Vector
    waypoint Vector
    dir Direction
}

var cwRotation = Matrix {
    0, -1,
    1, 0,
}

var ccwRotation = Matrix {
    0, 1,
    -1, 0,
}

func (a *Vector) add(b Vector) {
    a.x += b.x
    a.y += b.y
}

func (v *Vector) mul(m *Matrix) {
    newX := v.x * m.ij00 + v.y * m.ij01
    newY := v.x * m.ij10 + v.y * m.ij11
    v.x = newX
    v.y = newY
}

func (v *Vector) abs() int {
    return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)))
}

func (v *Vector) scale(scale int) Vector {
    newVec := Vector {v.x * scale, v.y * scale}
    return newVec
}

func inDir(dir Direction, amount int) Vector {
    x, y := 0, 0
    switch dir {
    case NORTH:
        y = -amount
    case EAST:
        x = amount
    case SOUTH:
        y = amount
    case WEST:
        x = -amount
    }
    return Vector{x, y}
}

func (me AoC12Solver) SolvePartOne(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

    ship := Ship{Vector{0, 0}, Vector{0, 0}, EAST}

	for scanner.Scan() {
		line := scanner.Text()
        command := line[0]
        amountRaw, _ := strconv.ParseInt(line[1:], 10, 0)
        amount := int(amountRaw)
        var delta Vector
        switch command {
        case 'F':
            delta = inDir(ship.dir, amount)
        case 'N':
            delta = inDir(NORTH, amount)
        case 'E':
            delta = inDir(EAST, amount)
        case 'S':
            delta = inDir(SOUTH, amount)
        case 'W':
            delta = inDir(WEST, amount)
        case 'R':
            steps := amount / 90
            ship.dir = Direction((int(ship.dir) + steps) % 4)
        case 'L':
            steps := amount / 90
            ship.dir = Direction((int(ship.dir) + 4 - steps) % 4)
        }
        ship.pos.add(delta)
	}
	fmt.Println("Result: " + strconv.Itoa(ship.pos.abs()))

}

func (me AoC12Solver) SolvePartTwo(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

    ship := Ship{Vector{0, 0}, Vector{10, -1}, EAST}

	for scanner.Scan() {
		line := scanner.Text()
        command := line[0]
        amountRaw, _ := strconv.ParseInt(line[1:], 10, 0)
        amount := int(amountRaw)
        switch command {
        case 'F':
            ship.pos.add(ship.waypoint.scale(amount))
        case 'N':
            ship.waypoint.add(inDir(NORTH, amount))
        case 'E':
            ship.waypoint.add(inDir(EAST, amount))
        case 'S':
            ship.waypoint.add(inDir(SOUTH, amount))
        case 'W':
            ship.waypoint.add(inDir(WEST, amount))
        case 'R':
            steps := amount / 90
            for i := 0; i < steps; i++ {
                ship.waypoint.mul(&cwRotation)
            }
        case 'L':
            steps := amount / 90
            for i := 0; i < steps; i++ {
                ship.waypoint.mul(&ccwRotation)
            }
        }
	}
	fmt.Println("Result: " + strconv.Itoa(ship.pos.abs()))
}

func (me AoC12Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC12Solver(12)
	aoc.SolvePuzzle(solver, sampleOnly)
}
