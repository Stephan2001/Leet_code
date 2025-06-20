package jumpgame

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		expected := 2
		result := jump(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{2, 3, 0, 1, 4}
		expected := 2
		result := jump(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func jump(nums []int, t testing.TB) int {
	curr_max := 0
	replaced := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > curr_max {
			curr_max = nums[i] - 1
			replaced++
		} else {
			curr_max--
		}
	}
	return replaced
}
