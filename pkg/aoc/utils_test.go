package aoc

import "testing"

func TestGCD(t *testing.T) {

	t.Run("expected GCDs", func(t *testing.T) {
		ans := GCD(20, 10)
		if ans != 10 {
			t.Errorf("GCD(20, 10) = %d; want 10", ans)
		}

		ans = GCD(20, 11)
		if ans != 1 {
			t.Errorf("GCD(20, 11) = %d; want 1", ans)
		}

		ans = GCD(100, 35)
		if ans != 5 {
			t.Errorf("GCD(100, 35) = %d; want 5", ans)
		}

		ans = GCD(1000, 300)
		if ans != 100 {
			t.Errorf("GCD(1000, 300) = %d; want 100", ans)
		}

		ans = GCD(17, 84)
		if ans != 1 {
			t.Errorf("GCD(17,84) = %d; want 1", ans)
		}
	})
}
