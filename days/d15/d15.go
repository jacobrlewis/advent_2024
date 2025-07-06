package d15

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Warehouse struct {
	Grid [][]rune
	X    int
	Y    int
}

// makeMove moves the robot in the direction provided
func (w *Warehouse) makeMove(dir rune) {
	dx, dy := 0, 0

	switch dir {
	case '^':
		dy = -1
	case '<':
		dx = -1
	case '>':
		dx = 1
	case 'v':
		dy = 1
	}

	nextX, nextY := w.X+dx, w.Y+dy

	// loop until hitting wall
	for w.Grid[nextY][nextX] != '#' {
		if w.Grid[nextY][nextX] == '.' {
			// found an empty square
			// place box into the free square (if not pushing a box, this will be overwritten correctly)
			w.Grid[nextY][nextX] = 'O'

			// move bot
			w.Grid[w.Y][w.X] = '.'
			w.X += dx
			w.Y += dy
			w.Grid[w.Y][w.X] = '@'
			break
		}

		nextX += dx
		nextY += dy
	}
}

// getScore returns the sum of all GPS coordinates of all boxes
// GPS = 100 * y + x
func (w Warehouse) getScore() int {
	sum := 0
	for y := range w.Grid {
		for x := range w.Grid[y] {
			if w.Grid[y][x] == 'O' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func (w Warehouse) print() {
	for y := range w.Grid {
		fmt.Println(string(w.Grid[y]))
	}
	fmt.Println("")
}

func Part1(file *os.File) int {

	w := Warehouse{}
	y := 0
	grid := make([][]rune, 0)

	// read until end of map
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			// found end of map
			break
		}

		if strings.Contains(line, "@") {
			// found starting spot, save it
			w.Y = y
			w.X = strings.Index(line, "@")
		}
		grid = append(grid, []rune(line))

		y += 1
	}
	w.Grid = grid

	// handle instructions
	for scanner.Scan() {
		steps := scanner.Text()

		for _, dir := range steps {
			fmt.Printf("Move %c:\n", dir)
			w.makeMove(dir)
			w.print()
		}
	}

	return w.getScore()
}

type DoubleWarehouse struct {
	Grid [][]rune
	X    int
	Y    int
}

func (w DoubleWarehouse) print() {
	for y := range w.Grid {
		fmt.Println(string(w.Grid[y]))
	}
	fmt.Println("")
}

// getScore returns the sum of all GPS coordinates of all boxes
// GPS = 100 * y + x
func (w DoubleWarehouse) getScore() int {
	sum := 0
	for y := range w.Grid {
		for x := range w.Grid[y] {
			if w.Grid[y][x] == '[' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

type Coord struct {
	X int
	Y int
}

// makeMove moves the robot in the direction provided
func (w *DoubleWarehouse) makeMove(dir rune) {
	dx, dy := 0, 0

	switch dir {
	case '^':
		dy = -1
	case '<':
		dx = -1
	case '>':
		dx = 1
	case 'v':
		dy = 1
	}

	nextX, nextY := w.X+dx, w.Y+dy

	checkSquares := []Coord{{nextX, nextY}}

	// boxes are stored by leftX and Y coord
	pushingBoxes := make([]Coord, 0)

	// get list of all boxes that are being affected
	// also check if there is a wall in the way
	for len(checkSquares) != 0 {
		checking := checkSquares[0]
		checkSquares = checkSquares[1:]

		if w.Grid[checking.Y][checking.X] == '[' {
			pushingBoxes = append(pushingBoxes, checking)

			// add new checking squares
			if dy == 0 {
				// we know we are checking to the right here
				checkSquares = append(checkSquares, Coord{checking.X + 2, checking.Y})
			} else {
				// add the 2 squares above/below where we found this fox
				checkSquares = append(checkSquares, Coord{checking.X, checking.Y + dy})
				checkSquares = append(checkSquares, Coord{checking.X + 1, checking.Y + dy})
			}
		}

		if w.Grid[checking.Y][checking.X] == ']' {
			pushingBoxes = append(pushingBoxes, Coord{checking.X - 1, checking.Y})

			// add new checking squares
			if dy == 0 {
				// we know we are checking to the left here
				checkSquares = append(checkSquares, Coord{checking.X - 2, checking.Y})
			} else {
				// add the 2 squares above/below where we found this fox
				checkSquares = append(checkSquares, Coord{checking.X - 1, checking.Y + dy})
				checkSquares = append(checkSquares, Coord{checking.X, checking.Y + dy})
			}
		}

		if w.Grid[checking.Y][checking.X] == '#' {
			// hitting a wall, no boxes can be pushed
			return
		}
	}

	// move each box
	slices.Reverse(pushingBoxes)
	for _, box := range pushingBoxes {
		w.Grid[box.Y][box.X] = '.'
		w.Grid[box.Y][box.X+1] = '.'

		w.Grid[box.Y+dy][box.X+dx] = '['
		w.Grid[box.Y+dy][box.X+dx+1] = ']'
	}

	// move bot
	w.Grid[w.Y][w.X] = '.'
	w.X += dx
	w.Y += dy
	w.Grid[w.Y][w.X] = '@'
}

func Part2(file *os.File) int {

	w := DoubleWarehouse{}
	y := 0
	grid := make([][]rune, 0)

	// read until end of map
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			// found end of map
			break
		}

		// double width of everything
		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")

		if strings.Contains(line, "@") {
			// found starting spot, save it
			w.Y = y
			w.X = strings.Index(line, "@")
		}
		grid = append(grid, []rune(line))

		y += 1
	}
	w.Grid = grid

	w.print()

	// handle instructions
	for scanner.Scan() {
		steps := scanner.Text()

		for _, dir := range steps {
			fmt.Printf("Move %c:\n", dir)
			w.makeMove(dir)
			w.print()
		}
	}

	return w.getScore()
}
