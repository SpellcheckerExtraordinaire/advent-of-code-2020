package main

import (
	"aoc-2020-go/day1/aoc1"
	"aoc-2020-go/day2/aoc2"
	"aoc-2020-go/day3/aoc3"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var solvers = []func(){aoc1.Solve, aoc2.Solve, aoc3.Solve}

func runAll() {
	for day, solver := range solvers {
		path := "./day" + strconv.Itoa(day+1)
		fmt.Println(path)
		os.Chdir(path)
		solver()
		os.Chdir("..")
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

	os.Mkdir(newDay, os.ModeDir)
	os.Mkdir(newDay+"/aoc"+dayNum, os.ModeDir)

	// copy template, substituting the correct day
	{
		data, _ := ioutil.ReadFile("./aoc/solver-template.nogo")
		code := string(data)
		code = strings.ReplaceAll(code, "X", dayNum)
		ioutil.WriteFile(newDay+"/aoc"+dayNum+"/aoc"+dayNum+".go", []byte(code), 0666)
	}

	{
		data, _ := ioutil.ReadFile("./aoc/runner.nogo")
		code := string(data)
		code = strings.ReplaceAll(code, "X", dayNum)
		ioutil.WriteFile(newDay+"/runDay"+dayNum+".go", []byte(code), 0666)
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
	ioutil.WriteFile("./cli-backup.nogo", data, 0666)
	code := string(data)
	prevDay := strconv.Itoa(len(dirs))
	importPath := "aoc-2020-go/day" + prevDay + "/aoc" + prevDay + "\""
	prevSolver := ", aoc" + prevDay + ".Solve"
	code = strings.Replace(code, importPath, importPath+"\n\t\""+"aoc-2020-go/day"+dayNum+"/aoc"+dayNum+"\"", 1)
	code = strings.Replace(code, prevSolver, prevSolver+", aoc"+dayNum+".Solve", 1)
	ioutil.WriteFile("./cli.go", []byte(code), 0666)

}

func main() {
	argLength := len(os.Args[1:])
	// default case
	if argLength == 0 {
		runAll()
	} else if os.Args[1] == "nextDay" {
		createNextDay()
	}
}
