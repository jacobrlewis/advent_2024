package d1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1(file *os.File) int {
	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}

	// create left and right lists (assuming rows of 2 ints separated by whitespace)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		i, err := strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}
		left = append(left, i)

		i, err = strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}
		right = append(right, i)
	}

	// sort lists
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	// get sum of absolute difference for each index
	sum := 0
	for i := range left {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum
}

func Part2(file *os.File) int {
	fmt.Println("Day 1 part 2!")
	return 0
}
