package main

import (
	"aoc-2020-go/day1/aoc1"
	"aoc-2020-go/day2/aoc2"
	"os"
)

func main() {
	os.Chdir("./day1")
	aoc1.Solve()
	os.Chdir("..")

	os.Chdir("./day2")
	aoc2.Solve()
	os.Chdir("..")
}
