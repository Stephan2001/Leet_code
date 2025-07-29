package besttimetobuyandsellstock

import (
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{100, 4, 200, 1, 3, 2}
		expected := 4
		result := longestConsecutive(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		expected := 9
		result := longestConsecutive(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := []int{1, 0, 1, 2}
		expected := 3
		result := longestConsecutive(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	longest := 1
	current := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] == nums[i-1]+1 {
			current++
		} else {
			if current > longest {
				longest = current
			}
			current = 1
		}
	}

	if current > longest {
		longest = current
	}

	return longest
}
