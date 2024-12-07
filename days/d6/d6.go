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

type path struct {
	Runes [][]rune
}

// given the current coord and direction, test if the guard gets stuck in a loop
func (p path) findLoop(x int, y int, d direction) bool {

	currentDir := d

	// any path longer than the area of the map is a loop
	maxSteps := len(p.Runes) * len(p.Runes[0])
	loopCheckSteps := 0

	for {

		j, i := nextCoord(x, y, currentDir)

		if i < 0 || i >= len(p.Runes) || j < 0 || j >= len(p.Runes[i]) {
			// made it to edge of map, no loop
			return false
		}

		if p.Runes[i][j] == '#' {
			// hitting a wall, turn direction, do not move
			currentDir = (currentDir + 1) % 4
			continue
		}

		// move to next place
		x, y = j, i
		loopCheckSteps++

		if loopCheckSteps > maxSteps {
			return true
		}
	}
}

func Part2(file *os.File) int {
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

	path := path{runes}

	obstacleCount := 0

	x, y := findGuardStart(lines)
	runes[y][x] = 'X'
	currentDir := up

	for {
		j, i := nextCoord(x, y, currentDir)

		if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[i]) {
			// made it to edge of map, break
			break
		}

		nextChar := runes[i][j]

		if nextChar == '#' {
			// hitting a wall, turn direction, do not move
			currentDir = (currentDir + 1) % 4
			continue
		}

		if nextChar == '.' {
			// moving to a new spot
			runes[i][j] = '#'
			loopFound := path.findLoop(x, y, currentDir)
			runes[i][j] = 'X'

			if loopFound {
				// each location is only tested for loops once, safe to just increment a counter
				obstacleCount++
			}
		}
		// no need to test if the next square has already been visited
		// it would have been tested earlier (before moving to it)

		// move to next place
		x, y = j, i
	}

	return obstacleCount
}
