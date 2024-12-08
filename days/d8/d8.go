package d8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type coord struct {
	I int
	J int
}

func Part1(file *os.File) int {

	locations := make(map[coord]bool)
	antennas := make(map[rune][]coord)

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

	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes[0]); j++ {
			if runes[i][j] == '.' || runes[i][j] == '#' {
				continue
			}

			// found antenna
			antennas[runes[i][j]] = append(antennas[runes[i][j]], coord{i, j})
		}
	}

	// for each antenna type calculate antinode locations
	for k := range antennas {
		for _, coord1 := range antennas[k] {
			for _, coord2 := range antennas[k] {
				if coord1 == coord2 {
					continue
				}

				nodeI := coord1.I - coord2.I + coord1.I
				nodeJ := coord1.J - coord2.J + coord1.J

				if nodeI >= 0 && nodeI < len(runes) && nodeJ >= 0 && nodeJ < len(runes[0]) {
					// antinode would be within map
					locations[coord{nodeI, nodeJ}] = true
				}
			}
		}
	}

	for c := range locations {
		runes[c.I][c.J] = '#'
	}

	for _, line := range runes {
		fmt.Println(string(line))
	}

	return len(locations)
}

func Part2(file *os.File) int {
	locations := make(map[coord]bool)
	antennas := make(map[rune][]coord)

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

	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes[0]); j++ {
			if runes[i][j] == '.' || runes[i][j] == '#' {
				continue
			}

			// found antenna
			antennas[runes[i][j]] = append(antennas[runes[i][j]], coord{i, j})
		}
	}

	// for each antenna type calculate antinode locations
	for k := range antennas {
		for _, coord1 := range antennas[k] {
			for _, coord2 := range antennas[k] {
				if coord1 == coord2 {
					continue
				}

				harmonicMult := 0
				// add a harmonic multiplier to the distance between the coords
				for {
					nodeI := harmonicMult*(coord1.I-coord2.I) + coord1.I
					nodeJ := harmonicMult*(coord1.J-coord2.J) + coord1.J

					if nodeI >= 0 && nodeI < len(runes) && nodeJ >= 0 && nodeJ < len(runes[0]) {
						// antinode would be within map
						locations[coord{nodeI, nodeJ}] = true
						harmonicMult++
					} else {
						break
					}
				}

			}
		}
	}

	for c := range locations {
		runes[c.I][c.J] = '#'
	}

	for _, line := range runes {
		fmt.Println(string(line))
	}

	return len(locations)
}
