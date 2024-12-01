package d1

import (
	"bufio"
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
	scanner := bufio.NewScanner(file)

	left := []int{}
	right := make(map[int]int)

	// create left list
	// create right map of occurances
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
		right[i] += 1
	}

	score := 0
	// for each value in the left list, multiply by number of occurances in the right list
	for _, v := range left {
		x := v * right[v]
		score += x
	}

	return score
}
