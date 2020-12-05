package main

import (
	"aoc-2020-go/day1/aoc1"
	"os"
)

func main() {
	os.Chdir("./day1")
	aoc1.Solve()
	os.Chdir("..")
}
