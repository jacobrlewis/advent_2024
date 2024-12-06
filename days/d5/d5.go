package d5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

// validateUpdates return the middle page parsed to an int if it is valid
func validateUpdate(rules map[string][]string, line string) int {
	pages := strings.Split(line, ",")
	for i, v := range pages {
		// for each page, check if any of the following pages should have been BEFORE
		for _, after := range pages[i:] {
			if slices.Contains(rules[v], after) {
				// rule broken, continue
				fmt.Println("invalid - " + line)
				return 0
			}
		}
	}
	// no rules broken, get middle page
	fmt.Println("valid - " + line)
	return aoc.StringToInt(pages[(len(pages)-1)/2])
}

func Part1(file *os.File) int {
	sum := 0
	readingRules := true

	scanner := bufio.NewScanner(file)

	// make rules
	// access a number to get the list of numbers that should be BEFORE it
	rules := make(map[string][]string)
	ruleRegex := regexp.MustCompile(`(?P<before>\d+)\|(?P<after>\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// found the empty line splitting rules and pages
			readingRules = false
			fmt.Println(rules)
			continue
		}

		if readingRules {
			matches := ruleRegex.FindStringSubmatch(line)
			beforeIndex := ruleRegex.SubexpIndex("before")
			afterIndex := ruleRegex.SubexpIndex("after")

			before := matches[beforeIndex]
			after := matches[afterIndex]

			rules[after] = append(rules[after], before)
			continue
		}

		sum += validateUpdate(rules, line)
	}

	return sum
}

func Part2(file *os.File) int {
	return 0
}
