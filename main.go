package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jacobrlewis/advent_2024/days/d1"
	"github.com/jacobrlewis/advent_2024/days/d10"
	"github.com/jacobrlewis/advent_2024/days/d11"
	"github.com/jacobrlewis/advent_2024/days/d2"
	"github.com/jacobrlewis/advent_2024/days/d3"
	"github.com/jacobrlewis/advent_2024/days/d4"
	"github.com/jacobrlewis/advent_2024/days/d5"
	"github.com/jacobrlewis/advent_2024/days/d6"
	"github.com/jacobrlewis/advent_2024/days/d7"
	"github.com/jacobrlewis/advent_2024/days/d8"
	"github.com/jacobrlewis/advent_2024/days/d9"
)

var day int
var part int
var example bool
var customFile string

func main() {
	flag.IntVar(&day, "d", 1, "Day (integer) to run")
	flag.IntVar(&part, "p", 1, "Part (integer) to run (1 or 2)")
	flag.BoolVar(&example, "example", false, "If given, uses example input file")
	flag.StringVar(&customFile, "file", "", "If given uses this file name + .txt")
	flag.Parse()

	// get problem to call
	funcs := [][]func(*os.File) int{
		{d1.Part1, d1.Part2},
		{d2.Part1, d2.Part2},
		{d3.Part1, d3.Part2},
		{d4.Part1, d4.Part2},
		{d5.Part1, d5.Part2},
		{d6.Part1, d6.Part2},
		{d7.Part1, d7.Part2},
		{d8.Part1, d8.Part2},
		{d9.Part1, d9.Part2},
		{d10.Part1, d10.Part2},
		{d11.Part1, d11.Part2}}
	problem := funcs[day-1][part-1]

	// get file input
	fileName := "input"
	if example {
		fileName = "example"
	}
	if customFile != "" {
		fileName = customFile
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
