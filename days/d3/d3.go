package d3

import (
	"bufio"
	"os"
	"regexp"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

func Part1(file *os.File) int {

	scanner := bufio.NewScanner(file)
	sum := 0

	// starts with mul( - any number of ints, comma, any number of ints, closing paren
	re := regexp.MustCompile(`mul\((?P<first>[0-9]+),(?P<second>[0-9]+)\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		firstIndex := re.SubexpIndex("first")
		secondIndex := re.SubexpIndex("second")

		for i := range matches {
			first := matches[i][firstIndex]
			second := matches[i][secondIndex]

			sum += aoc.StringToInt(first) * aoc.StringToInt(second)
		}
	}

	return sum
}

func Part2(file *os.File) int {

	scanner := bufio.NewScanner(file)
	sum := 0

	// starts with mul( - any number of ints, comma, any number of ints, closing paren
	// OR matches do()
	// OR matches don't()
	re := regexp.MustCompile(`mul\((?P<first>[0-9]+),(?P<second>[0-9]+)\)|(?P<activate>do\(\))|(?P<deactivate>don\'t\(\))`)

	active := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		firstIndex := re.SubexpIndex("first")
		secondIndex := re.SubexpIndex("second")
		activateIndex := re.SubexpIndex("activate")
		deactivateIndex := re.SubexpIndex("deactivate")

		for i := range matches {

			// fmt.Println(matches[i])
			if matches[i][activateIndex] == "do()" {
				// matched on "do()"
				// fmt.Println("matched do()")
				active = true
				continue
			}

			if matches[i][deactivateIndex] == "don't()" {
				// matched on "don't"
				// fmt.Println("matched don't()")
				active = false
				continue
			}

			// if the start of the file, or matched on do() more recently than don't()
			if active {
				first := matches[i][firstIndex]
				second := matches[i][secondIndex]

				sum += aoc.StringToInt(first) * aoc.StringToInt(second)
			}

		}
	}

	return sum
}
