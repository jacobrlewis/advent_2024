package d10

import (
	"fmt"
	"os"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

type coord struct {
	I int
	J int
}

func getCoordHeight(r rune) int {
	return int(r - '0')
}

// trailHeadScore finds all hiking trails from a head and returns how many reach a 9
func trailHeadScore(i int, j int, runes [][]rune) int {

	score := 0
	stack := []coord{{i, j}}

	visited := make(map[coord]bool)

	fmt.Printf("Checking %d %d", i, j)

	for {
		if len(stack) == 0 {
			break
		}

		// fmt.Println(stack)

		curCoord := stack[len(stack)-1]
		i = curCoord.I
		j = curCoord.J
		curHeight := getCoordHeight(runes[curCoord.I][curCoord.J])

		fmt.Printf("Currently at %+v height: %d\n", curCoord, curHeight)
		if curHeight == 9 {
			// found a peak
			score += 1
		}

		// make a move if possible
		if i > 0 && getCoordHeight(runes[i-1][j]) == curHeight+1 && !visited[coord{i - 1, j}] {
			// north
			fmt.Println("Moving North")
			stack = append(stack, coord{i - 1, j})
			continue
		}
		if j < len(runes[0])-1 && getCoordHeight(runes[i][j+1]) == curHeight+1 && !visited[coord{i, j + 1}] {
			// east
			fmt.Println("Moving east")
			stack = append(stack, coord{i, j + 1})
			continue
		}
		if i < len(runes)-1 && getCoordHeight(runes[i+1][j]) == curHeight+1 && !visited[coord{i + 1, j}] {
			// south
			fmt.Println("Moving south")
			stack = append(stack, coord{i + 1, j})
			continue
		}
		if j > 0 && getCoordHeight(runes[i][j-1]) == curHeight+1 && !visited[coord{i, j - 1}] {
			// west
			fmt.Println("Moving west")
			stack = append(stack, coord{i, j - 1})
			continue
		}

		// if no moves, move back in stack
		stack = stack[:len(stack)-1]
		visited[curCoord] = true
	}
	return score
}

func Part1(file *os.File) int {

	runes := aoc.Get2DRunes(file)
	sum := 0

	// find score of each 0
	for i, line := range runes {
		for j, r := range line {
			if r != '0' {
				continue
			}

			sum += trailHeadScore(i, j, runes)
		}
	}

	return sum
}

func trailHeadRating(i int, j int, runes [][]rune) int {
	score := 0
	curHeight := getCoordHeight(runes[i][j])
	// fmt.Printf("Currently at %d,%d height: %d\n", i, j, curHeight)
	if curHeight == 9 {
		// found a peak
		return 1
	}

	// make a move if possible
	if i > 0 && getCoordHeight(runes[i-1][j]) == curHeight+1 {
		// north
		// fmt.Println("Moving North")
		score += trailHeadRating(i-1, j, runes)
	}
	if j < len(runes[0])-1 && getCoordHeight(runes[i][j+1]) == curHeight+1 {
		// east
		// fmt.Println("Moving east")
		score += trailHeadRating(i, j+1, runes)
	}
	if i < len(runes)-1 && getCoordHeight(runes[i+1][j]) == curHeight+1 {
		// south
		// fmt.Println("Moving south")
		score += trailHeadRating(i+1, j, runes)
	}
	if j > 0 && getCoordHeight(runes[i][j-1]) == curHeight+1 {
		// west
		// fmt.Println("Moving west")
		score += trailHeadRating(i, j-1, runes)
	}

	return score
}

func Part2(file *os.File) int {
	runes := aoc.Get2DRunes(file)
	sum := 0

	// find score of each 0
	for i, line := range runes {
		for j, r := range line {
			if r != '0' {
				continue
			}

			sum += trailHeadRating(i, j, runes)
		}
	}

	return sum
}
