package d2

import (
	"bufio"
	"os"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

func rowSafe(nums []int, checkIndex int, increasing bool) bool {
	// reached end of list, must be safe
	if len(nums)-1 == checkIndex {
		return true
	}

	diff := nums[checkIndex] - nums[checkIndex+1]
	if increasing {
		diff = -diff
	}

	if 1 <= diff && diff <= 3 {
		// this index was safe, check the next one
		return rowSafe(nums, checkIndex+1, increasing)
	}
	return false
}

func Part1(file *os.File) int {

	scanner := bufio.NewScanner(file)
	safe := 0

	for scanner.Scan() {
		nums := aoc.GetNums(scanner.Text())
		if rowSafe(nums, 0, nums[0] < nums[1]) {
			safe += 1
		}
	}

	return safe
}

func removeIndex(nums []int, index int) []int {
	newNums := make([]int, 0)
	newNums = append(newNums, nums[:index]...)
	newNums = append(newNums, nums[index+1:]...)
	return newNums
}

func Part2(file *os.File) int {

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		nums := aoc.GetNums(scanner.Text())
		// check original list
		isSafe := rowSafe(nums, 0, nums[0] < nums[1])
		index := 0

		// loop until a safe list is found, or until each element has been removed and failed
		for !isSafe && index < len(nums) {
			newNums := removeIndex(nums, index)
			isSafe = rowSafe(newNums, 0, newNums[0] < newNums[1])
			index += 1
		}
		if isSafe {
			safeCount += 1
		}
	}

	return safeCount
}
