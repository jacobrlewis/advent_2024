package d13

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

type Machine struct {
	Ax     int
	Ay     int
	Bx     int
	By     int
	PrizeX int
	PrizeY int
}

func ParseMachine(buttonA string, buttonB string, prize string) Machine {
	m := Machine{}

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(buttonA, -1)
	m.Ax = aoc.StringToInt(matches[0])
	m.Ay = aoc.StringToInt(matches[1])

	matches = re.FindAllString(buttonB, -1)
	m.Bx = aoc.StringToInt(matches[0])
	m.By = aoc.StringToInt(matches[1])

	matches = re.FindAllString(prize, -1)
	m.PrizeX = aoc.StringToInt(matches[0])
	m.PrizeY = aoc.StringToInt(matches[1])

	return m
}

func ReadMachines(file *os.File) []Machine {
	machines := make([]Machine, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read machine lines
		buttonA := scanner.Text()
		scanner.Scan()
		buttonB := scanner.Text()
		scanner.Scan()
		prize := scanner.Text()
		scanner.Scan()

		// skip newline
		scanner.Text()

		machines = append(machines, ParseMachine(buttonA, buttonB, prize))
	}
	return machines
}

// ScoreMachine returns the token cost for a machine if possible, or 0 if invalid
func ScoreMachine(m Machine) int {
	min_score := 1000 // score can't be more than 100 A presses (3 * 100)
	found := false

	for a := range 100 {
		aSum := a * m.Ax
		bSum := m.PrizeX - aSum
		isClean := bSum%m.Bx == 0
		if !isClean {
			continue
		}

		b := bSum / m.Bx
		fmt.Printf("X: Found a,b %d,%d\n", a, b)

		// check that the a,b scores work for Y as well
		if !((a*m.Ay + b*m.By) == m.PrizeY) {
			continue
		}

		fmt.Printf("XY: Found a,b %d,%d\n", a, b)
		min_score = min(min_score, (3*a + b))
		found = true
	}

	if !found {
		return 0
	}

	return min_score
}

func Part1(file *os.File) int {
	total := 0

	machines := ReadMachines(file)
	fmt.Printf("%+v\n", machines)

	for _, m := range machines {
		total += ScoreMachine(m)
	}

	return total
}


func Part2(file *os.File) int {
	total := 0

	return total
}
