package majorityelement

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{3, 2, 3}
		expected := 3
		result := majorityElement(nums)

		if expected != result {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{2, 2, 1, 1, 1, 2, 2}
		expected := 2
		result := majorityElement(nums)

		if expected != result {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func majorityElement(nums []int) int {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		seen[nums[i]]++
		if seen[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
