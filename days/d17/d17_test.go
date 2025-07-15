package d17

import (
	"slices"
	"testing"
)

func TestComputer(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		var c Computer
		c.C = 9
		c.Instructions = []int{2, 6}

		c.Run()

		if c.B != 1 {
			t.Error("Expected c.B = 1")
		}
	})

	t.Run("2", func(t *testing.T) {
		var c Computer
		c.A = 10
		c.Instructions = []int{5, 0, 5, 1, 5, 4}

		c.Run()

		if !slices.Equal(c.Output, []int{0, 1, 2}) {
			t.Error("failed output")
		}
	})

	t.Run("3", func(t *testing.T) {
		var c Computer
		c.A = 2024
		c.Instructions = []int{0, 1, 5, 4, 3, 0}

		c.Run()

		expected := []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}
		if !slices.Equal(c.Output, expected) {
			t.Errorf("failed output actual|expected\n%v\n%v\n", c.Output, expected)
		}

		if c.A != 0 {
			t.Error("failed A value")
		}
	})
}
