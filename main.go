package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jacobrlewis/advent_2024/days/d1"
	"github.com/jacobrlewis/advent_2024/days/d2"
)

var day int
var part int
var example bool

func main() {
	flag.IntVar(&day, "d", 1, "Day (integer) to run")
	flag.IntVar(&part, "p", 1, "Part (integer) to run (1 or 2)")
	flag.BoolVar(&example, "example", false, "If given, uses example input file")
	flag.Parse()

	// get problem to call
	funcs := [][]func(*os.File) int{
		{d1.Part1, d1.Part2},
		{d2.Part1, d2.Part2}}
	problem := funcs[day-1][part-1]

	// get file input
	fileName := "input"
	if example {
		fileName = "example"
	}
	filePath := fmt.Sprintf("./days/d%d/%s.txt", day, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// solve problem
	fmt.Printf("Day %d Part %d...\n", day, part)
	fmt.Printf("Solution: %d\n", problem(file))
}
