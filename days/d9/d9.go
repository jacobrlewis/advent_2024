package d9

import (
	"bufio"
	"os"
	"slices"
)

func Part1(file *os.File) int {

	scanner := bufio.NewScanner(file)

	// read the single line
	scanner.Scan()

	blocks := make([]rune, 0)
	id := 0

	// create initial dot separated string
	for i, r := range []rune(scanner.Text()) {
		count := int(r - '0')

		char := '.'
		if i%2 == 0 {
			char = rune(id) + '0'
			id++
		}

		blocks = append(blocks, slices.Repeat([]rune{char}, count)...)
	}

	// fmt.Println(string(blocks))

	// sort, move last elements into open spaces
	for i, r := range blocks {
		// loop until we find an empty spot
		if r != '.' {
			continue
		}

		// loop backwards to find the last non '.' char
		for j := range len(blocks) {
			backIndex := len(blocks) - 1 - j

			if blocks[backIndex] == '.' {
				continue
			}

			if backIndex < i {
				// all blocks moved
				break
			}

			blocks[i] = blocks[backIndex]
			blocks[backIndex] = '.'
			break
		}
	}

	// fmt.Println(string(blocks))

	// get value
	sum := 0

	for i, r := range blocks {
		if r == '.' {
			break
		}
		sum += i * int(r-'0')
	}

	return sum
}

type area struct {
	Index int
	Size  int
}

func Part2(file *os.File) int {

	scanner := bufio.NewScanner(file)

	// read the single line
	scanner.Scan()

	blocks := make([]rune, 0)
	id := 0

	files := make([]area, 0)
	spaces := make([]area, 0)

	// create initial dot separated string
	for i, r := range []rune(scanner.Text()) {
		count := int(r - '0')

		char := '.'
		if i%2 == 0 {
			char = rune(id) + '0'
			id++
			// insert files in reverse order
			files = append(files, area{len(blocks), count})
		} else {
			spaces = append(spaces, area{len(blocks), count})
		}

		blocks = append(blocks, slices.Repeat([]rune{char}, count)...)
	}

	// sort, move last file into leftmost space that can fit it
	slices.Reverse(files)
	for fileIndex := range files {
		file := files[fileIndex]

		// go through spaces
		for spaceIndex := range spaces {
			space := spaces[spaceIndex]

			if space.Index > file.Index {
				break
			}

			if file.Size <= space.Size {
				// file fits
				fileValue := blocks[file.Index]

				for i := range file.Size {
					// fill space
					blocks[space.Index+i] = fileValue

					// overwrite file
					blocks[file.Index+i] = '.'
				}
				files[fileIndex].Index = space.Index
				spaces[spaceIndex].Index += file.Size
				spaces[spaceIndex].Size -= file.Size
				break
			}
		}
	}

	// get value
	sum := 0

	for i, r := range blocks {
		if r == '.' {
			continue
		}
		sum += i * int(r-'0')
	}

	return sum
}
