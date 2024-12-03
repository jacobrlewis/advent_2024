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
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}

	return ints
}
