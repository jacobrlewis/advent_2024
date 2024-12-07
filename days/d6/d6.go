package d6

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

func findGuardStart(lines []string) (int, int) {
	for y, line := range lines {
		for x, char := range line {
			if char == '^' {
				return x, y
			}
		}
	}
	panic("No guard!")
}

func nextCoord(x int, y int, d direction) (int, int) {
	switch d {
	case up:
		return x, y - 1
	case right:
		return x + 1, y
	case down:
		return x, y + 1
	case left:
		return x - 1, y
	default:
		panic("unknown direction")
	}
}

func Part1(file *os.File) int {

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

	x, y := findGuardStart(lines)
	direction := up

	// mark initial place
	touched := 1
	runes[y][x] = 'X'

	for {
		j, i := nextCoord(x, y, direction)

		if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[i]) {
			// made it to edge of map, break
			break
		}

		nextChar := runes[i][j]

		if nextChar == '#' {
			// hitting a wall, turn direction, do not move
			direction = (direction + 1) % 4
			continue
		}

		if nextChar == '.' {
			// moving to a unique spot, mark it and count it
			touched++
			runes[i][j] = 'X'
		}

		// move to next place
		x, y = j, i
	}

	// print finished path
	for _, line := range runes {
		fmt.Println(string(line))
	}

	return touched
}

func Part2(file *os.File) int {
	return 0
}
