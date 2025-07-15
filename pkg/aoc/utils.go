package aoc

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// GetNums returns a slice of ints from a row of numbers split by whitespace
func GetNums(row string) []int {
	words := strings.Fields(row)

	ints := make([]int, len(words))

	for i, v := range words {
		ints[i] = StringToInt(v)
	}

	return ints
}

// StringToInt converts a string to an int, panics on err
func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func Get2DRunes(file *os.File) [][]rune {
	bytes, err := io.ReadAll(bufio.NewReader(file))

	if err != nil {
		panic(err)
	}

	// 2D array of runes
	lines := strings.Split(string(bytes), "\n")
	runes := make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}
	return runes
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// GCD returns the greatest common divisor of 2 integers
func GCD(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// n ** m
func IntPow(n int, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
