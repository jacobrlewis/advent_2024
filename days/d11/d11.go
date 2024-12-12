package d11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

func Part1(file *os.File) int {

	// read the single line
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	nums := aoc.GetNums(scanner.Text())
	fmt.Printf("Initial: %v\n", nums)

	for range 25 {
		newNums := make([]int, 0)
		for _, num := range nums {

			if num == 0 {
				newNums = append(newNums, 1)
				continue
			}

			str := strconv.Itoa(num)
			strLength := len(str)
			if strLength%2 == 0 {

				first := str[:len(str)/2]
				last := str[len(str)/2:]

				newNums = append(newNums, aoc.StringToInt(first))
				newNums = append(newNums, aoc.StringToInt(last))
				continue
			}

			newNums = append(newNums, num*2024)
		}
		nums = newNums
		// fmt.Printf("%d: %v\n", blink+1, nums)
	}

	return len(nums)
}

type state struct {
	I          int
	blinksLeft int
}

type solver struct {
	cache map[state]int
}

func (s solver) getResult(i int, blinksLeft int) int {
	// fmt.Printf("i: %d blinks: %d\n", i, blinksLeft)
	curState := state{i, blinksLeft}
	ans, ok := s.cache[curState]
	if ok {
		return ans
	}

	if blinksLeft == 0 {
		s.cache[curState] = 1
		return 1
	}

	if i == 0 {
		return s.getResult(1, blinksLeft-1)
	}

	str := strconv.Itoa(i)
	strLength := len(str)
	if strLength%2 == 0 {

		first := str[:len(str)/2]
		last := str[len(str)/2:]

		left := s.getResult(aoc.StringToInt(first), blinksLeft-1)
		right := s.getResult(aoc.StringToInt(last), blinksLeft-1)
		s.cache[curState] = left + right

		return left + right
	}

	ans = s.getResult(i*2024, blinksLeft-1)
	s.cache[curState] = ans
	return ans
}

func Part2(file *os.File) int {

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	nums := aoc.GetNums(scanner.Text())
	sum := 0
	cache := make(map[state]int)
	s := solver{cache}

	for _, num := range nums {
		sum += s.getResult(num, 75)
	}

	return sum
}
