package d7

import (
	"bufio"
	"os"
	"strings"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

type operation int

const (
	Add int = iota
	Mul
)

func solveLine(total int, nums []int) int {
	operators := make([]operation, len(nums)-1)

	for {
		calculation := nums[0]
		for i, o := range operators {
			if o == operation(Mul) {
				calculation *= nums[i+1]
			} else {
				calculation += nums[i+1]
			}
		}

		if total == calculation {
			return total
		}

		// binary increment on operators array
		for j := 0; j < len(operators)+1; j++ {

			// catch overflow error
			if j == len(operators) {
				// all values already set to true
				// next step would be overflow
				// this means no solution
				return 0
			}

			if operators[j] == operation(Mul) {
				// flip 1 to 0 as long as possible (carrying the 1)
				operators[j] = operation(Add)
				continue
			}
			// found a 0: flip to 1 and break loop, that was our increment
			operators[j] = operation(Mul)
			break
		}

	}

}

func Part1(file *os.File) int {

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		left := strings.Split(line, ":")[0]
		total := aoc.StringToInt(left)

		nums := aoc.GetNums(strings.Split(line, ":")[1])

		sum += solveLine(total, nums)
	}

	return sum
}

func Part2(file *os.File) int {
	return 0
}
