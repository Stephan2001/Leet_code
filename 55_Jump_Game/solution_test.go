package jumpgame

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		expected := true
		result := canJump(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{3, 2, 1, 0, 4}
		expected := false
		result := canJump(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func canJump(nums []int, t testing.TB) bool {
	curr_max := 0
	for _, jump := range nums {
		if curr_max == 0 && jump < 1 {
			return false
		}
		if jump > curr_max {
			curr_max = jump - 1
		} else {
			curr_max--
		}
	}
	return true
}
