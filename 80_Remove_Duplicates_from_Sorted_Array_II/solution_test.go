package removeduplicatesfromsortedarrayii

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 3}
		expected := 5
		result := removeDuplicates(nums)

		if expected != result {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
		expected := 7
		result := removeDuplicates(nums)

		if expected != result {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func removeDuplicates(nums []int) int {
	seen := make(map[int]int)
	k := 0
	for i := 0; i < len(nums); i++ {
		val, exists := seen[nums[i]]
		if !exists || val < 2 {
			k++
			seen[nums[i]]++
		}
	}
	return k
}
