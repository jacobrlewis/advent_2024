package d4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// matchesMas takes in the array of strings, and the coordinates of the x coordiante
// the x and y multipliers tell it what direction to check (pos = down and right)
func matchesMas(lines []string, i int, j int, xMultiplier int, yMultiplier int) int {
	n := 0

	m := lines[i+(1*yMultiplier)][j+(1*xMultiplier)]
	a := lines[i+(2*yMultiplier)][j+(2*xMultiplier)]
	s := lines[i+(3*yMultiplier)][j+(3*xMultiplier)]

	if m == 'M' && a == 'A' && s == 'S' {
		n = 1
	}

	return n
}

func Part1(file *os.File) int {
	count := 0

	bytes, err := io.ReadAll(bufio.NewReader(file))

	if err != nil {
		panic(err)
	}

	// 2D array of chars
	lines := strings.Split(string(bytes), "\n")

	for i, line := range lines {
		for j, char := range line {
			if char == 'X' {

				rightSafe := j < len(line)-3
				leftSafe := j >= 3
				upSafe := i >= 3
				downSafe := i < len(line)-3

				if rightSafe {
					// EAST
					count += matchesMas(lines, i, j, 1, 0)

					if upSafe {
						// NORTHEAST
						count += matchesMas(lines, i, j, 1, -1)
					}
					if downSafe {
						// SOUTHEAST
						count += matchesMas(lines, i, j, 1, 1)
					}
				}

				if leftSafe {
					// WEST
					count += matchesMas(lines, i, j, -1, 0)

					if upSafe {
						// NORHTWEST
						count += matchesMas(lines, i, j, -1, -1)
					}
					if downSafe {
						// SOUTHWEST
						count += matchesMas(lines, i, j, -1, 1)
					}
				}

				if upSafe {
					// NORTH
					count += matchesMas(lines, i, j, 0, -1)
				}

				if downSafe {
					// SOUTH
					count += matchesMas(lines, i, j, 0, 1)
				}
			}
		}
	}

	return count
}

// delete empty line at end of file
func Part2(file *os.File) int {

	count := 0

	bytes, err := io.ReadAll(bufio.NewReader(file))

	if err != nil {
		panic(err)
	}

	// 2D array of chars
	lines := strings.Split(string(bytes), "\n")

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == 'A' {
				fmt.Println(i, j)
				// S	or	M
				//  A		 A
				//   M		  S
				sam1 := lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M'
				mas1 := lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S'

				//   M	or 	  S
				//  A		 A
				// S		M
				mas2 := lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S'
				sam2 := lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M'

				if (mas1 || sam1) && (mas2 || sam2) {
					count += 1
				}
			}
		}
	}

	return count
}
