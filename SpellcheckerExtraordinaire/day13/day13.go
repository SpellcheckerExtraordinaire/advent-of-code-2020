package day13

import (
	"aoc-2020-go/aoc"
	"fmt"
	"strconv"
	"strings"
)

type AoC13Solver uint

// returnd gcD, s and t
func extendedEuclidianAlgorithm(a, b int64) (int64, int64, int64) {
    if (b == 0) {
        return a, 1, 0
    }
    remainder := a % b
    gcd, s, t := extendedEuclidianAlgorithm(b, remainder)
    return gcd, t, s - (a/b) * t
}

func (me AoC13Solver) SolvePartOne(input string) {
    lines := strings.Fields(input)
    timestamp, _ := strconv.ParseInt(lines[0], 10, 0)
    ids := strings.Split(lines[1], ",")

    soonest := int64(99999)
    bus := int64(-1)
    for _, id := range ids {
        if id == "x" {
            continue
        }
        interval, _ := strconv.ParseInt(id, 10, 0)
        remainder := timestamp % interval
        waitTime := interval - remainder
        if waitTime < soonest {
            bus = interval
            soonest = waitTime
        }
    }
    result := bus * soonest
	fmt.Println("Result: " + strconv.Itoa(int(result)))
}

func (me AoC13Solver) SolvePartTwo(input string) {
    lines := strings.Fields(input)
    ids := strings.Split(lines[1], ",")
    congruencySystem := make([][]int, 0, len(ids))
    for i, id := range ids{
        if id == "x" {
            continue
        }
        _interval, _ := strconv.ParseInt(id, 10, 0)
        interval := int(_interval)
        congruencySystem = append(congruencySystem, []int{ (-i) % interval, interval})
    }

    bigM := int64(1)
    for _, congruency := range congruencySystem {
        bigM *= int64(congruency[1])
    }

    x := int64(0)
    for _, congruency := range congruencySystem {
        ai := int64(congruency[0])
        mi := int64(congruency[1])
        bigMi := bigM / mi
        // gcD and t denoted as _ to supress compiler errors for unused variables
        _, s, _ := extendedEuclidianAlgorithm(bigMi, mi)
        //fmt.Printf("mi * t + s * Mi: %v * %v + %v * %v\n", mi, t, s, bigMi)
        x += ai * s * bigMi
    }

    for x < 0 {
        x += bigM
    }

	fmt.Println("Result: " + strconv.Itoa(int(x)))
}

func (me AoC13Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC13Solver(13)
	aoc.SolvePuzzle(solver, sampleOnly)
}
