package d12

import (
	"fmt"
	"os"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

type coord struct {
	i int
	j int
}

type garden struct {
	runes   [][]rune
	regions map[coord]int
}

func (g garden) markRegion(c coord, regionId int) {
	// depth first search to fill region
	g.regions[c] = regionId
	r := g.runes[c.i][c.j]
	var nextCoord coord

	if c.i != 0 {
		nextCoord = coord{c.i - 1, c.j}

		if g.runes[nextCoord.i][nextCoord.j] == r && g.regions[nextCoord] == 0 {
			g.markRegion(nextCoord, regionId)
		}
	}

	if c.i != len(g.runes)-1 {
		nextCoord = coord{c.i + 1, c.j}

		if g.runes[nextCoord.i][nextCoord.j] == r && g.regions[nextCoord] == 0 {
			g.markRegion(nextCoord, regionId)
		}
	}

	if c.j != 0 {
		nextCoord = coord{c.i, c.j - 1}

		if g.runes[nextCoord.i][nextCoord.j] == r && g.regions[nextCoord] == 0 {
			g.markRegion(nextCoord, regionId)
		}
	}

	if c.j != len(g.runes[0])-1 {
		nextCoord = coord{c.i, c.j + 1}

		if g.runes[nextCoord.i][nextCoord.j] == r && g.regions[nextCoord] == 0 {
			g.markRegion(nextCoord, regionId)
		}
	}
}

func Part1(file *os.File) int {
	total := 0

	runes := aoc.Get2DRunes(file)
	newRegionId := 1
	regions := make(map[coord]int)

	g := garden{runes, regions}

	areas := make(map[int]int)
	perims := make(map[int]int)

	for i, line := range runes {
		for j, r := range line {

			perim := 0

			// north
			if i == 0 || runes[i-1][j] != r {
				perim += 1
			}
			// east
			if j == 0 || runes[i][j-1] != r {
				perim += 1
			}
			// south
			if i == len(runes)-1 || runes[i+1][j] != r {
				perim += 1
			}
			// west
			if j == len(line)-1 || runes[i][j+1] != r {
				perim += 1
			}

			curRegion := g.regions[coord{i, j}]

			if curRegion == 0 {
				fmt.Printf("New region at %d %d\n", i, j)
				curRegion = newRegionId
				newRegionId++
				g.markRegion(coord{i, j}, curRegion)
			}

			areas[curRegion] += 1
			perims[curRegion] += perim
		}
	}

	for k := range areas {
		total += areas[k] * perims[k]
		// fmt.Println(total)
	}

	fmt.Println(areas)
	fmt.Println(perims)

	return total
}

type border struct {
	i         int
	j         int
	direction aoc.Direction
}

type garden2 struct {
	runes [][]rune
	// coord -> regionId
	regions map[coord]int
	// regionId -> (coord -> direction)
	border map[int]map[border]bool
	// regionId -> num sides for that region
	sides map[int]int
}

func (g garden2) addBorder(c coord, d aoc.Direction, regionId int) {
	_, ok := g.border[regionId]

	if !ok {
		g.border[regionId] = make(map[border]bool)
	}
	borderMap := g.border[regionId]

	var leftExists bool
	var rightExists bool

	switch d {
	case aoc.North:
		leftExists = borderMap[border{c.i, c.j - 1, d}]
		rightExists = borderMap[border{c.i, c.j + 1, d}]
	case aoc.East:
		leftExists = borderMap[border{c.i - 1, c.j, d}]
		rightExists = borderMap[border{c.i + 1, c.j, d}]
	case aoc.South:
		leftExists = borderMap[border{c.i, c.j - 1, d}]
		rightExists = borderMap[border{c.i, c.j + 1, d}]
	case aoc.West:
		leftExists = borderMap[border{c.i - 1, c.j, d}]
		rightExists = borderMap[border{c.i + 1, c.j, d}]
	}

	borderMap[border{c.i, c.j, d}] = true
	if leftExists && rightExists {
		// this means we are checking a border that would join two existing sides
		// only one of them should have been counted. reduce from count.
		g.sides[regionId] -= 1
		return
	}

	if !leftExists && !rightExists {
		g.sides[regionId] += 1
	}

}

func (g garden2) markRegionTrackBorders(c coord, regionId int) {
	// depth first search to fill region
	// fmt.Printf("Marking %v for region %d\n", c, regionId)
	g.regions[c] = regionId
	r := g.runes[c.i][c.j]
	var nextCoord coord

	// BFS to travel to each coord in a region

	// regardless of order, there can always be shapes that result in checking parts
	// of the same side, with coords in the center of the side not being checked yet
	// this means temporarily overcounting sides
	// when the holes are discovered, the overcounting is corrected
	nextCoord = coord{c.i - 1, c.j}
	if c.i != 0 && g.runes[nextCoord.i][nextCoord.j] == r {
		if g.regions[nextCoord] == 0 {
			g.markRegionTrackBorders(nextCoord, regionId)
		}
	} else {
		g.addBorder(nextCoord, aoc.North, regionId)
	}

	nextCoord = coord{c.i, c.j + 1}
	if c.j != len(g.runes[0])-1 && g.runes[nextCoord.i][nextCoord.j] == r {
		if g.regions[nextCoord] == 0 {
			g.markRegionTrackBorders(nextCoord, regionId)
		}
	} else {
		g.addBorder(nextCoord, aoc.East, regionId)
	}

	nextCoord = coord{c.i + 1, c.j}
	if c.i != len(g.runes)-1 && g.runes[nextCoord.i][nextCoord.j] == r {
		if g.regions[nextCoord] == 0 {
			g.markRegionTrackBorders(nextCoord, regionId)
		}
	} else {
		g.addBorder(nextCoord, aoc.South, regionId)
	}

	nextCoord = coord{c.i, c.j - 1}
	if c.j != 0 && g.runes[nextCoord.i][nextCoord.j] == r {
		if g.regions[nextCoord] == 0 {
			g.markRegionTrackBorders(nextCoord, regionId)
		}
	} else {
		g.addBorder(nextCoord, aoc.West, regionId)
	}

}

func Part2(file *os.File) int {
	total := 0

	runes := aoc.Get2DRunes(file)
	newRegionId := 1
	regions := make(map[coord]int)

	g := garden2{runes, regions, make(map[int]map[border]bool), make(map[int]int)}

	areas := make(map[int]int)

	for i, line := range runes {
		for j := range line {

			curRegion := g.regions[coord{i, j}]

			if curRegion == 0 {
				// fmt.Printf("New region at %d %d\n", i, j)
				curRegion = newRegionId
				newRegionId++
				g.markRegionTrackBorders(coord{i, j}, curRegion)
			}

			areas[curRegion] += 1
		}
	}

	for regionId := range areas {
		// fmt.Printf("RegionId: %d = %d * %d\n", regionId, areas[regionId], g.sides[regionId])

		total += g.sides[regionId] * areas[regionId]
	}

	return total
}
