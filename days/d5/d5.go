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

// sortInvalidUpdate takes in an invalid update, and sorts by finding a page that should after something that follows it
// if a rule is found to broken, it is corrected by moving the first value directly after the found rule breaking value
func sortInvalidUpdate(rules map[string][]string, pages []string) []string {

	i := 0
	for i < len(pages) {

		innerLoopBroken := false

		for j := i + 1; j < len(pages); j++ {
			ruleBreaker := slices.Contains(rules[pages[i]], pages[j])

			if ruleBreaker {
				// rule broken, move the before to directly after the found match
				newSlice := make([]string, 0)
				newSlice = append(newSlice, pages[:i]...) // before the found value
				newSlice = append(newSlice, pages[i+1:j+1]...) // skip the found value, go to the AFTER val
				newSlice = append(newSlice, pages[i]) // add the found value
				newSlice = append(newSlice, pages[j+1:]...) // rest of the list, after the AFTER val
				pages = newSlice
				// break inner loop
				innerLoopBroken = true
				break
			}
		}

		if !innerLoopBroken {
			// current i index did not find a rule break, increment index
			i++
		}
	}

	return pages
}

// findValue returns the middle value from an invalid update after correcting it. 0 if given a valid update
func findUpdateValue(rules map[string][]string, line string) int {
	pages := strings.Split(line, ",")
	for i, v := range pages {
		// for each page, check if any of the following pages should have been BEFORE
		for _, after := range pages[i:] {
			if slices.Contains(rules[v], after) {
				// rule broken, sort and get value
				fmt.Printf("before sort %v\n", pages)
				sorted := sortInvalidUpdate(rules, pages)
				fmt.Printf("after sort %v\n", pages)
				return aoc.StringToInt(sorted[(len(sorted)-1)/2])
			}
		}
	}
	// no rules broken, get middle page
	return 0
}

func Part2(file *os.File) int {
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

		sum += findUpdateValue(rules, line)
	}

	return sum
}
