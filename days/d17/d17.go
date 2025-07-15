package d17

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/jacobrlewis/advent_2024/pkg/aoc"
)

type Computer struct {
	A            int
	B            int
	C            int
	Instructions []int
	Current      int
	Output       []int
}

func parseInput(c *Computer, file *os.File) {

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)
	var line string

	scanner.Scan()
	line = scanner.Text()
	c.A = aoc.StringToInt(re.FindAllString(line, 1)[0])

	scanner.Scan()
	line = scanner.Text()
	c.B = aoc.StringToInt(re.FindAllString(line, 1)[0])

	scanner.Scan()
	line = scanner.Text()
	c.C = aoc.StringToInt(re.FindAllString(line, 1)[0])

	// new line
	scanner.Scan()

	scanner.Scan()
	line = scanner.Text()
	nums := re.FindAllString(line, -1)

	ints := make([]int, len(nums))
	for i, s := range nums {
		ints[i] = aoc.StringToInt(s)
	}
	c.Instructions = ints
}

func (c *Computer) ComboOperand(o int) int {
	if o <= 3 {
		return o
	}

	switch o {
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		panic("invalid combo operand")
	}
}

func (c *Computer) RunInstruction(i int, o int) {
	fmt.Printf("Running i: %d o: %d\n", i, o)

	switch i {
	case 0:
		// adv: division, A over 2 ^ combo operand. truncated written to A
		c.A = c.A / aoc.IntPow(2, c.ComboOperand(o))
	case 1:
		// bxl: bitwise XOR
		c.B ^= o
	case 2:
		// bst
		c.B = c.ComboOperand(o) % 8
	case 3:
		// jnz
		if c.A != 0 {
			c.Current = o
		}
	case 4:
		// bxc
		c.B = c.C ^ c.B
	case 5:
		// out
		c.Output = append(c.Output, c.ComboOperand(o)%8)
	case 6:
		// bdv
		c.B = c.A / aoc.IntPow(2, c.ComboOperand(o))
	case 7:
		// cdv
		c.C = c.A / aoc.IntPow(2, c.ComboOperand(o))
	default:
		panic("invalid instruction")
	}

	if !(i == 3 && c.A != 0) {
		// move up 2 unless we jumped
		c.Current += 2
	}
}

func (c *Computer) Run() {
	for c.Current < len(c.Instructions) {
		fmt.Printf("%+v\n", c)
		c.RunInstruction(c.Instructions[c.Current], c.Instructions[c.Current+1])
	}
}

func Part1(file *os.File) int {

	var c Computer
	parseInput(&c, file)

	c.Run()

	for _, d := range c.Output {
		fmt.Printf("%d,", d)
	}
	fmt.Println()

	return 0
}

func Part2(file *os.File) int {
	return 0
}
