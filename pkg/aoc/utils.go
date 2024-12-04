package aoc

import (
	"strconv"
	"strings"
)

// GetNums returns a slice of ints from a row of numbers
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
