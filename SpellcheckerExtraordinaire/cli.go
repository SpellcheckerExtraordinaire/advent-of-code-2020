package main

import (
	"aoc-2020-go/day1"
	"aoc-2020-go/day2"
	"aoc-2020-go/day3"
	"aoc-2020-go/day4"
	"aoc-2020-go/day5"
	"aoc-2020-go/day6"
	"aoc-2020-go/day7"
	"aoc-2020-go/day8"
	"aoc-2020-go/day9"
    "aoc-2020-go/day10"
	"aoc-2020-go/day11"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var solvers = []func(bool){day1.Solve, day2.Solve, day3.Solve, day4.Solve, day5.Solve, day6.Solve, day7.Solve, day8.Solve, day9.Solve, day10.Solve, day11.Solve}

func run(day int, sampleOnly bool) {
	path := "./day" + strconv.Itoa(day+1)
	os.Chdir(path)
	solvers[day](sampleOnly)
	os.Chdir("..")
}

func runAll() {
	for day, _ := range solvers {
		run(day, false)
	}
}

func createNextDay() {
	// figure out which day we're on
	files, _ := ioutil.ReadDir("./")
	dirs := make([]os.FileInfo, 0, 25)
	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "day") {
			dirs = append(dirs, file)
		}
	}
	dayNum := strconv.Itoa(len(dirs) + 1)
	newDay := "./day" + dayNum

	os.Mkdir(newDay, 0777)

	// copy template, substituting the correct day
	{
		data, _ := ioutil.ReadFile("./aoc/solver-template.nogo")
		code := string(data)
		code = strings.ReplaceAll(code, "X", dayNum)
		ioutil.WriteFile(newDay+"/day"+dayNum+".go", []byte(code), 0777)
	}

	// create files for puzzle input
	os.Create(newDay + "/sample.txt")
	os.Create(newDay + "/puzzle.txt")

	// fetch puzzle input
	/*
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    2 * time.Second,
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}
		resp, _ := client.Get("https://adventofcode.com/")
		resp, _ = client.Get("https://adventofcode.com/2020/day/" + dayNum + "/input")
		//fmt.Println("https://adventofcode.com/2020/day/" + dayNum + "/input")
		fmt.Println(resp)

		This makes Aoc Server angry :(
	*/

	// modify myself >:)
	data, _ := ioutil.ReadFile("./cli.go")
	ioutil.WriteFile("./cli-backup.nogo", data, 0777)
	code := string(data)
	prevDay := strconv.Itoa(len(dirs))
	importPath := "aoc-2020-go/day" + prevDay + "\""
	prevSolver := ", day" + prevDay + ".Solve"
	code = strings.Replace(code, importPath, importPath+"\n\t\""+"aoc-2020-go/day"+dayNum+"\"", 1)
	code = strings.Replace(code, prevSolver, prevSolver+", day"+dayNum+".Solve", 1)
	ioutil.WriteFile("./cli.go", []byte(code), 0777)

}

func main() {
	argLength := len(os.Args[1:])
	// default case
	if argLength == 0 {
		runAll()
	} else if os.Args[1] == "nextDay" {
		createNextDay()
	} else {
		solverIndex, err := strconv.ParseInt(os.Args[1], 10, 0)
		if err != nil {
			fmt.Println("Invalid Argument, " + err.Error())
		}

		run(int(solverIndex-1), argLength == 2)
	}
}
