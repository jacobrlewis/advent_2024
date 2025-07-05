package d14

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

func Part1(file *os.File) int {

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`-?\d+`)

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		x := aoc.StringToInt(matches[0])
		y := aoc.StringToInt(matches[1])
		dx := aoc.StringToInt(matches[2])
		dy := aoc.StringToInt(matches[3])

		// flip negative movement so modulo works
		if dx < 0 {
			dx = 101 + dx
		}
		if dy < 0 {
			dy = 103 + dy
		}

		finalX := (x + 100*dx) % 101
		finalY := (y + 100*dy) % 103

		fmt.Printf("x:%d y:%d dx:%d, dy:%d -> %d,%d\n", x, y, dx, dy, finalX, finalY)

		if finalX < 50 && finalY < 51 {
			q1 += 1
		} else if finalX > 50 && finalY < 51 {
			q2 += 1
		} else if finalX < 50 && finalY > 51 {
			q3 += 1
		} else if finalX > 50 && finalY > 51 {
			q4 += 1
		}
		fmt.Printf("%d %d %d %d\n", q1, q2, q3, q4)
	}

	return q1 * q2 * q3 * q4
}

type Bot struct {
	X  int
	Y  int
	DX int
	DY int
}

func Part2(file *os.File) int {

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`-?\d+`)

	bots := make([]Bot, 0)

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		x := aoc.StringToInt(matches[0])
		y := aoc.StringToInt(matches[1])
		dx := aoc.StringToInt(matches[2])
		dy := aoc.StringToInt(matches[3])

		// flip negative movement so modulo works
		if dx < 0 {
			dx = 101 + dx
		}
		if dy < 0 {
			dy = 103 + dy
		}

		bots = append(bots, Bot{x, y, dx, dy})
	}

	day := -1

	for {
		day += 1

		grid := make([][]int, 103)
		for i := range grid {
			grid[i] = make([]int, 101)
		}

		for i, bot := range bots {
			bots[i].X = (bot.DX + bot.X) % 101
			bots[i].Y = (bot.DY + bot.Y) % 103

			grid[bot.Y][bot.X] += 1
		}

		// find a tree by looking for a row with a lot of consecutive squares filled
		dayString := fmt.Sprintf("\n\nDay %d\n", day)
		found := false

		for _, line := range grid {
			streak := 0
			for _, num := range line {
				if num == 0 {
					dayString += "."
					streak = 0
				} else {
					dayString += fmt.Sprintf("%d", num)

					streak += 1
				}

				// I bumped up this number incrementally until I found one that printed out a tree
				if streak > 8 {
					found = true
				}
			}
			dayString += "\n"
		}

		if found {
			fmt.Println(dayString)
			break
		} else {
			fmt.Printf("%d ", day)
		}
	}

	return day
}
