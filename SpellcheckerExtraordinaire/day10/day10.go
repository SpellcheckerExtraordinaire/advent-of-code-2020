package day10

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type AoC10Solver uint

func countDifferences(adapters []int) (int, int) {
    oneCount := 0
    threeCount := 0
    last := 0
    for _, adapter := range adapters {
        difference := adapter - last
        if difference > 3 {
            fmt.Println("somethings not right")
        }
        if difference == 1 {
            oneCount++
        }
        if difference == 3 {
            threeCount++
        }
        last = adapter
    }
    return oneCount, threeCount
}

func countArrangementsCached(adapters []int, cur int, cache *map[int]int) int {
    cached, exists := (*cache)[adapters[cur]]
    if exists {
        return cached
    } else if cur == len(adapters) - 1 {
        return 1
    }
    branches := 0
    for i:= cur + 1; i < cur + 4 && i < len(adapters); i++{
        if adapters[i] - adapters[cur] <= 3 {
            branches += countArrangementsCached(adapters, i, cache)
        }
    }
    (*cache)[adapters[cur]] = branches
    return branches
}

func createAdapterList(input string) []int {
    scanner := bufio.NewScanner(strings.NewReader(input))
    scanner.Split(bufio.ScanLines)

    adapters := make([]int, 1, 1)
    adapters[0] = 0
    for scanner.Scan() {
        line := scanner.Text()
        adapter, _ := strconv.ParseInt(line, 10, 0)
        adapters = append(adapters, int(adapter))
    }
    sort.Ints(adapters)
    adapters = append(adapters, adapters[len(adapters) - 1] + 3)
    return adapters
}

func (me AoC10Solver) SolvePartOne(input string) {
    adapters := createAdapterList(input)
    oneCount, threeCount := countDifferences(adapters)
    fmt.Println("Result: " + strconv.Itoa(oneCount * threeCount))
}

func (me AoC10Solver) SolvePartTwo(input string) {
    adapters := createAdapterList(input)
    fmt.Println(adapters)
    cache := make(map[int]int)
    found := countArrangementsCached(adapters, 0, &cache)
    fmt.Println(cache)
    fmt.Println("Arrangements: " + strconv.Itoa(found))
}

func (me AoC10Solver) Day() uint {
    return uint(me)
}

func Solve(sampleOnly bool) {
    solver := AoC10Solver(10)
    aoc.SolvePuzzle(solver, sampleOnly)
}
