package day7

import (
	"aoc-2020-go/aoc"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type AoC7Solver uint

type Bag struct {
	name        string
	contains    map[string]int
	containedIn map[string]Bag
}

var stringToBag map[string]Bag = make(map[string]Bag)
var canContainGold map[string]bool = make(map[string]bool)

func (bag *Bag) String() string {
	containsString := "contains [ "
	for b, _ := range bag.contains {
		containsString += b + " "
	}
	containsString += "]"

	containedString := "is contained in [ "
	for _, b := range bag.containedIn {
		containedString += b.name + " "
	}
	containedString += "]"

	return bag.name + " " + containsString + ", " + containedString
}

func createOrRetrieveBag(color string) Bag {
	bag, exists := stringToBag[color]
	if !exists {
		bag = Bag{color, make(map[string]int), make(map[string]Bag)}
		stringToBag[color] = bag
	}
	return bag
}

func processRule(rule string) {
	split := strings.Fields(rule)
	color := split[0] + split[1]
	bag := createOrRetrieveBag(color)
	for i := 4; i < len(split)-2; i += 4 {
		amount, err := strconv.ParseInt(split[i], 10, 0)
		if err != nil {
			continue
		}
		newColor := split[i+1] + split[i+2]
		bag.contains[newColor] = int(amount)
		newBag := createOrRetrieveBag(newColor)
		newBag.containedIn[color] = bag
	}
}

func recursiveTraversal(cur Bag) {
	for color := range cur.containedIn {
		_, existsAlready := canContainGold[color]
		if existsAlready || color == "shinygold" {
			continue
		}

		canContainGold[color] = true
		recursiveTraversal(stringToBag[color])
	}
}

func summingTraversal(cur Bag) int {
	totalBags := 1
	for color, amount := range cur.contains {
		totalBags += amount * summingTraversal(stringToBag[color])
	}
	return totalBags
}

func (me AoC7Solver) SolvePartOne(input string) {
	stringToBag = make(map[string]Bag)
	canContainGold = make(map[string]bool)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		rule := scanner.Text()
		processRule(rule)
	}

	recursiveTraversal(stringToBag["shinygold"])
	fmt.Println("Number of bags that can contain shiny gold: " + strconv.Itoa(len(canContainGold)))
}

func (me AoC7Solver) SolvePartTwo(input string) {
	stringToBag = make(map[string]Bag)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		rule := scanner.Text()
		processRule(rule)
	}

	bagsNeeded := summingTraversal(stringToBag["shinygold"]) - 1

	fmt.Println("Bags needed: " + strconv.Itoa(bagsNeeded))
}

func (me AoC7Solver) Day() uint {
	return uint(me)
}

func Solve(sampleOnly bool) {
	solver := AoC7Solver(7)
	aoc.SolvePuzzle(solver, sampleOnly)
}
