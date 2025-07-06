package d16

import (
	"fmt"
	"os"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

type Snapshot struct {
	X     int
	Y     int
	DX    int
	DY    int
	Score int
}

func Part1(file *os.File) int {
	maze := aoc.Get2DRunes(file)

	start := Snapshot{0, 0, 1, 0, 0}

	endX, endY := 0, 0

	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == 'S' {
				start.X = x
				start.Y = y
			}
			if maze[y][x] == 'E' {
				endX = x
				endY = y
				maze[y][x] = '.'
			}
		}
	}

	bestScores := make([][]int, len(maze))
	for i := range bestScores {
		bestScores[i] = make([]int, len(maze[0]))
	}

	branches := []Snapshot{start}

	for len(branches) > 0 {
		current := branches[0]
		branches = branches[1:]

		if current.X == endX && current.Y == endY {
			continue
		}

		// add next available moves
		if maze[current.Y+current.DY][current.X+current.DX] == '.' {
			// move straight
			nextScore := current.Score + 1
			nextX := current.X + current.DX
			nextY := current.Y + current.DY

			if nextScore < bestScores[nextY][nextX] || bestScores[nextY][nextX] == 0 {
				// only move if this is an improvement (or new)
				bestScores[nextY][nextX] = nextScore

				branches = append(branches,
					Snapshot{
						nextX,
						nextY,
						current.DX,
						current.DY,
						nextScore,
					})
			}
		}

		nextScore := current.Score + 1000

		leftDy := -1 * current.DX
		leftDx := current.DY

		if maze[current.Y+leftDy][current.X+leftDx] == '.' {
			bestScore := bestScores[current.Y+leftDy][current.X+leftDx]
			if nextScore+1 < bestScore || bestScore == 0 {
				// turn left in place
				branches = append(branches,
					Snapshot{
						current.X,
						current.Y,
						leftDx,
						leftDy,
						nextScore,
					})
			}
		}

		rightDy := -1 * leftDy
		rightDx := -1 * leftDx
		if maze[current.Y+rightDy][current.X+rightDx] == '.' {
			bestScore := bestScores[current.Y+rightDy][current.X+rightDx]
			if nextScore+1 < bestScore || bestScore == 0 {
				// turn right in place
				branches = append(branches,
					Snapshot{
						current.X,
						current.Y,
						rightDx,
						rightDy,
						nextScore,
					})
			}
		}
	}

	// print the paths used
	for i := range bestScores {
		for _, num := range bestScores[i] {
			fmt.Printf("%06d\t", num)
		}
		fmt.Println()
	}

	return bestScores[endY][endX]
}

func Part2(file *os.File) int {
	return 0
}
